package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	model2 "oauth2/Server/model"
	"oauth2/ThirdParty/model"
)

const AccessToken_RedirectUri = "http://localhost:9001/homepage/getToken"
const ExpireTime = 20

func AccessToken(app *gin.Engine) {
	app.POST("homepage/getToken", getToken)
}

func getToken(context *gin.Context) {
	accessTokenInfo := model2.AccessTokenInfo{}
	_ = context.BindJSON(&accessTokenInfo)
	err := model.InsertAccessTokenInfo(accessTokenInfo)
	if err == false {
		log.Println(err)
		return
	}
	model.InsertAccessTokenInfo(accessTokenInfo)
	log.Println("successfully get token")
	context.JSON(http.StatusOK, gin.H{"message": "successfully get token"})
}

