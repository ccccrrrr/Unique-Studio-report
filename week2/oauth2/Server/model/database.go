package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB_server *gorm.DB
)

func InitServerDatabase() {

	path := "root: @(localhost:3306)/db2?charset=utf8mb4&parseTime=True&loc=Local"

	_db, err := gorm.Open(mysql.Open(path), &gorm.Config{})

	if err != nil {
		log.Println(err)
		return
	}

	_ = _db.AutoMigrate(&AuthzCodeInfo{})
	_ = _db.AutoMigrate(&AccessTokenInfo{})
	_ = _db.AutoMigrate(&UserInfo{})
	_ = _db.AutoMigrate(&Picture{})

	DB_server = _db

}