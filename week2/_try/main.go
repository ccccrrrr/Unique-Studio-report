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
	server.GET("/server/authorization", func(c *gin.Context){
		clientId := c.Query("client_id")
		code := c.Query("response_type")
		scope := "read-only"
		redirect_uri := c.Query("redirect_uri")
		if len(c.Query("scope")) == 0 {
			scope = c.Query("scope")
		}
		db.Update(&AuthRequest{
			ResponseType: code,
			ClientId:     clientId,
			RedirectUri:  redirect_uri,
			Scope:        scope,
		})
		c.Redirect(http.StatusPermanentRedirect, "/server/" + redirect_uri + "?code=" + code)
	})
	//g1 := server.Group("/server/" + redirect_uri)
	path := ""
	json := ClientInfo{}
		//through json, get clientID, clientPassword


	// in server redirect page, get code
	// it needs to add a verifying operation
	server.GET("/server/redirect", func(c *gin.Context){
		getCode := c.Query("code")
		fmt.Println(getCode)
		// store code in db1?
	})



	// APP send JSON format information to service
	server.POST("/server/redirect", func(c * gin.Context){
		err = c.BindJSON(&json)
			fmt.Printf("%+v\n", &json)
			//path := "/server/oauth/token?client_id=" + json.ClientId +
			//	"&client_secret=" + json.ClientPassword + "&grant_type=" + json.GrantType +
			//	"&code=" + json.Code + "&redirect_uri=" + json.RedirectUri

			AuthorizationCode := "ccccrrrr"
			// through authorization, issue an authorization code to app
			c.Redirect(http.StatusPermanentRedirect, "http://localhost:9090/app/getresource?code=" + AuthorizationCode)
			c.JSON(http.StatusOK, gin.H{
				"authorizationCode": "200",
			})
			//c.Redirect(http.StatusPermanentRedirect, path)
		})
	app := server.Group("/app")

	// app receive authorization code
	app.GET("/getresource", func(c *gin.Context){

	})
	server.POST(path, func(c *gin.Context){
				c.JSON(http.StatusOK, json)
		})
	//4. APP display the authorization page
	//5. user input ID and password and choose life span
	//6. service issue an authorization code to App
	//7. App present authorization code
	//8. service issue an access token
	//9.

	_ = server.Run(":9090")
}
