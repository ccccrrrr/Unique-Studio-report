package model

import (
	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
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
