package main

import (
	"blog/dao"
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

func InitConfig()  {
	workDir, err := os.Getwd()
	if err != nil {
		log.Println("failed to get work dir, err:",err)
		return
	}
	viper.SetConfigName("application")       // 设置读取文件的名字
	viper.SetConfigType("yml")               // 设置读取文件的类型
	viper.AddConfigPath(workDir + "/config") // 设置文件路径
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}