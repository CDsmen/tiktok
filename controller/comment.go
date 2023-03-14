package controller

import (
	"net/http"
	"strconv"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")

	var commentList []service.Comment
	videoid_int64, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "videoId err"},
		})
		return
	}
	err = service.CommentList(videoid_int64, token, &commentList)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "CommentList get err"},
		})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: commentList,
	})
}
