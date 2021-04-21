package main

import (
	"awesomeProject1/controller"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ClientInfo struct {
	ClientId       string `json:"client_id"`
	ClientPassword string `json:"client_password"`
	GrantType      string `json:"grant_type"`
	Code           string `json:"code"`
	RedirectUri    string `json:"redirect_uri"`
}

type AuthRequest struct {
	gorm.Model
	ResponseType string
	ClientId     string
	RedirectUri  string
	Scope        string
	//	State string
	//	CodeChallenge string
	//	CodeChallengeMethod string
}
type RequestPermission struct {
	LoginId  string `json:"login_id"`
	Password string `json:"password"`
	// tbc
}


func generateAutoCode(loginId string, password string) string {
	return "ccccrrrr"
}

func generateAccessToken(authCode string) string {
	return "kkk"
}
func main() {

	server := gin.Default()

	model.StartDatabase()

	//db, err := gorm.Open("mysql", "root: @(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")

	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	// set form
	//db.CreateTable(&AuthRequest{})
	//db.CreateTable(&ClientInfo{})
	//db.AutoMigrate(&model.User{})
	//defer db.Close()

	// need client id verify
	controller.Authorization_endpoint(server)
	controller.Login(server)
/*	server.GET("/server/authorization_endpoint", func(c *gin.Context) {

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
	})
*/
	// server redirect page
	// check requested permission
	// get json format information
/*	server.GET("/server/redirect", func(c *gin.Context) {
		json := RequestPermission{}
		_ = c.BindJSON(&json)
		loginId := json.LoginId
		password := json.Password
		if controller.Check(loginId, password) {
			//authorizationCode := generateAutoCode()
			//c.JSON(200, gin.H{
			//	"message": "need authorization code",
			//})
			c.Redirect(http.StatusTemporaryRedirect, "/server/sendAuthCode?loginId="+loginId+"&password="+password)
		} else {
			c.String(200, "wrong password or username\n")
		}
	})
*/
	// step 6 issue a short-lived authorization code
	// but I can just use restClient to get the authorization code ;(
/*	server.GET("/server/sendAuthCode", func(c *gin.Context) {
		loginId := c.Query("loginId")
		password := c.Query("password")
		authCode := generateAutoCode(loginId, password)
		c.JSON(http.StatusOK, gin.H{
			"authorizationCode": authCode,
		})
	})
*/
	// step 7 also in restClient see requestToTokenEndpoint.http
	// server receive code and issue an access token
/*	server.POST("/server/token_endpoint", func(c *gin.Context) {
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
	})
*/
	// step 8 issue an access token to app
	// need an authorization_code
/*	server.GET("/server/token_endpoint", func(c *gin.Context) {
		authCode := c.Query("authorization_code")
		accessToken := generateAccessToken(authCode)
		c.JSON(http.StatusOK, gin.H{
			"access token": accessToken,
		})
	})
*/
	_ = server.Run(":9090")
}
