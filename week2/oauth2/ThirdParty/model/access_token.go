package model

import (
	model2 "oauth2/Server/model"
)

func InsertAccessTokenInfo(info model2.AccessTokenInfo) bool {
	appDB.Create(&info)
	return true
}

// GetNewestAccessTokenInfo just use latest access_token received from server
func GetNewestAccessTokenInfo() (*model2.AccessTokenInfo, error) {
	var temp model2.AccessTokenInfo
	if err := appDB.Table("access_token_infos").Last(&temp).Error; err != nil {
		return nil,err
	}
	return &temp, nil
}
