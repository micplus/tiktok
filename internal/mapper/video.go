package mapper

import (
	"log"
	"tiktok/internal/model"
)

const (
	limit = 30
)

func VideosByTimeUserID(now, userID int64) ([]*model.Video, error) {
	var videos []*model.Video

	db.Raw(`
	SELECT * FROM videos WHERE videos.user_id=1 
	LEFT JOIN user_videos ON user_videos.user_id=videos.user_id
	`, userID)


`
	WITH result_videos AS (
		SELECT * FROM videos WHERE user_id=?
	)
	
`, userID
}

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

func VideosByIDs(ids []int64) ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Order("created_at DESC").
		Limit(limit).
		Find(&videos, ids).Error
	if err != nil {
		log.Println("mapper.VideosByIDs: ", err)
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
