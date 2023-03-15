package service

import (
	"strconv"
	"tiktok/dao"
	"tiktok/myRedis"
)

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
