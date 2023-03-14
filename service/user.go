package service

import (
	"strconv"
	"tiktok/dao"
	"tiktok/myRedis"
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

func UserInfo(userid string, token string, user *User) error {
	// token不存在
	err := myjwt.FindToken(token)
	if err != nil {
		return err
	}

	// 解析token
	claim, err := myjwt.VerifyAction(token)
	if err != nil {
		return err
	}

	// 鉴权不通过
	if strconv.FormatInt(claim.UserID, 10) != userid {
		return err
	}

	// 获取user信息
	err = dao.User_info(claim.UserID, user)
	if err != nil {
		return err
	}

	// 获取TotalFavorited
	if n, err := myRedis.RdbUsF.Exists(myRedis.Ctx, userid).Result(); n > 0 {
		if err != nil {
			return err
		}
		user.TotalFavorited, err = myRedis.RdbUsF.Get(myRedis.Ctx, userid).Result()
		if err != nil {
			return err
		}
	} else {
		totalFavorited, err := dao.User_UsF(user.Id)
		if err != nil {
			return err
		}
		user.TotalFavorited = strconv.FormatInt(totalFavorited, 10)
		myRedis.RdbUsF.Set(myRedis.Ctx, userid, totalFavorited, 0)
	}

	// 获取WorkCount
	if n, err := myRedis.RdbUsV.Exists(myRedis.Ctx, userid).Result(); n > 0 {
		if err != nil {
			return err
		}
		user.WorkCount, err = myRedis.RdbUsV.Get(myRedis.Ctx, userid).Int64()
		if err != nil {
			return err
		}
	} else {
		workCount, err := dao.User_UsV(user.Id)
		if err != nil {
			return err
		}
		user.WorkCount = workCount
		myRedis.RdbUsV.Set(myRedis.Ctx, userid, workCount, 0)
	}

	// 获取FavoriteCount
	if n, err := myRedis.RdbU2F.Exists(myRedis.Ctx, userid).Result(); n > 0 {
		if err != nil {
			return err
		}
		user.FavoriteCount, err = myRedis.RdbU2F.Get(myRedis.Ctx, userid).Int64()
		if err != nil {
			return err
		}
	} else {
		favoriteCount, err := dao.User_U2F(user.Id)
		if err != nil {
			return err
		}
		user.FavoriteCount = favoriteCount
		myRedis.RdbU2F.Set(myRedis.Ctx, userid, favoriteCount, 0)
	}

	return nil
}
