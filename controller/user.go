package controller

import (
	"net/http"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userId, token, err := service.Login(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Login Failed"},
		})
		return
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userId,
			Token:    token,
		})
		return
	}
}
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userId, token, err := service.Register(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Register Failed"},
		})
		return
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userId,
			Token:    token,
		})
		return
	}
}
func UserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	token := c.Query("token")

	var user service.User

	err := service.UserInfo(userid, token, &user)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get UserInfo err"},
		})
		return
	}else{
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
		return
	}

}
