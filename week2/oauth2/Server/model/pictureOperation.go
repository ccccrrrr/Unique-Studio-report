package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type Picture struct {
	gorm.Model
	PictureName string `json:"picture_name"`
	CreateUser  string `json:"create_user"`
}

func Upload(picture_name string, user_name string) bool {
	var p1 []Picture
	var p Picture
	if err := DB_server.Table("pictures").Where("picture_name = ?", picture_name).Find(&p1).Error; err == nil {
		log.Println(p1)
		for _, info := range p1 {
			if info.CreateUser == user_name {
				log.Println("[upload error] already upload the picture!")
				return false
			}
		}
	}
	p.PictureName= picture_name
	p.CreateUser = user_name
	DB_server.Table("pictures").Create(&p)
	return true
}

func DeletePicture(picture_name string, username string) error {
	var tmp []Picture
	if err := DB_server.Table("pictures").Where("picture_name = ?", picture_name).Find(&tmp).Error; err != nil {
		return err
	}
		for _, info := range tmp {
			if info.CreateUser== username {
				DB_server.Table("pictures").Delete(&info)
				return nil
			}
	}
	return errors.New("[picture delete error] no picture name match.")
}

func SearchPicture(username string) error {
	var pictures []Picture
	if err := DB_server.Where("create_user = ?", username).Find(&pictures).Error; err != nil {
		log.Println(err)
		return err
	}
	for _, info := range pictures {
		log.Println(info)
	}
	return nil
}