package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DeleteInfo struct {
	DeleteUser string `json:"delete_user"`
	DeletePictureName string `json:"delete_picture_name"`
}

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

func DeletePicture(picture DeleteInfo) bool {
	tmp := Picture{}
	log.Println("yes")
	log.Println(picture)
	// therefore we must search for every picture in the
	if err := Db.Table("pictures").Where("name = ?", picture.DeletePictureName).First(&tmp).Error; err == nil{
		log.Println(tmp)

		Db.Table("pictures").Delete(&tmp)
		//for _, info := range tmp {
		//	log.Println(info)
		//	if info.CreateUserName == picture.DeleteUser {
		//		Db.Table("pictures").Delete(&info)
		//		return true
		//	}
		//}
		return true
	}
	return false
}
