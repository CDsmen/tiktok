package controller

import (
	"net/http"
	"tiktok/myjwt"
	"tiktok/service"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userId, err := service.Login(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Login Failed"},
		})
		return
	} else {
		if userId == 0 {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User not exist"},
			})
			return
		} else {
			// 获得token
			claims := &myjwt.JWTClaims{
				UserID:   userId,
				Username: username,
				Password: password,
			}
			claims.IssuedAt = time.Now().Unix()
			claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(myjwt.ExpireTime)).Unix()
			signedToken, err := myjwt.GetToken(claims)
			if err != nil {
				c.String(http.StatusNotFound, err.Error())
				return
			}

			myjwt.TakenGetMap(signedToken)

			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   userId,
				Token:    signedToken,
			})
			return
		}
	}
}
