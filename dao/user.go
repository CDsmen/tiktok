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
	var t_f int64
	err := DB.Raw("CALL get_t_f(?)", user_id).Scan(&t_f).Error
	if err != nil {
		return 0, err
	}
	return t_f, nil
}

func User_UsV(user_id int64) (int64, error) {
	var w_c int64
	err := DB.Raw("CALL get_w_c(?)", user_id).Scan(&w_c).Error
	if err != nil {
		return 0, err
	}
	return w_c, nil
}

func User_U2F(user_id int64) (int64, error) {
	var f_c int64
	err := DB.Raw("CALL get_f_c(?)", user_id).Scan(&f_c).Error
	if err != nil {
		return 0, err
	}
	return f_c, nil
}
