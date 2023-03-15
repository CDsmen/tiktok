package controller

import (
	"net/http"
	"strconv"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

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
