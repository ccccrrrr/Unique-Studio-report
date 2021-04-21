package main

import (
	"fmt"
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type ClientInfo struct {
	ClientId string	`json:"client_id"`
	ClientPassword string `json:"client_password"`
	GrantType string `json:"grant_type"`
	Code string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
}

type AuthRequest struct{
	gorm.Model
	ResponseType string
	ClientId 	string
	RedirectUri string
	Scope string
//	State string
//	CodeChallenge string
//	CodeChallengeMethod string
}

func main() {

	server := gin.Default()

	db, err := gorm.Open("mysql", "root: @(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//set form
	db.CreateTable(&AuthRequest{})
	defer db.Close()

	//1. Request from User to App to link to ServiceABC

	//2. App make authorization request to serviceABC

	// 3. serviceABC return an authorization page to APP
	// a request was sent from APP to server
	// and the server receive it and redirect to redirect page with code
	//server.GET("/server/authorization", func(c *gin.Context){
	//	clientId := c.Query("client_id")
	//	code := c.Query("response_type")
	//	scope := "read-only"
	//	redirect_uri := c.Query("redirect_uri")
	//	if len(c.Query("scope")) == 0 {
	//		scope = c.Query("scope")
	//	}
	//	db.Update(&AuthRequest{
	//		ResponseType: code,
	//		ClientId:     clientId,
	//		RedirectUri:  redirect_uri,
	//		Scope:        scope,
	//	})
	//	c.Redirect(http.StatusPermanentRedirect, "/server/" + redirect_uri + "?code=" + code)
	//})
	//g1 := server.Group("/server/" + redirect_uri)
	//path := ""
	//json := ClientInfo{}
		//through json, get clientID, clientPassword


	// in server redirect page, get code
	// it needs to add a verifying operation
	//server.GET("/server/redirect", func(c *gin.Context){
	//	getCode := c.Query("code")
	//	fmt.Println(getCode)
	//	// store code in db1?
	//})



	server.GET("/server/authorization_endpoint", func(c *gin.Context){
		//responseType := c.Query("response_type")
		clientId := c.Query("client_id")
		redirectUri := c.Query("redirect_uri")
		//scope := c.Query("scope")
		state := c.Query("state")
		//codeChallenge := c.Query("code_challenge")
		//codeChallengeMethod := c.Query("code_challenge_method")

		// temporary
		AuthorizationCode := clientId
		path := "http://localhost:9090/server/" + redirectUri + "?code=" + AuthorizationCode + "&state=" + state
		c.Redirect(302, path)
	})

	// server receive code and issue an access token
	server.POST("/server/token_endpoint", func(c *gin.Context){
		//grantType := c.PostForm("grant_type")
		//authorizationCode := c.PostForm("code")
		//redirectUri := c.PostForm("redirect_uri")
		//codeVerifier := c.PostForm("code_verifier")
		c.JSON(http.StatusOK, gin.H{
			"access_token": "",
			"token_type": "",
			"expires_in": "",
			"refresh_token":"",
			"scope": "",
		})
	})
	//4. APP display the authorization page
	//5. user input ID and password and choose life span
	//6. service issue an authorization code to App
	//7. App present authorization code
	//8. service issue an access token
	//9.

	_ = server.Run(":9090")
}
