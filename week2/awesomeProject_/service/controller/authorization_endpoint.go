package controller

import (
	model2 "awesomeproject1/service/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type sendAccessToken struct {
	AccessToken string
	ExpireIn    time.Time
	Scope       string
//	ScopeNumber int
}

var letters = []byte("abcdefghjkmnpqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~")

func init() {
	rand.Seed(time.Now().Unix())
}

//type scope struct {
//	Scope       string    `json:"scope"`
//	TimeSpan    time.Time `json:"time_span"`
//	PresentTime time.Time `json:"present_time"`
//}

func Authorization(server *gin.Engine) {
	server.GET("/server", getAuthorizationPage)
	server.POST("/server/authorization", _Authorization)
	server.GET("/server/authorization", getAuthorization)
	server.GET("/server/choose_scope", ChooseScope)
	server.GET("/server/authorization_endpoint", getAuthorizationEndpoint)
	server.POST("/server/authorization_endpoint", generateAuthCode)
	server.GET("/server/redirect", getClientInfo)
	server.GET("/server/sendAuthCode", sendAuthCode)
	server.POST("/server/token_endpoint", getAuthCode)
	// step 7 and 8
	server.GET("/server/token_endpoint", issueAccessToken)
}

func generateAuthCode(c *gin.Context) {
	// need something
	//responseType := c.Query("responsetype")
	//if responseType != "200" {
	//	c.Redirect(http.StatusMovedPermanently, "localhost:9090/HomePage")
	//}
	client_id := c.Query("client_id")
	scope := c.Query("scope")
	expire := c.Query("expire")
	_, err := c.Request.Cookie("username")
	log.Println(err)
	if err == nil {
//	if isLogin(client_id, c) { // use cookie
		//wordList, _ := generate.GenRandomMix(10)
		//code := wordList.Chance
		// store authzcode in db
		code := model2.GenerateTempAuthCode(client_id)
		model2.InsertAuthCodeInfo(model2.AuthInfo{
			client_id,
			code,
			scope,
			expire,
			time.Now(),
		})
		model2.AuthCodeToUser(client_id, scope, expire, code)
		path := "http://localhost:9090/HomePage?code="
		path += code
		_, _ = http.Post(path, "application/x-www-form-urlencoded", nil)
	//	c.Redirect(http.StatusMovedPermanently, path)
	} else {
		// you have to login first
		c.Redirect(http.StatusMovedPermanently, "http://localhost:9090/login")
	}

	//redirect_uri := c.Query("redirect_uri")

}

func getAuthorizationPage(c *gin.Context) {

}

