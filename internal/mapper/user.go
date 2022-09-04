package mapper

import "tiktok/internal/model"

func CheckUniqueUsername(username string) (bool, error) {
	res := db.Where("username = ?", username).Limit(1).Find(&model.UserLogin{})
	cnt, err := res.RowsAffected, res.Error
	return cnt == 0, err
}

func CreateUser(user *model.User) error {
	return db.Create(user).Error
}

func UserByID(id int64) (*model.User, error) {
	user := &model.User{}
	res := db.Where("id = ?", id).Limit(1).Find(user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}

func LoginByUsername(username string) (*model.UserLogin, error) {
	ul := &model.UserLogin{}
	res := db.Where("username = ?", username).Limit(1).Find(ul)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return ul, nil
}
