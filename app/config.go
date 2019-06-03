package app

import (
	"encoding/json"
	"flag"
	"fmt"
	"liushuai-gin-api/util/cli/db"
	"liushuai-gin-api/util/golog"
	"os"
)

// 用于解析配置文件
type config struct {
	Log      *log.Config     `json:"log"`
	Redis    *db.RedisConfig `json:"redis"`
	Mysql    *db.MysqlConfig `json:"mysql"`
	Debug    bool            `json:"debug"`     // 调试日志模式
	BindPort int             `json:"bind_port"` //启动端口
}

// Config 全局配置
var Config *config

// Parse 指定配置文件filename执行解析
func (c *config) parse(filename string) error {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&c)
	if err != nil {
		return err
	}
	/*将日志信息打印到配置文件里,追代码，file为空时不处理，level，rotate，有几个常量，使用时直接引包使用他到方法，*/
	log.Infof("log config (%s)", c.Log.DebugString())
	err = log.ParseConfig(c.Log)
	if err != nil {
		return err
	}
	return nil
}

// Parse 初始化配置实例
func Parse() {
	flag.Parse() //解析命令行参数使用
	filename := "./app/config.json"
	if flag.Arg(0) != "" {
		filename = flag.Arg(0)
	}
	log.Infof("parse config: %s", filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatal(err)
	}

	Config = new(config)
	err := Config.parse(filename)
	if err != nil {
		log.Fatal(err)
	}
	//获取启动指令参数
	args := flag.Args()
	for n, v := range args {
		fmt.Println(n, v)
	}
}
