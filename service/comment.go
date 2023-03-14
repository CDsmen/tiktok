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
	err = dao.Comment_add(video_id, claim.UserID, text, comment)
	if err != nil {
		return err
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
	err = dao.Comment_del(comment_id)
	if err != nil {
		return err
	}
	return nil
}
