package service

import (
	"tiktok/dao"
	"tiktok/myjwt"
	"time"
)

func Login(username string, password string) (int64, string, error) {

	// 获得Id
	userId, err := dao.User_login(username, password)
	if err != nil || userId == 0 {
		return userId, "", err
	}

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
		return userId, "", err
	}

	myjwt.TakenGetMap(signedToken)

	return userId, signedToken, nil
}

func Register(username string, password string) (int64, string, error) {
	// 获得Id
	userId, err := dao.User_register(username, password)
	if err != nil || userId == 0 {
		return userId, "", err
	}

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
		return userId, "", err
	}

	myjwt.TakenGetMap(signedToken)

	return userId, signedToken, nil
}
