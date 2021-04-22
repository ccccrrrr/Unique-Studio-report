package controller

import (
	"awesomeproject1/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type instruction struct {
	Direct string `json:"direct"`
}

func Login(server *gin.Engine){
	server.GET("/login", Choose_)
	server.POST("/login", Choose)
	server.POST("/login/register", Register)
	server.GET("/login/register", Register_)
	server.GET("/login/login", _login)
	server.POST("/login/login", _login_)
}

func Register_(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "register information",
	})
}

func _login_(c *gin.Context){

	info := model.User{}
	err := c.BindJSON(&info)
	//tmp := model.User{}
	log.Println(info)
	if err != nil {
		log.Println(err)
		return
	}
	if model.Login(info) {
		//model.Db.Where("user_name = ?", info.UserName).First(&tmp)
		//model.Db.Model(&tmp).Update("last_login_time", time.Now())
		c.JSON(http.StatusOK, gin.H{
			"massage": "login successfully",
			"username": info.UserName,
		})
		//c.Redirect(http.StatusPermanentRedirect, "/picture")
	}else {
		c.JSON(http.StatusOK, gin.H{
			"massage": "fail to login",
		})
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "login information",
	//})
}

func Choose_(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", nil)
}

func Choose(c *gin.Context){
	json := instruction{}
	_ = c.BindJSON(&json)
	if json.Direct == "login" {
		c.Redirect(http.StatusPermanentRedirect, "/login/login")
	}else if json.Direct == "register" {
		c.Redirect(http.StatusMovedPermanently, "/login/register")
	}else if json.Direct == "error" {
		c.JSON(http.StatusOK, gin.H{
			"warning": "wrong selection!",
		})
	}else {
		c.JSON(http.StatusBadRequest, gin.H{
			"warning": "error",
		})
	}
}

func Register(c *gin.Context){
	info := model.User{}
	err := c.BindJSON(&info)
	info.LastLoginTime = time.Now()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v", info)
	if err = model.InsertUser(info); err == nil {
		//log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "fail to register",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successful register",
	})
}

func _login(c *gin.Context){

	c.HTML(http.StatusOK, "login-login.html", nil)

	//info := model.User{}
	//err := c.BindJSON(&info)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//if model.Login(info) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"massage": "login successfully",
	//		"username": info.UserName,
	//	})
	//}else {
	//	c.JSON(http.StatusOK, gin.H{
	//		"massage": "fail to login",
	//	})
	//}

}
