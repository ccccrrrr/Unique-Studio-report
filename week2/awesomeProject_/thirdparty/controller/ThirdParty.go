package controller

import (
	"awesomeproject1/thirdparty/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)
var (
	_verify AccessTokenInfo
	access_token string
	client_id string
)

type AccessTokenInfo struct {
	ExpireTime  time.Time `json:"expire_time"`
	UserName    string    `json:"user_name"`
	AccessToken string    `json:"access_token"`
	Scope       string    `json:"scope"`
	ScopeNumber int       `json:"scope_number"`
}

func ThirdParty(app *gin.Engine) {
	app.GET("/HomePage", getAppHomePage)
//	app.DELETE("/HomePage", deletePicture)
	app.PUT("/HomePage", getAccessToken)
	app.POST("/HomePage", PictureOperation)
	app.GET("/HomePage/getAuth", getAuthPage)
	app.POST("/HomePage/getAuth", sendAuthzInfo)
}

func getAuthPage(c *gin.Context){
	c.HTML(http.StatusOK, "app-getAuth.html", nil)
}
func sendAuthzInfo(c *gin.Context){
	_client_id := c.PostForm("client_id")
	scope := c.PostForm("scope")
	expire := c.PostForm("expire")
	path := "http://localhost:9090/server/authorization_endpoint" +
	"?client_id=" + _client_id +
	"&scope=" + scope +
	"&expire=" + expire
	log.Println("get path: " + path)
	_, _ = http.Get(path)
//	http.Redirect(c.Writer, c.Request, "http://localhost:9001/HomePage", http.StatusPermanentRedirect)
//	c.Redirect(http.StatusPermanentRedirect, "/HomePage")
	c.HTML(http.StatusOK, "app-getAuth.html", nil)
}

func getAccessToken(c *gin.Context){
	// two choice
	var info AccessTokenInfo
	_ = c.BindJSON(&info)
	_verify = info
}
//
//func receiveInfo(c *gin.Context) {
//	clientId := c.PostForm("client_id")
//	scope := c.PostForm("scope")
//	expire := c.PostForm("expire")
//	var info model2.GetInfo1
//	//	_ = c.BindJSON(&info)
//	info.TimeLimit = expire
//	info.UseName = clientId
//	info.Scope = scope
//	_info := model2.GenerateInfo(info, c)
//	path := "http://localhost:9090/server/authorization_endpoint?"
//	//	path += "responsetype=200"
//	path += "client_id=" + _info.ClientId
//	//	path += "&redirect_uri=" + _info.RedirectUri
//	path += "&scope=" + _info.Scope
//	path += "&expire=" + _info.TimeLimit
//	//	log.Println("start redirect...")
//	_, err := http.Post(path, "application/x-www-form-urlencoded", nil)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	//	c.Redirect(http.StatusPermanentRedirect, path)
//}
//
//func getgetAuthPage(c *gin.Context) {
//	c.HTML(http.StatusOK, "app-getAuth.html", nil)
//}
//
//func deletePicture(c *gin.Context) {
//	//log.Println(c.Query("user_name"))
//	info := model2.DeleteInfo{}
//	_ = c.BindJSON(&info)
//	info.DeleteUser = user_name
//	if model2.IsValid(user_name) {
//		//if model.IsValid(info.DeleteUser) {
//		res := model2.DeletePicture(info)
//		log.Println(info)
//		if res == true {
//			log.Println("delete successfully")
//		} else {
//			log.Println("unsuccessful delete")
//		}
//	} else {
//		log.Println("invalid user name")
//	}
//}

func PictureOperation(c *gin.Context) {
	http.Post("http://localhost:9090/server/isValidToken?token=" + _verify.AccessToken,"text", nil)
	if _verify.AccessToken == ""{
		log.Println("this is not a valid access token!")
		return
	}
	operation := c.PostForm("submit")
	if operation == "delete" {
		picturePath := c.PostForm("picture_path")
		model.SendDelete(picturePath, _verify.AccessToken)
	}else if operation == "show" {
		model.SendShow(_verify.AccessToken)
	}else if operation == "add" {
		picturePath := c.PostForm("picture_path")
		model.SendAdd(picturePath, _verify.AccessToken)
	}else {

	}
	//if c.ContentType() == "application/json;charset=UTF-8" {
	//	_verify.ExpireTime, _ = time.Parse(time.RFC3339, c.Query("expireIn"))
	//	_verify.Scope = c.Query("scope")
	//	_verify.AccessToken = c.Query("access_token")
	//	//_body, _ := ioutil.ReadAll(c.Request.Body)
	//	//_ = json.Unmarshal(_body, &__body)
	//	//_verify.ExpireTime = __body.ExpireIn
	//	//_verify.Scope = __body.Scope
	//	//_verify.AccessToken = __body.AccessToken
	//	temp, _ := c.Request.Cookie("username")
	//	log.Println("thirparty.go")
	//	_verify.UserName = temp.Value
	//	_verify.ScopeNumber = generateChoice(_verify.Scope)
	//
	//}
	//info := model2.Picture{}
	//// I think the picture should not store the userpassword, right?
	//_ = c.BindJSON(&info)
	//log.Println(info)
	//info.CreateUserName = user_name
	//if model2.IsValid(user_name) {
	//	//if model.IsValid(info.CreateUserName) {
	//	res := model2.Upload("", info)
	//	if res == true {
	//		log.Println("upload successfully")
	//	} else {
	//		log.Println("unsuccessful upload operation")
	//	}
	//} else {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "fail...",
	//	})
	//}
}

func getAppHomePage(c *gin.Context) {
	//username := c.Query("user_name")
	//user_name = username
	//log.Println(user_name)
	//if model.IsValid(username) {
	//	c.HTML(http.StatusOK, "app-home-login.html", nil)
	//}else {
	//	c.HTML(http.StatusOK, "app-home-notlogin.html", nil)
	//}

	code := c.Query("code")
	if code == "" {
		//c.HTML(http.StatusOK, "app.gohtml", _verify)
		c.HTML(http.StatusOK, "app.gohtml", gin.H{
			"ScopeNumber": 0,
		})
	} else {
		//var body = strings.NewReader("code=" + code)
		//_, _ = http.Post(
		//	"http://localhost:9090/server/token_endpoint?code="+code,
		//	"application/x-www-form-urlencoded",
		//	body)
		//		log.Println(err)
		//http.Redirect(c.Writer, c.Request, "http://localhost:9090/server/token_endpoint", http.StatusMovedPermanently)
	}
}
