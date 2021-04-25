package controller

import (
	"awesomeproject1/service/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func PictureStore(server *gin.Engine) {
	server.GET("/picture", BasicPicturePage)
	server.POST("/picture", managePicture)
}

func managePicture(c *gin.Context) {
	if Username == ""{
		c.HTML(http.StatusOK, "not-login.html", nil)
		return
	}
	choice := c.PostForm("submit")
	log.Printf("choice: " + choice)
	switch choice {
	case "delete":
		picture_name := c.PostForm("picture_name")
		// I dont know this operation does not work
		err := os.RemoveAll("/home/ccrr/go/src/awesomeProject/picture_store/" + picture_name)
		if err != nil {
			log.Println(err)
			return
		}
		model.DeletePicture(picture_name, Username)
	case "add":
		f, err := c.FormFile("picture")
		if err != nil {
			log.Println(err)
			return
		}
		name := f.Filename
		dst := "/home/ccrr/go/src/awesomeProject1/picture_store/" + f.Filename
		err = c.SaveUploadedFile(f, dst)
		if err != nil {
			log.Printf("file store error: %v", err)
			return
		}
		model.Upload(name, Username)
	case "show":
		model.Show(Username)
	}
	c.HTML(http.StatusOK, "picture-store.html", nil)
}

func BasicPicturePage(c *gin.Context) {
	c.HTML(http.StatusOK, "picture-store.html", nil)
}
