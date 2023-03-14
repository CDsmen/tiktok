package service

import(
	"tiktok/dao"
)

func Login(username string, password string)(int64, error){
	return dao.User_login(username,password)
}