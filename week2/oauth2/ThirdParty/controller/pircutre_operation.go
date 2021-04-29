package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oauth2/ThirdParty/model"
)

func PictureOperation(app *gin.Engine) {
	app.GET("/home/picture", PicturePage)
	app.POST("/home/picture", uploadPicture)
	app.DELETE("/home/picture", deletePicture)
	app.PUT("/home/picture", uploadPicture)
}

func PicturePage(c *gin.Context) {
	c.HTML(http.StatusOK, "picture.html", nil)
}

func uploadPicture(context *gin.Context) {
	accessTokenInfo, err := model.GetNewestAccessTokenInfo()
	if err != nil {
		log.Println(err)
		return
	}
	f, err := context.FormFile("picture")
	if err != nil {
		log.Println(err)
		return
	}
	name := f.Filename
	request, _ := http.NewRequest(http.MethodPut, "http://localhost:9090/picture", nil)
	request.FormFile()
	client := http.Client{}
	client.Do(sponse)
	context.HTML(http.StatusOK, "picture.html", nil)
}

func deletePicture(context *gin.Context) {
	http.NewRequest(http.MethodDelete, "http://localhost:9090/picture", nil)
	context.HTML(http.StatusOK, "picture.html", nil)
}

func getPicture(context *gin.Context) {
	http.NewRequest(http.MethodGet, "http://localhost:9090/picture", nil)
	context.HTML(http.StatusOK, "picture.html", nil)

}
