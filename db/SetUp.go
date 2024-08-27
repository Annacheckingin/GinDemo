package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup() {
	dsn := "root:123@tcp(127.0.0.1:3306)/vapor?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db = _db
}

var Db *gorm.DB
