package model

import (
	"gorm.io/gorm"
	"log"
)

type Picture struct {
	gorm.Model
	Name           string `json:"picture_name"`
	CreateUserName string `json:"create_user_name"`
}

func Show(user_name string) {
	var p1 []Picture
	if err := Db.Table("pictures").Where("create_user_name = ?", user_name).Find(&p1).Error; err == nil {
		for _, info := range p1 {
			log.Println(info)
		}
	} else {
		log.Printf("[warning] user %v has not uploaded any picture", user_name)
	}
}

func Upload(picture_name string, user_name string) bool {
	var p1 []Picture
	var p Picture
	if err := Db.Table("pictures").Where("name = ?", picture_name).Find(&p1).Error; err == nil {
		log.Println(p1)
		for _, info := range p1 {
			if info.CreateUserName == user_name {
				log.Println("[upload error] already upload the picture!")
				return false
			}
		}
	}
		log.Println("yes")
		p.Name = picture_name
		p.CreateUserName = user_name
		Db.Table("pictures").Create(&p)

	return true
}

func DeletePicture(picture_name string, username string) {
	var tmp []Picture
	if err := Db.Table("pictures").Where("name = ?", picture_name).Find(&tmp).Error; err == nil {
		for _, info := range tmp {
			if info.CreateUserName == username {
				Db.Table("pictures").Delete(&info)
			}
		}
	}
}
