package controller

import (
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	base_path = "~/Desktop/picture_store"
)

func PictureStore(server *gin.Engine){
	server.GET("/picture", BasicPicturePage)
	server.GET("/picture/manage", UploadPicture)
}

func BasicPicturePage(c *gin.Context){
	c.HTML(http.StatusOK, "picture-store.html", nil)
}

func UploadPicture(c *gin.Context){
	pictureInfo := model.Picture{}
	_ = c.BindJSON(&pictureInfo)
	model.Upload(base_path, pictureInfo)
	c.Redirect(http.StatusPermanentRedirect, "/picture")
}