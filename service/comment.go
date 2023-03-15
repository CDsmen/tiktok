package service

import (
	"strconv"
	"tiktok/dao"
	"tiktok/myRedis"
	"tiktok/myjwt"
)

func CommentList(video_id int64, token string, commentsList *[]Comment) error {
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

	err = dao.Comment_list(video_id, commentsList)
	if err != nil {
		return err
	}

	// 补充user
	for id := range *commentsList {
		err = UserInfo(strconv.FormatInt((*commentsList)[id].Userid, 10), token, &(*commentsList)[id].User)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddComment(video_id int64, text string, token string, comment *Comment) error {
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

	// 先修改数据库
	err = dao.Comment_add(video_id, claim.UserID, text, comment)
	if err != nil {
		return err
	}
	err = UserInfo(strconv.FormatInt(comment.Userid, 10), token, &comment.User)
	if err != nil {
		return err
	}
	// 再修改redis
	if n, err := myRedis.RdbVsC.Exists(myRedis.Ctx, strconv.FormatInt(video_id, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		_, err = myRedis.RdbVsC.Incr(myRedis.Ctx, strconv.FormatInt(video_id, 10)).Result()

		if err != nil {
			return err
		}
	}
	return nil
}

func DelComment(comment_id int64, token string) error {
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
	var video_id int64
	err = dao.Comment_del(comment_id, &video_id)
	if err != nil {
		return err
	}
	if n, err := myRedis.RdbVsC.Exists(myRedis.Ctx, strconv.FormatInt(video_id, 10)).Result(); n > 0 {
		if err != nil {
			return err
		}
		_, err = myRedis.RdbVsC.Decr(myRedis.Ctx, strconv.FormatInt(video_id, 10)).Result()

		if err != nil {
			return err
		}
	}
	return nil
}
