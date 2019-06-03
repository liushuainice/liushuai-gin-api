package main

import (
	"fmt"
	"ginWebTest/util/cli/db"
	log "ginWebTest/util/golog"
)

func init() {
	config := db.RedisConfig{
		Cluster:  "",
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	}
	// Redis操作句柄
	if err := db.InitRedisCli(&config); err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
}

type TastA struct {
	name string
}

func spinTast(list []*TastA) {
	lLen := len(list) - 1
	fmt.Println(lLen)
	for i := 0; i < (lLen+1)/2; i++ {
		list[i], list[lLen-i] = list[lLen-i], list[i]
		//fmt.Println(*list[i])
	}
}
func main() {

	//tasts:=[]TastA{{"aa"},{"bb"},{"cc"},{"dd"}}
	//tasts1:=[]*TastA{{"aa"},{"bb"},{"cc"},{"dd"},{"ee"}}
	tasts1 := []*TastA{}
	fmt.Println(tasts1)
	spinTast(tasts1)
	fmt.Println(tasts1)

	//i, e := db.RDS.LPush("asdf", 1, 2, 3, 4, 5, 6, 7, 8).Result()
	//fmt.Println(i,e)
	/*	if fields, err := db.RDS.LRange("asdf",0,-1).Result(); err == nil || err == redis.Nil { //无异常
		//if fields, err := db.RDS.LRange("asdf",-1,0).Result(); err == nil || err == redis.Nil { //无异常
		fmt.Println(111,fields,err,len(fields),111)
			for e := range fields {
			e:=e+100
			fmt.Println(e)
			}



		}else {
			fmt.Println(2222,fields,err,len(fields),222)
		}*/

	/*re, er := db.RDS.Get("aa").Int()
	if er == nil {
		fmt.Println(555,re)
	} else {
		fmt.Println(666,er)
	}*/

	//err := db.RDS.Set("aa", 1, 0).Err()
	//fmt.Println(777,err)
	/*errr := db.RDS.Incr("aaa").Err()
	if errr != nil {
		fmt.Println(888,errr)
	}
	i,_:=db.RDS.Get("aaa").Int()
	fmt.Println(999,i)*/

	//// 这里设置过期时间.--1秒过期
	//err = client.Set("age", "20", 1 * time.Second).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//client.Incr("age") // 自增

	//hash 插入和获取
	/*	err:=hashInsertUserTest()
		fmt.Println(err)
		result, e := db.RDS.HGetAll("test:hashkey").Result()
		if e==nil{
			//fmt.Println(result)
			for k, v := range result {
				fmt.Println(k,"==",v)
			}
		}*/
	//删除key
	//err := db.RDS.Del("test:hashkey").Err()
	//fmt.Println(err)
}
