package model

import (
	"gorm.io/gorm"
	"log"
)

type Picture struct {
	gorm.Model
	Name string `json:"picture_name"`
	Path string `json:"picture_path"`
	CreateUserName string `json:"create_user_name"`
	CreateUserPassword string `json:"create_user_password"`
//	CreateTime time.Time `json:"create_time"`
//	Size string `json:"picture_size"`
}

func Upload(base_path string, picture Picture) bool {
	// different directory path need to add
	var p1 Picture
	//picture.CreateTime = time.Now()
	// also need to see username and picture
	if err := Db.Table("pictures").Where("name = ?", picture.Name).First(&p1).Error; err == nil {
		log.Println("not found!")
		return false
	}else {
		Db.Table("pictures").Create(&picture)
	}
	return true
}