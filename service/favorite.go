package service

import (
	"strconv"
	"tiktok/dao"
	"tiktok/myRedis"
	"tiktok/myjwt"
)

func AddFavorite(video_id int64, token string) error {
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

	err = dao.Favorite_add(claim.UserID, video_id)
	if err != nil {
		return err
	}

	if n, err := myRedis.RdbVsF.Exists(myRedis.Ctx, strconv.FormatInt(video_id, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		_, err = myRedis.RdbVsF.Incr(myRedis.Ctx, strconv.FormatInt(video_id, 10)).Result()

		if err != nil {
			return err
		}
	}

	if n, err := myRedis.RdbU2F.Exists(myRedis.Ctx, strconv.FormatInt(claim.UserID, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		_, err = myRedis.RdbU2F.Incr(myRedis.Ctx, strconv.FormatInt(claim.UserID, 10)).Result()

		if err != nil {
			return err
		}
	}

	vuserid, err := dao.Video_Vid2Uid(video_id)
	if err != nil {
		return err
	}

	if n, err := myRedis.RdbUsF.Exists(myRedis.Ctx, strconv.FormatInt(vuserid, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		_, err = myRedis.RdbUsF.Incr(myRedis.Ctx, strconv.FormatInt(vuserid, 10)).Result()

		if err != nil {
			return err
		}
	}
	return nil
}

func DelFavorite(video_id int64, token string) error {
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

	err = dao.Favorite_del(claim.UserID, video_id)
	if err != nil {
		return err
	}

	if n, err := myRedis.RdbVsF.Exists(myRedis.Ctx, strconv.FormatInt(video_id, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		_, err = myRedis.RdbVsF.Decr(myRedis.Ctx, strconv.FormatInt(video_id, 10)).Result()

		if err != nil {
			return err
		}
	}

	if n, err := myRedis.RdbU2F.Exists(myRedis.Ctx, strconv.FormatInt(claim.UserID, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		_, err = myRedis.RdbU2F.Decr(myRedis.Ctx, strconv.FormatInt(claim.UserID, 10)).Result()

		if err != nil {
			return err
		}
	}

	vuserid, err := dao.Video_Vid2Uid(video_id)
	if err != nil {
		return err
	}

	if n, err := myRedis.RdbUsF.Exists(myRedis.Ctx, strconv.FormatInt(vuserid, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		_, err = myRedis.RdbUsF.Decr(myRedis.Ctx, strconv.FormatInt(vuserid, 10)).Result()

		if err != nil {
			return err
		}
	}
	return nil
}
