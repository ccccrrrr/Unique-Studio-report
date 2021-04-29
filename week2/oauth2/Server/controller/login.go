package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oauth2/Server/model"
)

const ExpireTime = 20

func Login(server *gin.Engine) {
	server.GET("/login", getLoginPage)
	server.POST("/login", loginOperation)
}

func loginOperation(context *gin.Context) {
	info := model.UserInfo{}
	info.UserName = context.PostForm("username")
	info.UserPassword = context.PostForm("userpassword")
	if model.IsValid(info) {
		context.SetCookie("username", info.UserName, ExpireTime, "/", "localhost", false, true)
		context.HTML(http.StatusOK, "login.gohtml", gin.H{"message": 1})
	} else {
		context.HTML(http.StatusOK, "login.gohtml", gin.H{"message": 0})
	}
}

func getLoginPage(context *gin.Context) {
	_, returnValue := model.CheckCookie(context, "username")
	if returnValue == false {
		context.HTML(http.StatusOK, "login.gohtml", gin.H{"message": -1})
	} else {
		context.HTML(http.StatusOK, "login.gohtml", gin.H{"message": 2})
	}

}
