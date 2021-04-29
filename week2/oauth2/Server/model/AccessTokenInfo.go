package model

import (
	"errors"
	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type AccessTokenInfo struct {
	gorm.Model
	UserName    string `json:"user_name" gorm:"column:user_name"`
	AccessToken string `json:"access_token" gorm:"column:access_token"`
	AuthzCode   string `json:"auth_code" gorm:"column:auth_code"`
	ExpireTime  string `json:"expire_time" gorm:"column:expire_time"`
	Scope       string `json:"scope" gorm:"column:scope"`
	RedirectUri string `json:"redirect_uri" gorm:"column:redirect_uri"`
}

func GenerateAccessToken(authzCode string, redirect_uri string) (*AccessTokenInfo, bool) {
	authzCodeInfo, returnType := FindAuthzInfoInAuthzCode(authzCode)
	if returnType == false {
		return nil, false
	}
	accessTokenInfo := AccessTokenInfo{}
	accessTokenInfo.AuthzCode = authzCode
	accessTokenInfo.Scope = authzCodeInfo.Scope
	accessTokenInfo.RedirectUri = authzCodeInfo.RedirectUri
	accessTokenInfo.ExpireTime = authzCodeInfo.ExpireTime
	accessTokenInfo.UserName = authzCodeInfo.UserName
	accessTokenInfo.AccessToken = randstr.Hex(20)
	return &accessTokenInfo, true
}

func InsertAccessTokenInfo(info *AccessTokenInfo) bool {
	DB_server.Create(info)
	return true
}

func IsValidToken(access_token string) (*AccessTokenInfo, error) {
	var temp *AccessTokenInfo
	if err := DB_server.Where("access_token = ?", access_token).First(&temp).Error; err != nil {
		return nil, err
	}
	expire_time, _ := strconv.Atoi(temp.ExpireTime)
	if time.Now().After(temp.CreatedAt.Add(time.Minute * time.Duration(expire_time))) {
		return nil, errors.New("access token has expired")
	}
	return temp, nil
}

func GetInfoInAccessToken(access_token string) (*AccessTokenInfo, error) {
	var temp * AccessTokenInfo
	if err := DB_server.Where("access_token = ?", access_token).First(&temp).Error; err != nil {
		return nil, err
	}
	return temp, nil
}

func IsValidDeleteOperation(access_token string) bool {
	info, err := GetInfoInAccessToken(access_token)
	if err != nil {
		log.Println(err)
		return false
	}

	if info.Scope == "read-write" {
		return true
	} else {
		return false
	}
}

func IsValidUploadOperation(access_token string) bool {
	info, err := GetInfoInAccessToken(access_token)
	if err != nil {
		log.Println(err)
		return false
	}

	if info.Scope == "read-write" {
		return true
	} else {
		return false
	}
}

func IsValidGetOperation(access_token string) bool {
	info, err := GetInfoInAccessToken(access_token)
	if err != nil {
		log.Println(err)
		return false
	}

	if info.Scope == "read-write" || info.Scope == "read" {
		return true
	} else {
		return false
	}
}