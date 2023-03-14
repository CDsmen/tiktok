package main

import (
	"tiktok/dao"
	"tiktok/myRedis"

	"github.com/gin-gonic/gin"
)

func main() {
	dao.InitDB()
	myRedis.InitRDB()

	r := gin.Default()
	initRouter(r)

	r.Run()
}
