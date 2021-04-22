package controller

import (
	"awesomeproject1/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func ThirdParty(app *gin.Engine) {
	app.GET("/HomePage", getAppHomePage)
	app.DELETE("/HomePage", deletePicture)
	app.POST("/HomePage", uploadPicture)
}

func deletePicture(c *gin.Context){
	//log.Println(c.Query("user_name"))
	info := model.DeleteInfo{}
	_ = c.BindJSON(&info)
	if model.IsValid(info.DeleteUser) {
		res := model.DeletePicture(info)
		log.Println(info)
		if res == true {
			log.Println("delete successfully")
		}else {
			log.Println("unsuccessful delete")
		}
	}else {
		log.Println("invalid user name")
	}
}

func uploadPicture(c *gin.Context){
	info := model.Picture{}
	_ = c.BindJSON(&info)
	log.Println(info)
	if model.IsValid(info.CreateUserName) {
		res := model.Upload("", info)
		if res == true {
			log.Println("upload successfully")
		}else {
			log.Println("unsuccessful upload operation")
		}
	}else {
		c.JSON(http.StatusOK, gin.H{
			"message": "fail...",
		})
	}
}


func getAppHomePage(c *gin.Context) {
	username := c.Query("user_name")
	if model.IsValid(username) {
		c.HTML(http.StatusOK, "app-home-login.html", nil)
	}else {
		c.HTML(http.StatusOK, "app-home-notlogin.html", nil)
	}
}

