package main

import(
	"tiktok/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.InitDB()
	dao.InitRDB()

	r := gin.Default()
	initRouter(r)

	r.Run()
}
