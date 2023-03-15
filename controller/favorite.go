package controller

import (
	"net/http"
	"strconv"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	videoid_int64, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "videoId err"},
		})
		return
	}
	if actionType == "1" {
		err := service.AddFavorite(videoid_int64, token)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Favorite fail"})
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Favorite succeeded"})
		return
	} else if actionType == "2" {
		err := service.DelFavorite(videoid_int64, token)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Delete favorite fail"})
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Delete favorite succeeded"})
		return
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "FavoriteAction Type err"})
	}
}
