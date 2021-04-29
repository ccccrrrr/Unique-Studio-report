package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(app *gin.Engine) {
	app.GET("/home", getHomePage)
}

func getHomePage(context *gin.Context) {
	context.HTML(http.StatusOK, "home.html", nil)
}

