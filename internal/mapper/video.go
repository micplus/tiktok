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
	err := db.Model(&model.Video{}).
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
	err := db.Where("user_id=?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&videos).Error
	if err != nil {
		log.Println("mapper.VideosByTime: ", err)
		return nil, err
	}
	return videos, nil
}

func CreateVideo(video *model.Video) error {
	return db.Create(video).Error
}
