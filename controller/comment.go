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
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	text := c.Query("comment_text")
	commentId := c.Query("comment_id")

	if actionType == "1" {
		videoid_int64, err := strconv.ParseInt(videoId, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "videoId err"},
			})
			return
		}
		var comment service.Comment
		err = service.AddComment(videoid_int64, text, token, &comment)
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "AddComment err"},
			})
			return
		} else {
			c.JSON(http.StatusOK, CommentListResponse{
				Response:    Response{StatusCode: 0},
				CommentList: []service.Comment{comment},
			})
			return
		}
	} else if actionType == "2" {
		commentid_int64, err := strconv.ParseInt(commentId, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "commentId err"},
			})
			return
		}
		err = service.DelComment(commentid_int64, token)
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "DelComment err"},
			})
			return
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0, StatusMsg: "Delete succeeded"},
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "actionType err"},
		})
		return
	}
}
