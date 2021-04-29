package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oauth2/Server/model"
)

func AuthAndLogin(server *gin.Engine) {
	server.GET("/auth-and-login", getAuthAndLoginPage)
	server.POST("/auth-and-login", VerifyInfo)
}

func VerifyInfo(c *gin.Context) {
	choice := c.PostForm("submit")
	if choice == "login" {
		user_name := c.PostForm("username")
		user_password := c.PostForm("userpassword")
		if model.IsValid(model.UserInfo{UserName: user_name,UserPassword: user_password}) {
			c.SetCookie("username", user_name, ExpireTime, "/", "localhost", false, true)
			c.HTML(http.StatusOK, "auth-and-login.gohtml", gin.H{"message": 1})
		}
	}else {
		info, err := c.Request.Cookie("username")
		if err != nil {
			log.Println(err)
			return
		}
		scope := c.PostForm("scope")
		expire_time := c.PostForm("expire")
		redirect_uri := c.PostForm("redirect_uri")
		authCodeInfo, _ := model.GenerateAuthzCode(info.Value, scope, expire_time, redirect_uri)
		model.StoreAuthCodeInfo(authCodeInfo)
		c.Redirect(http.StatusFound, redirect_uri + "?code=" + authCodeInfo.AuthCode)
	}
}

func getAuthAndLoginPage(c *gin.Context) {
	info, err := c.Request.Cookie("username")
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "auth-and-login.gohtml", gin.H{"message": -1})
	}else {
		if !model.IsValidUserName(info.Value, c) {
			c.HTML(http.StatusOK, "auth-and-login.gohtml", gin.H{"message": -1})
		} else {
			c.HTML(http.StatusOK, "auth-and-login.gohtml", gin.H{"message": 1})
		}
	}
}
