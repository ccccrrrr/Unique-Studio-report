package controller

import (
	"awesomeproject1/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

var letters = []byte("abcdefghjkmnpqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~")

func init() {
	rand.Seed(time.Now().Unix())
}

type scope struct {
	Scope       string    `json:"scope"`
	TimeSpan    time.Time `json:"time_span"`
	PresentTime time.Time `json:"present_time"`
}

func Authorization(server *gin.Engine) {
	server.POST("/server/authorization", _Authorization)
	server.GET("/server/authorization", getAuthorization)
	server.GET("/server/choose_scope", ChooseScope)
	server.GET("/server/authorization_endpoint", getAuthorizationEndpoint)
	server.GET("/server/redirect", getClientInfo)
	server.GET("/server/sendAuthCode", sendAuthCode)
	server.POST("/server/token_endpoint", getAuthCode)
	// step 7 and 8
	server.GET("/server/token_endpoint", issueAccessToken)
}

func getAuthorization(c *gin.Context){
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "OK",
	//})
	json := map[string]string{}
	_ = c.BindJSON(&json)
	if model.CheckState(json["user_name"]) {
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
	if model.CheckState(json["user_name"]) {
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
	c.JSON(http.StatusOK, gin.H{
		"access_token":  "",
		"token_type":    "",
		"expires_in":    "",
		"refresh_token": "",
		"scope":         "",
	})
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
	json := model.RequestPermission{}
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
