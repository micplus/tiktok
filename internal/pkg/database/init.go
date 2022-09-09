package database

import (
	"log"
	"tiktok/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init() {
	var err error
	DB, err = sqlx.Connect("mysql", config.DSN)
	if err != nil {
		log.Fatalln("database.init: ", err)
	}
}
