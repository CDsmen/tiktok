package dao

import (
	"context"
	"log"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RdbV2C *redis.Client

func InitRDB() {
	log.Println("InitRDB start")
	RdbV2C := redis.NewClient(&redis.Options{
		Addr:     "172.17.235.73:6379",
		Password: "",
		DB:       0,
	})
	_, err := RdbV2C.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}
}
