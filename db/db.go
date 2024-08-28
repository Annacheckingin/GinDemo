package db

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(gin *gin.Engine) {

}

func init() {
	dsn := "root:123@tcp(127.0.0.1:3306)/vapor?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db = _db
}

// Db is the database connection
var Db *gorm.DB
