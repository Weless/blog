package main

import (
	"blog/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	InitConfig()
	dao.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	r.Run(":8000")
}

func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Println("failed to get work dir, err:", err)
		return
	}
	fmt.Println(workDir)
	viper.AddConfigPath(workDir + "/config") // 设置文件路径
	viper.SetConfigName("application")       // 设置读取文件的名字
	viper.SetConfigType("yml")               // 设置读取文件的类型

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
