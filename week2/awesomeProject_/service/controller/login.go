package controller

import (
	model2 "awesomeproject1/service/model"
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
	server.GET("/login/checkcookie", checkCookie)
}

func checkCookie(c *gin.Context) {
	cookie, err := c.Cookie("username")
	if err == nil {
		c.JSON(http.StatusOK, cookie)
	}else {
		c.JSON(http.StatusOK, "")
	}
}

func Register_(c *gin.Context) {
	//type User struct {
	//	UserName     string `json:"username" gorm:"user_name"`
	//	UserPassword string `json:"userpassword" gorm:"user_password"`
	//	LastLoginTime time.Time
	//}
	c.HTML(http.StatusOK, "Login-register.html", nil)
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "register information",
	//})
}

func _login_(c *gin.Context) {
	info := model2.User{}
	info.UserName = c.PostForm("username")
	info.UserPassword = c.PostForm("userpassword")
	info.LastLoginTime = time.Now()
	//	log.Println(c.GetHeader("Cookie"))
	tmp, err := c.Request.Cookie("username")
	if err == nil && tmp.Value == info.UserName {
		Username = tmp.Value
		c.JSON(http.StatusOK, gin.H{"message": "already have user log in"})
		return
	}
	if model2.Login(info) {
		//model.Db.Where("user_name = ?", info.UserName).First(&tmp)
		//model.Db.Model(&tmp).Update("last_login_time", time.Now())
		Username = info.UserName
		c.SetCookie("username", info.UserName, 1000, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"massage":  "login successfully",
			"username": info.UserName,
		})
		//c.Redirect(http.StatusPermanentRedirect, "/picture")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"massage": "fail to login",
		})
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "login information",
	//})
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
	info := model2.User{
		c.PostForm("username"),
		c.PostForm("userpassword"),
		time.Now(),
	}
	//err := c.BindJSON(&info)
	//info.LastLoginTime = time.Now()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	if info.UserName == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "user name is null",
		})
	}
	log.Printf("%+v", info)
	if err := model2.InsertUser(info); err == nil {
		//log.Println(err)
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

func isLogin(username string, c *gin.Context) bool {
	cookieInfo, err := c.Cookie("username")
	log.Println(err)
	//req, _ := http.Post("http://localhost:9090/login", "text", nil)
	log.Println(err)
	log.Println("login.go")
	if err == nil && cookieInfo == username {
		return true
	}
	if err == nil && cookieInfo != username {
		log.Println("not the same user...")
	}
	return false
}
