package model

import (
	"gorm.io/gorm"
	"log"
)

type UserInfo struct {
	gorm.Model
	UserName     string
	UserPassword string
}

func IsValid(user UserInfo) bool {
	temp := UserInfo{}
	username := user.UserName
	userpassword := user.UserPassword
	if err := DB_server.Where("user_name = ?", username).First(&temp).Error; err != nil {
		log.Println(err)
		return false
	} else {
		if temp.UserPassword == userpassword {
			return true
		} else {
			return false
		}
	}
}

func DeleteUser(username string, userpassword string) bool {
	return false
}

func CreateUser(user UserInfo) bool {
	err := DB_server.Create(&user)
	if err == nil {
		return true
	} else {
		return false
	}
}

func IsExistUserName(username string) bool {
	temp := UserInfo{}
	if err := DB_server.Where("user_name = ?", username).First(&temp).Error; err != nil {
		log.Println(err)
		return false
	} else {
		return false
	}
}
