package model

import (
	"gorm.io/gorm"
	"time"
)

const ExpireTime = time.Duration(time.Second * 60)

type AuthCodeInfo struct {
	gorm.Model
	AuthzCode string `json:"authz_code" gorm:"column:authz_code"`
	UserName  string `json:"user_name" gorm:"user_name"`
}

func InsertAuthCodeInfo(authcodeinfo AuthCodeInfo) bool {
	appDB.Create(&authcodeinfo)
	return true
}