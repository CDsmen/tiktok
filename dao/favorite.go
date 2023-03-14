package dao

func Favorite_add(user_id int64, video_id int64) error {
	err := DB.Exec("CALL add_favorite(?, ?)", user_id, video_id).Error
	return err
}

func Favorite_del(user_id int64, video_id int64) error {
	err := DB.Exec("CALL del_favorite(?, ?)", user_id, video_id).Error
	return err
}

func Favorite_list(user_id int64, video interface{}) error {
	err := DB.Raw("CALL list_favorite(?)", user_id).Scan(video).Error
	return err
}
