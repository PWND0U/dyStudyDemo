package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var c *Config

func C() *Config {
	if c == nil {
		panic("配置信息加载出错")
	}
	return c
}

func loadConfig() {
	//获取项目的执行路径
	path, err := os.Getwd()
	fmt.Println("开始从" + path + "加载配置文件...")
	if err != nil {
		panic(err)
	}
	// 1. 初始化 Viper 库
	vip := viper.New()
	vip.AddConfigPath(path + "/config") //设置读取的文件路径
	vip.SetConfigName("config")         //设置读取的文件名
	// 2. 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	vip.SetConfigType("yaml") //设置文件的类型
	// vip.WatchConfig()
	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}

	err = vip.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}

func defaultConfig() {
	c = &Config{
		App:      defaultApp(),
		Database: defaultDatabase(),
		Log:      defaultLog(),
	}
}

func init() {
	fmt.Println("开始初始化配置...")
	defaultConfig()
	loadConfig()
}
