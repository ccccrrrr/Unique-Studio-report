package controller

import (
	"awesomeproject1/service/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var Username string

type instruction struct {
	Direct string `json:"direct"`
}

func Login(server *gin.Engine) {
	server.GET("/login", Choose_)
	server.POST("/login", Choose)
	server.POST("/login/register", Register)
	server.GET("/login/register", Register_)
	server.GET("/login/login", _login)
	server.POST("/login/login", _login_)
}

func Register_(c *gin.Context) {
	c.HTML(http.StatusOK, "Login-register.html", nil)
}

func _login_(c *gin.Context) {
	info := model.User{}
	info.UserName = c.PostForm("username")
	info.UserPassword = c.PostForm("userpassword")
	info.LastLoginTime = time.Now()
	tmp, err := c.Request.Cookie("username")
	if err == nil && tmp.Value == info.UserName {
		Username = tmp.Value
		c.JSON(http.StatusOK, gin.H{"message": "already have user log in"})
		return
	}
	if model.Login(info) {
		Username = info.UserName
		c.SetCookie("username", info.UserName, 1000, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"massage":  "login successfully",
			"username": info.UserName,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"massage": "fail to login",
		})
	}
}

func Choose_(c *gin.Context) {
	info, err := c.Request.Cookie("username")
	if err == nil && info.Value != "" {
		Username = info.Value
		c.HTML(http.StatusOK, "have-login.gohtml", gin.H{
			"username": info,
		})
	} else {
		Username = ""
		log.Println(c.Request.Cookie("username"))
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func Choose(c *gin.Context) {
	json := instruction{}
	_ = c.BindJSON(&json)
	if json.Direct == "login" {
		c.Redirect(http.StatusPermanentRedirect, "/login/login")
	} else if json.Direct == "register" {
		c.Redirect(http.StatusMovedPermanently, "/login/register")
	} else if json.Direct == "error" {
		c.JSON(http.StatusOK, gin.H{
			"warning": "wrong selection!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"warning": "error",
		})
	}
}

func Register(c *gin.Context) {
	info := model.User{
		c.PostForm("username"),
		c.PostForm("userpassword"),
		time.Now(),
	}
	if info.UserName == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "user name is null",
		})
	}
	log.Printf("%+v", info)
	if err := model.InsertUser(info); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "fail to register",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "successful register",
		"user_name": info.UserName,
	})
}

func _login(c *gin.Context) {
	c.HTML(http.StatusOK, "login-login.html", nil)
}

func IsLogin(username string) bool {
	if username == Username {
		return true
	}else{
		return false
	}
}
