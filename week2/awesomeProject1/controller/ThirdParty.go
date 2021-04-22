package controller

import (
	"awesomeproject1/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ThirdParty(app *gin.Engine) {
	app.GET("/HomePage", getAppHomePage)
	app.POST("/HomePage", jumpToAuth)
	app.PUT("/HomePage", getPictureInfo)
}

func getPictureInfo(c *gin.Context){
	//json := model.Picture{}
	//_ = c.BindJSON(&json)
	//Info.PictureInfo = json
}

func getAppHomePage(c *gin.Context) {
	username := c.Query("user_name")
	if model.IsValid(username) {
		c.HTML(http.StatusOK, "app-home-login.html", nil)
	}else {
		c.HTML(http.StatusOK, "app-home-notlogin.html", nil)
	}
}

func jumpToAuth(c *gin.Context) {
	var json map[string]interface{}
	_ = c.BindJSON(&json)
	log.Println(json["message"])
	path := "/server/authorization"
	c.Redirect(http.StatusTemporaryRedirect, path)
}
