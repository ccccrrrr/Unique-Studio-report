package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	model2 "oauth2/Server/model"
	"oauth2/ThirdParty/model"
)

const RedirectUri = "http://localhost:9001/receive_token"

type AccessTokenInfo struct {
	gorm.Model
	UserName    string `json:"user_name" gorm:"column:user_name"`
	AccessToken string `json:"access_token" gorm:"column:access_token"`
	AuthzCode   string `json:"auth_code" gorm:"column:auth_code"`
	ExpireTime  string `json:"expire_time" gorm:"column:expire_time"`
	Scope       string `json:"scope" gorm:"column:scope"`
}

func AuthzCode(app *gin.Engine) {
	app.GET("/home/getAuth", getAuthPage)
	app.GET("/home/getAuth/success", getAuthCode)
}

func getAuthPage(c *gin.Context) {
	c.HTML(http.StatusOK, "getAuth.html", nil)
}

func getAuthCode(c *gin.Context) {
	code := c.Query("code")
	redirect_uri := RedirectUri
	path := "http://localhost:9090/accesstoken_endpoint" + "?code=" + code + "&redirect_uri=" + redirect_uri
	response, _ := http.Post(path, "application/json", nil)
	dec := json.NewDecoder(response.Body)
	accessTokenInfo := model2.AccessTokenInfo{}
	err := dec.Decode(&accessTokenInfo)
	if err != nil {
		log.Println(err)
		return
	}
	_ = model.InsertAccessTokenInfo(accessTokenInfo)
	c.HTML(http.StatusOK, "getAuth.html", nil)
}