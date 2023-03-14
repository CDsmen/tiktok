package myRedis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RdbVsClist *redis.Client // 视频id - 被评论列表
var RdbVsF *redis.Client     // 视频id - 被点赞总数
var RdbVsC *redis.Client     // 视频id - 被评论总数
var RdbUsF *redis.Client     // 用户id - 被点赞总数
var RdbUsV *redis.Client     // 用户id - 做视频总数
var RdbU2F *redis.Client     // 用户id - 点赞别人总数

func InitRDB() {
	log.Println("InitRDB start")
	RdbVsClist = redis.NewClient(&redis.Options{
		Addr:     "172.17.235.73:6379",
		Password: "",
		DB:       0,
	})
	_, err := RdbVsClist.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	RdbVsF = redis.NewClient(&redis.Options{
		Addr:     "172.17.235.73:6379",
		Password: "",
		DB:       1,
	})
	_, err = RdbVsF.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	RdbVsC = redis.NewClient(&redis.Options{
		Addr:     "172.17.235.73:6379",
		Password: "",
		DB:       2,
	})
	_, err = RdbVsC.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	RdbUsF = redis.NewClient(&redis.Options{
		Addr:     "172.17.235.73:6379",
		Password: "",
		DB:       3,
	})
	_, err = RdbUsF.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	RdbUsV = redis.NewClient(&redis.Options{
		Addr:     "172.17.235.73:6379",
		Password: "",
		DB:       4,
	})
	_, err = RdbUsV.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	RdbU2F = redis.NewClient(&redis.Options{
		Addr:     "172.17.235.73:6379",
		Password: "",
		DB:       5,
	})
	_, err = RdbU2F.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

}
