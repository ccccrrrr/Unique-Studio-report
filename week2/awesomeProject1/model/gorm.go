package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)
var (
	Db *gorm.DB
	//dbUser *gorm.DB
	//dbPicture *gorm.DB
)
func StartDatabase() {
	path := "root: @(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(path), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	_ = _db.AutoMigrate(&User{})
	_ = _db.AutoMigrate(&Picture{})

	//dbUser = _db
	//dbPicture = _db.Table("pictures")

	Db = _db

}
