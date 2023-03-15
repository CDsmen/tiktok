package controller

import (
	"net/http"
	"strconv"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

func PublishList(c *gin.Context) {
	userid := c.Query("user_id")
	token := c.Query("token")

	userid_int64, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "userid err"},
		})
		return
	}

	var videoList []service.Video
	err = service.ListPublish(userid_int64, token, &videoList)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "list_publish error"},
		})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
	})
}
