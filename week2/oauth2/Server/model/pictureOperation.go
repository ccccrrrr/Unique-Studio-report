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

func Show(user_name string) error {
	var p1 []Picture
	if err := DB_server.Table("pictures").Where("create_user_name = ?", user_name).Find(&p1).Error; err == nil {
		for _, info := range p1 {
			log.Println(info)
		}
	} else {
		return err
	}
	return nil
}

func Upload(picture_name string, user_name string) bool {
	var p1 []Picture
	var p Picture
	if err := DB_server.Table("pictures").Where("name = ?", picture_name).Find(&p1).Error; err == nil {
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
	if err := DB_server.Table("pictures").Where("name = ?", picture_name).Find(&tmp).Error; err != nil {
		return err
	}
		for _, info := range tmp {
			if info.CreateUser== username {
				DB_server.Table("pictures").Delete(&info)
				return nil
			}
	}
	return errors.New("no picture name match.")
}
