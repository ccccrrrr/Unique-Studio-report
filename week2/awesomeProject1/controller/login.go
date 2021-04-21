package controller

import (
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(server *gin.Engine){
	server.POST("/login/register", Register)
	server.POST("/login/login", _login)
}

func Register(c *gin.Context){
	info := model.User{}
	err := c.BindJSON(&info)
	if err != nil {
		log.Println(err)
		return
	}
	if err = model.InsertUser(info); err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successful register",
	})
}

func _login(c *gin.Context){
	info := model.User{}
	err := c.BindJSON(&info)
	if err != nil {
		panic(err)
	}
	if model.Login(info) {
		c.JSON(http.StatusOK, gin.H{
			"massage": "login successfully",
			"username": info.UserName,
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"massage": "fail to login",
		})
	}

}
