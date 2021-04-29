package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"oauth2/Server/model"
)

type AccessTokenInfo struct {
	gorm.Model
	UserName    string `json:"user_name" gorm:"column:user_name"`
	AccessToken string `json:"access_token" gorm:"column:access_token"`
	AuthzCode   string `json:"auth_code" gorm:"column:auth_code"`
	ExpireTime  string `json:"expire_time" gorm:"column:expire_time"`
	Scope       string `json:"scope" gorm:"column:scope"`
}

func AccessTokenEndpoint(server *gin.Engine) {
	server.POST("/accesstoken_endpoint", getAuthCode)
}

func getAuthCode(c *gin.Context) {
	code := c.Query("code")
	redirect_uri := c.Query("redirect_uri")
	accessTokenInfo, returnType := model.GenerateAccessToken(code, redirect_uri)
	if returnType == false {
		log.Println("get auth code error")
		c.JSON(http.StatusInternalServerError, nil)
	}else {
		model.InsertAccessTokenInfo(accessTokenInfo)
//		_, _ = http.Post()
		c.JSON(http.StatusOK, accessTokenInfo)
	}
}
