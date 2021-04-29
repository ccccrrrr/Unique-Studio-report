package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	model2 "oauth2/Server/model"
)

var (
	appDB *gorm.DB
)

func InitDataBase() {

	path := "root: @(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"

	_db, err := gorm.Open(mysql.Open(path), &gorm.Config{})

	if err != nil {
		log.Println(err)
		return
	}

	_ = _db.AutoMigrate(&model2.AccessTokenInfo{})
	_ = _db.AutoMigrate(&AuthCodeInfo{})

	appDB = _db

}

func Create(access_token model2.AccessTokenInfo) bool {
	appDB.Create(&access_token)
	return true
}

func Delete(access_token model2.AccessTokenInfo) bool {
	return false
}

func FindUserInAccessToken(accesstoken string) (*model2.AccessTokenInfo, bool) {
	temp := model2.AccessTokenInfo{}
	if err := appDB.Where("access_token = ?", accesstoken).First(&temp).Error; err == nil {
		return &temp, true
	} else {
		return nil, false
	}
}
