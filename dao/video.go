package dao

func Feed(latestTime string, video interface{}) error {
	err := DB.Raw("CALL feed(?)", latestTime).Scan(video).Error
	return err
}

func Video_byU(user_id int64, video interface{}) error {
	err := DB.Raw("CALL list_video(?)", user_id).Scan(video).Error
	return err
}

func Video_IsFavorite(user_id int64, video_id int64) (bool, error) {
	var isFavorite bool
	err := DB.Raw("CALL is_favorite(?, ?)", user_id, video_id).Scan(&isFavorite).Error
	if err != nil {
		return false, err
	}
	return isFavorite, nil
}

func Video_VsF(video_id int64) (int64, error) {
	var favoriteCount int64
	err := DB.Raw("CALL get_VsF(?)", video_id).Scan(&favoriteCount).Error
	if err != nil {
		return 0, err
	}
	return favoriteCount, nil
}

func Video_VsC(video_id int64) (int64, error) {
	var commentCount int64
	err := DB.Raw("CALL get_VsC(?)", video_id).Scan(&commentCount).Error
	if err != nil {
		return 0, err
	}
	return commentCount, nil
}

func Video_add() {

}
