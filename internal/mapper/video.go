package mapper

import (
	"log"
	"tiktok/internal/model"
)

const (
	limit = 30
)

func VideosByTime(now int64) ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Joins("User").
		Where("created_at<?", now).
		Order("created_at DESC").
		Limit(limit).
		Find(&videos).Error
	if err != nil {
		log.Println("mapper.VideosByTime: ", err)
		return nil, err
	}
	return videos, nil
}

func VideosByUserID(userID int64) ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Joins("User", db.Where(&model.User{ID: userID})).
		Order("created_at DESC").
		Limit(limit).
		Find(&videos).Error
	if err != nil {
		log.Println("mapper.VideosByTime: ", err)
		return nil, err
	}
	return videos, nil
}

func CreateVideoWithUserID(video *model.Video, userID int64) error {
	user := &model.User{ID: userID}
	return db.Model(user).Association("Published").Append(video)
}
