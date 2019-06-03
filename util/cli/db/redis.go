package db

import (
	"strings"

	"github.com/go-redis/redis"
)

// RDS a global Redis client
// External logic can use it while
// use the one in martini context firstly
var RDS redis.UniversalClient

// RedisConfig 用于解析redis配置
type RedisConfig struct {
	Cluster  string `json:"cluster"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func newRedisClient(addr string, pwd string, db int) (redis.UniversalClient, error) {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{addr},
		Password: pwd,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func newRedisClusterClient(nodesAddr string) (redis.UniversalClient, error) {
	addrsArray := strings.Split(nodesAddr, ",")
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: addrsArray,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewUniversalRedisClient create universal redis client with config
// cluster client or single instance client
func NewUniversalRedisClient(conf *RedisConfig) (redis.UniversalClient, error) {
	var client redis.UniversalClient
	var err error

	if conf.Cluster != "" {
		client, err = newRedisClusterClient(conf.Cluster)
	} else {
		client, err = newRedisClient(conf.Addr, conf.Password, conf.DB)
	}
	return client, err
}

// InitRedisCli create a redis handler
func InitRedisCli(conf *RedisConfig) error {
	client, err := NewUniversalRedisClient(conf)
	if err != nil {
		return err
	}
	RDS = client
	return nil
}
