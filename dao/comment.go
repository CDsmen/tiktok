package dao

func Comment_add(video_id int64, user_id int64, text string, comment interface{}) error {
	err := DB.Raw("CALL add_comment(?, ?, ?)", video_id, user_id, text).Scan(comment).Error
	return err
}

func Comment_del(comment_id int64,videoid *int64) error {
	err := DB.Raw("CALL del_comment(?)", comment_id).Scan(videoid).Error
	return err
}

func Comment_list(video_id int64, commentsList interface{}) error {
	err := DB.Raw("CALL list_comment(?)", video_id).Scan(commentsList).Error
	return err
}
