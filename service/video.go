package service

import (
	"strconv"
	"tiktok/dao"
	"tiktok/myRedis"
	"tiktok/myjwt"
	"time"
)

func Feed(latestTime string, token string, videoList *[]Video) (int64, error) {
	// token不存在
	err := myjwt.FindToken(token)
	if err != nil {
		return 0, err
	}
	// 解析token
	_, err = myjwt.VerifyAction(token)
	if err != nil {
		return 0, err
	}

	if latestTime == "0" || latestTime == "" {
		latestTime = strconv.FormatInt(time.Now().Unix(), 10)
	}
	if len([]rune(latestTime)) > 11 {
		latestTime = latestTime[0:10]
	}

	err = dao.Feed(latestTime, videoList)
	if err != nil {
		return 0, err
	}

	nextTime := time.Now().Unix()
	for id := range *videoList {
		err = FullVideo(&(*videoList)[id], token)
		if err != nil {
			return 0, err
		}
		if (*videoList)[id].CreateTime < nextTime {
			nextTime = (*videoList)[id].CreateTime
		}
	}
	return nextTime, nil
}

func ListPublish(user_id int64, token string, videoList *[]Video) error {
	// token不存在
	err := myjwt.FindToken(token)
	if err != nil {
		return err
	}

	// 解析token
	_, err = myjwt.VerifyAction(token)
	if err != nil {
		return err
	}

	err = dao.Video_list(user_id, videoList)
	if err != nil {
		return err
	}

	for id := range *videoList {
		err = FullVideo(&(*videoList)[id], token)
		if err != nil {
			return err
		}
	}
	return nil
}

func FullVideo(video *Video, token string) error {
	err := UserInfo(strconv.FormatInt(video.Userid, 10), token, &(*video).Author)
	if err != nil {
		return err
	}

	// 补充IsFavorite
	video.IsFavorite, err = dao.Video_IsFavorite(video.Userid, video.Id)
	if err != nil {
		return err
	}

	// 补充FavoriteCount
	if n, err := myRedis.RdbVsF.Exists(myRedis.Ctx, strconv.FormatInt(video.Id, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		video.FavoriteCount, err = myRedis.RdbVsF.Get(myRedis.Ctx, strconv.FormatInt(video.Id, 10)).Int64()
		if err != nil {
			return err
		}
	} else {
		video.FavoriteCount, err = dao.Video_VsF(video.Id)
		if err != nil {
			return err
		}
		myRedis.RdbVsF.Set(myRedis.Ctx, strconv.FormatInt(video.Id, 10), video.FavoriteCount, 0)
	}

	// 补充CommentCount
	if n, err := myRedis.RdbVsC.Exists(myRedis.Ctx, strconv.FormatInt(video.Id, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		video.CommentCount, err = myRedis.RdbVsC.Get(myRedis.Ctx, strconv.FormatInt(video.Id, 10)).Int64()
		if err != nil {
			return err
		}
	} else {
		video.CommentCount, err = dao.Video_VsC(video.Id)
		if err != nil {
			return err
		}
		myRedis.RdbVsC.Set(myRedis.Ctx, strconv.FormatInt(video.Id, 10), video.CommentCount, 0)
	}

	return nil
}