func getAuthorization(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "OK",
	//})
	_json := map[string]string{}
	_ = c.BindJSON(&_json)
	if model2.CheckState(_json["user_name"]) {
		c.Redirect(http.StatusPermanentRedirect, "/server/choose_scope")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}

func ChooseScope(c *gin.Context) {
	//	json := scope{}
	//	_ = c.BindJSON(&json)
	//	c.Redirect()
}

func _Authorization(c *gin.Context) {
	// need the name of the user
	json := map[string]string{}
	_ = c.BindJSON(&json)
	if model2.CheckState(json["user_name"]) {
		c.Redirect(http.StatusPermanentRedirect, "/server/choose_scope")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}

func issueAccessToken(c *gin.Context) {
	authCode := c.Query("authorization_code")
	accessToken := generateAccessToken(authCode)
	c.JSON(http.StatusOK, gin.H{
		"access token": accessToken,
	})
}

func getAuthCode(c *gin.Context) {
	//grantType := c.PostForm("grant_type")
	//authorizationCode := c.PostForm("code")
	//redirectUri := c.PostForm("redirect_uri")
	//codeVerifier := c.PostForm("code_verifier")
//	log.Println("token_endpoint get auth code")
	code := c.Query("code")
//	log.Println("code = " + code)
	//	user_name, _ := c.Request.Cookie("username")
	accessTokenInfo := model2.GenerateAccessToken(code, 50)
	log.Println("accesstoken" + accessTokenInfo.AccessToken)
	var data = sendAccessToken{
		accessTokenInfo.AccessToken,
		accessTokenInfo.ExpireTime,
		accessTokenInfo.Scope,
		//generateChoice(accessTokenInfo.Scope),
	}
	body, _ := json.Marshal(data)
	log.Println(body)
//	c.Redirect(http.StatusPermanentRedirect)
	_, _ = http.Post("http://localhost:9090/HomePage?access_token=" + accessTokenInfo.AccessToken +
		"&expireIn=" + accessTokenInfo.ExpireTime.String() + "&scope=" + accessTokenInfo.Scope,
		"application/json;charset=UTF-8",
		nil)
}

func sendAuthCode(c *gin.Context) {
	loginId := c.Query("loginId")
	password := c.Query("password")
	authCode := generateAutoCode(loginId, password)
	c.JSON(http.StatusOK, gin.H{
		"authorizationCode": authCode,
	})
}

func getClientInfo(c *gin.Context) {
	json := model2.RequestPermission{}
	_ = c.BindJSON(&json)
	loginId := json.LoginId
	password := json.Password
	if Check(loginId, password) {
		//authorizationCode := generateAutoCode()
		//c.JSON(200, gin.H{
		//	"message": "need authorization code",
		//})
		c.Redirect(http.StatusTemporaryRedirect, "/server/sendAuthCode?loginId="+loginId+"&password="+password)
	} else {
		c.String(200, "wrong password or username\n")
	}
}

func getAuthorizationEndpoint(c *gin.Context) {
	// need something
	//responseType := c.Query("responsetype")
	//if responseType != "200" {
	//	c.Redirect(http.StatusMovedPermanently, "localhost:9090/HomePage")
	//}
	client_id := c.Query("client_id")
	scope := c.Query("scope")
	expire := c.Query("expire")
	code := model2.GenerateTempAuthCode(client_id)
	log.Println(code)
	if isLogin(client_id, c) { // use cookie
		//wordList, _ := generate.GenRandomMix(10)
		//code := wordList.Chance
		// store authzcode in db
		model2.AuthCodeToUser(client_id, scope, expire, code)
		path := "http://localhost:9090/HomePage?code="
		path += code
		log.Println(path)
		c.Redirect(http.StatusMovedPermanently, path)
	} else {
		// you have to login first
		c.Redirect(http.StatusMovedPermanently, "http://localhost:9090/login")
	}

	//redirect_uri := c.Query("redirect_uri")

	//c.JSON(200, gin.H{
	//	"message": "...",
	//})
}

func _getAuthorizationEndpoint(c *gin.Context) {

	// Tells the authorization server which grant to execute
	//responseType := c.Query("response_type")

	// clientID: public identifier
	// github: 6779ef20e75817b79602
	//clientId := c.Query("client_id")

	// holds a URL. A successful response from this endpoint results in a redirect to this URL
	redirectUri := c.Query("redirect_uri")

	// A space-delimited list of permissions that the application requires.
	// changes are needed later
	// scope := c.Query("scope")

	// An opaque value, used for security purposes. If this request parameter is set in the request, then it is returned to the application as part of the redirect_uri.
	// I don't know how to solve this problem
	state := c.Query("state")

	//codeChallenge := c.Query("code_challenge")
	var codeChallenge string
	codeChallengeMethod := c.Query("code_challenge_method")
	codeVerifier := generateCodeVerifier()
	if codeChallengeMethod == "plain" {
		codeChallenge = codeVerifier
	} else {
		// changes are needed
		// it might cost much time to write S256 method, so I assume codeChallengeMethod always be plain
		codeChallenge = codeVerifier
	}
	// temporary
	AuthorizationCode := codeChallenge
	path := "/server/" + redirectUri + "?code=" + AuthorizationCode + "&state=" + state
	fmt.Println(path)
	//c.Redirect(302, "http://localhost:9090/server/redirect")
	c.Redirect(302, path)
}

func generateCodeVerifier() string {
	length := rand.Intn(86) + 43
	var codeVerifier string
	for i := 0; i < length; i++ {
		codeVerifier += string(letters[rand.Intn(43)])
	}
	return codeVerifier
}

func Check(loginId string, password string) bool {
	return true
}

func generateAccessToken(authCode string) string {
	return "kkk"
}

func generateAutoCode(loginId string, password string) string {
	return "ccccrrrr"
}

func generateChoice(s string) int {
	if s == "read" {
		return 1
	}else if s == "read-write" {
		return 2
	}else{return 0}
}