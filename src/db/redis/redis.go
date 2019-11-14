package redis

import (
	"project/src/config"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

//Client redis实例
var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Config.Redis.Host, config.Config.Redis.Port),
		Password: config.Config.Redis.Password,
		DB:       0,
	})
}

//Paths 前缀
var Paths = map[string]string{
	"printToken":    "dis:3",
	"homeStatic":    "dis:4",
	"requestRecord": "dis:5",
}

//GetValue 通过key获取value值
func GetValue(key string) (string, error) {
	var value string
	value, err := Client.Get(key).Result()
	if err == redis.Nil {
		logrus.Error("key dose not exists")
		return value, err
	} else if err != nil {
		logrus.Error("product find basicInfo err ---->", err)
		return value, err
	}
	return value, nil
}
