package model

import (
	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type AuthzCodeInfo struct {
	gorm.Model
	UserName    string `json:"user_name" gorm:"column:user_name"`
	AuthCode    string `json:"auth_code" gorm:"column:auth_code"`
	RedirectUri string `json:"redirect_uri" gorm:"column:redirect_uri"`
	Scope       string `json:"scope" gorm:"column:scope"`
	ExpireTime  string `json:"expire_time" gorm:"expire_time"`
}

func GenerateAuthzCode(username string, scope string, expire string, redirectUri string) (AuthzCodeInfo, bool) {
	temp := AuthzCodeInfo{}
	temp.UserName = username
	temp.Scope = scope
	temp.ExpireTime = expire
	temp.RedirectUri = redirectUri
	temp.AuthCode = randstr.Hex(10)
	return temp, true
}

func GetUsernameInAuthCode(auth_code string) (string, bool) {
	return "", true
}

func GetAuthCodeInUsername(user_name string) (string, bool) {
	return "", true
}

func StoreAuthCodeInfo(user AuthzCodeInfo) bool {
	DB_server.Create(&user)
	return true
}

func CheckAuthzCode(code string) bool {
	temp := AuthzCodeInfo{}
	if err := DB_server.Where("auth_code = ?", code).First(&temp).Error; err == nil {
		expire_time, _ := strconv.Atoi(temp.ExpireTime)
		if time.Now().After(temp.CreatedAt.Add(time.Duration(expire_time) * time.Minute)) {
			DB_server.Delete(&temp)
			return false
		}else {
			return true
		}
	}
	return false
}

func FindAuthzInfoInAuthzCode(code string) (*AuthzCodeInfo, bool) {
	temp := AuthzCodeInfo{}
	if err := DB_server.Where("auth_code = ?", code).First(&temp).Error; err == nil {
		return &temp, true
	}
	return nil, false
}