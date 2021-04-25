package controller

import (
	"awesomeproject1/service/model"
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type sendAccessToken struct {
	AccessToken string
	ExpireIn    time.Time
	Scope       string
	//	ScopeNumber int
}

var letters = []byte("abcdefghjkmnpqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~")

func Authorization(server *gin.Engine) {
	server.GET("/server/authorization_endpoint", getAuthorizationEndpoint)
	server.POST("/server/authorization_endpoint", generateAuthCode)
	server.GET("/server/sendAuthCode", sendAuthCode)
	server.POST("/server/token_endpoint", getAuthCode)
	server.GET("/server/token_endpoint", issueAccessToken)
}

func generateAuthCode(c *gin.Context) {
	client_id := c.Query("client_id")
	scope := c.Query("scope")
	expire := c.Query("expire")
	_, err := c.Request.Cookie("username")
	log.Println(err)
	if err == nil {
		code := model.GenerateTempAuthCode(client_id)
		model.InsertAuthCodeInfo(model.AuthInfo{
			UserName : client_id,
			Code: code,
			Scope: scope,
			ExpireTime: expire,
			TimeNow: time.Now(),
		})
		model.AuthCodeToUser(client_id, scope, expire, code)
		path := "http://localhost:9090/HomePage?code="
		path += code
		_, _ = http.Post(path, "application/x-www-form-urlencoded", nil)
	} else {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:9090/login")
	}
}

func issueAccessToken(c *gin.Context) {
	authCode := c.Query("grant_type")
	detail, _ := model.FindInfoThroughAuthCode(authCode)
	log.Printf("authCode:%v", authCode)
	expire_time, _ := strconv.Atoi(detail.ExpireTime)
	accessTokenInfo := model.GenerateAccessToken(authCode, expire_time)
	accessTokenInfo.ScopeNumber = model.GenerateScopeNumber(accessTokenInfo.Scope)
	bytes, _ := json.Marshal(accessTokenInfo)
	reqs, _ := http.NewRequest(http.MethodPut, "http://localhost:9001/HomePage/getAuth", bytes2.NewReader(bytes))
	client := http.Client{}
	_, _ = client.Do(reqs)
}

func getAuthCode(c *gin.Context) {
	code := c.Query("code")
	accessTokenInfo := model.GenerateAccessToken(code, 50)
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
	_, _ = http.Post("http://localhost:9090/HomePage?access_token="+accessTokenInfo.AccessToken+
		"&expireIn="+accessTokenInfo.ExpireTime.String()+"&scope="+accessTokenInfo.Scope,
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

func getAuthorizationEndpoint(c *gin.Context) {
	client_id := c.Query("client_id")
	scope := c.Query("scope")
	expire := c.Query("expire")
	code := model.GenerateTempAuthCode(client_id)
	//	log.Println(code)
	log.Printf("username: %v", Username)
	if Username == "" {
		_, _ = http.Get("http://localhost:9090/login")
	}
	log.Printf("username: %v", Username)
	model.AuthCodeToUser(client_id, scope, expire, code)
	path := "http://localhost:9001/HomePage/getAuth?code="
	path += code
	_, _ = http.Get(path)
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

func generateAutoCode(loginId string, password string) string {
	return "ccccrrrr"
}

func generateChoice(s string) int {
	if s == "read" {
		return 1
	} else if s == "read-write" {
		return 2
	} else {
		return 0
	}
}
