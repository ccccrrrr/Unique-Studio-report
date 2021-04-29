package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oauth2/Server/model"
)

func Register(server *gin.Engine) {
	server.GET("/register", getRegisterPage)
	server.POST("/register", registerOperation)
}

func registerOperation(context *gin.Context) {
	info := model.UserInfo{}
	info.UserName = context.PostForm("username")
	info.UserPassword = context.PostForm("userpassword")
	if model.IsExistUserName(info.UserName) {
		context.HTML(http.StatusOK, "register.gohtml", gin.H{"message": 0})
	} else {
		model.CreateUser(info)
		context.HTML(http.StatusOK, "register.gohtml", gin.H{"message": 1})
	}
}

func getRegisterPage(context *gin.Context) {
	context.HTML(http.StatusOK, "register.gohtml", nil)
}
