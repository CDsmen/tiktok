package dao

func User_register(username string, password string) (int64, error) {
	var userId int64
	err := DB.Raw("CALL register(?, ?)", username, password).Scan(&userId).Error
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func User_login(username string, password string) (int64, error) {
	var userId int64
	err := DB.Raw("CALL login(?, ?)", username, password).Scan(&userId).Error
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func User_info(user_id int64, user interface{}) error {
	err := DB.Raw("CALL user_info(?)", user_id).Scan(user).Error
	return err
}

func User_UsF(user_id int64) (int64, error) {
	var totalFavorited int64
	err := DB.Raw("CALL get_t_f(?)", user_id).Scan(&totalFavorited).Error
	if err != nil {
		return 0, err
	}
	return totalFavorited, nil
}

func User_UsV(user_id int64) (int64, error) {
	var workCount int64
	err := DB.Raw("CALL get_w_c(?)", user_id).Scan(&workCount).Error
	if err != nil {
		return 0, err
	}
	return workCount, nil
}

func User_U2F(user_id int64) (int64, error) {
	var favoriteCount int64
	err := DB.Raw("CALL get_f_c(?)", user_id).Scan(&favoriteCount).Error
	if err != nil {
		return 0, err
	}
	return favoriteCount, nil
}
