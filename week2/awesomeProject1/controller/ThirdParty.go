package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ThirdParty(app *gin.Engine) {
	app.GET("/HomePage", getAppHomePage)
	app.POST("/HomePage", jumpToAuth)
}

func getAppHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "app-home.html", nil)
}

func jumpToAuth(c *gin.Context) {
	var json map[string]interface{}
	_ = c.BindJSON(&json)
	log.Println(json["message"])
	path := "/server/authorization"
	c.Redirect(http.StatusTemporaryRedirect, path)
}
