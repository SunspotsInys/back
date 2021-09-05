package configs

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

func init() {
	// 读取配置文件路径
	var configPath string = "../config.yaml"
	flag.StringVar(&configPath, "c", "../config.yaml", "配置文件路径config path")
	flag.Parse()
	log.Println(configPath)

	// 读取配置文件内容
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}

	// 将配置文件写入结构体
	err = viper.Unmarshal(Conf)
	if err != nil {
		log.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
}
