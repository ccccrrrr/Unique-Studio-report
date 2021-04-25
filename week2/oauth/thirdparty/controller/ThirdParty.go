package controller

import (
	"awesomeproject1/thirdparty/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type AccessTokenInfo struct {
	ExpireTime  time.Time `json:"expire_time"`
	UserName    string    `json:"user_name"`
	AccessToken string    `json:"access_token"`
	Scope       string    `json:"scope"`
	ScopeNumber int       `json:"scope_number"`
}

var (
	access_token string
	accessTokenInfo AccessTokenInfo
)

func ThirdParty(app *gin.Engine) {
	app.GET("/HomePage", getAppHomePage)
	app.PUT("/HomePage", getAccessToken)
	app.POST("/HomePage", PictureOperation)
	app.GET("/HomePage/getAuth", getAuthPage)
	app.POST("/HomePage/getAuth", sendAuthzInfo)
	app.PUT("/HomePage/getAuth", getTokenInfo)
}

func getTokenInfo(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &accessTokenInfo)
	access_token = accessTokenInfo.AccessToken
}

func getAuthPage(c *gin.Context) {
	authCode := c.Query("code")
	if authCode != "" {
		path := "http://localhost:9090/server/token_endpoint?"
		path += "grant_type=" + authCode
		_, _ = http.Get(path)
	}
	c.HTML(http.StatusOK, "app-getAuth.html", nil)
}

func sendAuthzInfo(c *gin.Context) {
	_client_id := c.PostForm("client_id")
	scope := c.PostForm("scope")
	expire := c.PostForm("expire")
	path := "http://localhost:9090/server/authorization_endpoint" +
		"?client_id=" + _client_id +
		"&scope=" + scope +
		"&expire=" + expire
	_, _ = http.Get(path)
	c.HTML(http.StatusOK, "app-getAuth.html", nil)
}

func getAccessToken(c *gin.Context) {
	// two choice
	var info AccessTokenInfo
	_ = c.BindJSON(&info)
//	_verify = info
}

func PictureOperation(c *gin.Context) {
	if time.Now().After(accessTokenInfo.ExpireTime) {
		accessTokenInfo = AccessTokenInfo{}
		return
	}
	operation := c.PostForm("submit")
	if operation == "delete" {
		picturePath := c.PostForm("picture_path")
		model.SendDelete(picturePath, access_token)
	} else if operation == "show" {
		model.SendShow(access_token)
	} else if operation == "add" {
		picturePath := c.PostForm("picture_path")
		model.SendAdd(picturePath, access_token)
	} else {

	}
}

func getAppHomePage(c *gin.Context) {
//	code := c.Query("code")
	if time.Now().After(accessTokenInfo.ExpireTime) {
		accessTokenInfo = AccessTokenInfo{}
//		return
	}
//	if code == "" {
		c.HTML(http.StatusOK, "app.gohtml", accessTokenInfo)
//	} else {

//	}
}
