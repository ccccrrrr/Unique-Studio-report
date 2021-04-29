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

const Authz_RedirectUri = "http://localhost:9001/homepage/getAuth/success"
const AuthzUrl = "http://localhost:9090/auth-and-login"

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
	//	app.POST("/home/getAuth", sendRequest)
	//	app.PUT("/home/getAuth", redirectLoginPage)
	app.GET("/home/getAuth/success", getAuthCode)
}

func getAuthPage(c *gin.Context) {
	c.HTML(http.StatusOK, "getAuth.html", nil)
}

func getAuthCode(c *gin.Context) {
	code := c.Query("code")
	redirect_uri := "http://localhost:9001/receive_token"
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
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

//
//func getAuthCode(context *gin.Context) {
//	code := context.Query("code")
//	username := context.PostForm("username")
//	if code != "" && username != "" {
//		path := "http://localhost:9090/accesstoken_endpoint"
//		path += "?code=" + code
//		path += "&redirect_uri=" + "http://localhost:9001/home/access_token"
//		_, _ = http.Post(path, "application/x-www-form-urlencoded", nil)
//		model.InsertAuthCodeInfo(model.AuthCodeInfo{AuthzCode: code, UserName: username})
//	} else {
//		log.Println("get auth error")
//	}
//	context.HTML(http.StatusOK, "getauth-success.html", nil)
////}
//
//func redirectLoginPage(context *gin.Context) {
//	context.Redirect(http.StatusTemporaryRedirect, "http://localhost:9090/auth-and-login")
//}
//
//func sendRequest(context *gin.Context) {
//	user_name := context.PostForm("username")
//	scope := context.PostForm("scope")
//	redirect_uri := Authz_RedirectUri
//	expire := context.PostForm("expire")
//	path := AuthzUrl
//	path += "?user_name=" + user_name
//	path += "&scope=" + scope
//	path += "&expire=" + expire
//	path += "&redirect_uri=" + redirect_uri
//	_, _ = http.Get(path)
//}
//
//func getAuthPage(context *gin.Context) {
//	context.HTML(http.StatusOK, "home-getAuth.html", nil)
//}
