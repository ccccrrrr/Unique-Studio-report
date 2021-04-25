package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	_ = _db.AutoMigrate(&AuthInfo{})
	_ = _db.AutoMigrate(&AccessTokenInfo{})

	Db = _db

}

func InsertAuthCodeInfo(authInfo AuthInfo) {
	var tmp AuthInfo
	if err := Db.Where("user_name = ?", authInfo.UserName).First(&tmp).Error; err != nil {
		// expireTime, _ := strconv.Atoi(tmp.ExpireTime)
		// update
			Db.Delete(&tmp)
			Db.Create(&authInfo)
	}else {
		//
		Db.Create(&authInfo)
	}
}

func IsValidToken(accessToken string) bool {
	var temp AccessTokenInfo
	// need to verify expire time
	if err := Db.Where("access_token = ?", accessToken).First(&temp).Error; err != nil {
		return false
	}else {
		return true
	}
}
