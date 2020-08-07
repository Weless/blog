package main

import (
	"blog/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	err := InitConfig()
	if err != nil {
		fmt.Println("failed to init config, err:", err)
		return
	}
	err = dao.InitDB()
	if err != nil {
		fmt.Println("failed to connect to database, err:" + err.Error())
		return
	}
	err = dao.InitRedis()
	if err != nil {
		fmt.Println("fail to connect to redis,err:", err.Error())
		return
	}
	r := gin.Default()
	r = CollectRoute(r)
	r.Run(":8000")
}

func InitConfig() error {
	workDir, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.AddConfigPath(workDir + "/config") // 设置文件路径
	viper.SetConfigName("application")       // 设置读取文件的名字
	viper.SetConfigType("yml")               // 设置读取文件的类型

	err = viper.ReadInConfig()
	return err
}
