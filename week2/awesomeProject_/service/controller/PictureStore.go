package controller

import (
	model2 "awesomeproject1/service/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var (
	base_path = "~/Desktop/picture_store"
	UserName string
	UserPassword string
)

//type TotalInformation struct {
//	PictureInfo model.Picture
//	UserInfo 	model.Picture
//}


func PictureStore(server *gin.Engine){
	server.GET("/picture", BasicPicturePage)
	server.PUT("/picture", getPicture)
//	server.POST("/picture", getUserInfo)
	server.GET("/picture/manage", UploadPicture)
}


func getPicture(c *gin.Context) {

	json := model2.Picture{}
	_ = c.BindJSON(&json)

	// I don't know why the time is not stored in mysql
	//json.CreateTime = time.Now()
	log.Println(json)
	u := model2.User{
		UserName:      json.CreateUserName,
		UserPassword:  json.CreateUserPassword,
		LastLoginTime: time.Now(),
	}
	// need to judge if the same picture has been stored
	if model2.Login(u) {
		// need to reorganize the picture
		model2.Db.Table("pictures").Create(json)
		c.JSON(http.StatusOK, gin.H{
			"message": "upload successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "wrong password or the login is out of date. you have to log in first",
		})
	}
}

func BasicPicturePage(c *gin.Context){
	c.HTML(http.StatusOK, "picture-store.html", nil)
}

func UploadPicture(c *gin.Context){
	pictureInfo := model2.Picture{}
	_ = c.BindJSON(&pictureInfo)
	model2.Upload(base_path, pictureInfo)
	c.Redirect(http.StatusPermanentRedirect, "/picture")
}