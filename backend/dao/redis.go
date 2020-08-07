package dao

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RDB *redis.Client

func InitRedis() error {
	addr := viper.GetString("redis.addr")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	dbNum := viper.GetInt("redis.DB")

	options := redis.Options{
		Addr:     addr + ":" + port,
		Password: password,
		DB:       dbNum,
	}
	RDB = redis.NewClient(&options)
	_, err := RDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
