package mapper

import (
	"log"
	"tiktok/config"
	"tiktok/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(config.Detail.DataSourceName()), &gorm.Config{})
	if err != nil {
		log.Println("mapper.init: ", err)
	}
	if err := db.AutoMigrate(&model.Video{}, &model.User{}, &model.UserLogin{}, &model.Comment{}); err != nil {
		log.Println("mapper.init: ", err)
	}
}
