package controller

import (
	"net/http"
	"strconv"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	// 获取上传的文件
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "获取上传的文件 error",
		})
		return
	}

	err = service.AddVideo(title, token, file)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Publish error",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  file.Filename + " uploaded successfully",
	})
}

func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	token := c.Query("token")

	var videoList []service.Video
	nextTime, err := service.Feed(latestTime, token, &videoList)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "feed error"},
		})
		return
	} else if len(videoList) == 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0, StatusMsg: "没有更多视频了"},
			VideoList: videoList,
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			NextTime:  nextTime,
			VideoList: videoList,
		})
	}
}

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
