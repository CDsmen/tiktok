package service

import (
	"strconv"
	"tiktok/dao"
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
		var user User
		err = UserInfo(strconv.FormatInt((*commentsList)[id].Userid, 10), token, &(*commentsList)[id].User)
		if err != nil {
			return err
		}
		(*commentsList)[id].User = user
	}
	return nil
}
