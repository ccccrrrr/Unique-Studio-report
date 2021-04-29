package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oauth2/Server/model"
)

func PictureOperation(server *gin.Engine) {
	server.GET("/picture", checkPicture)
	server.DELETE("/picture", deletePicture)
	server.PUT("/picture", postPicture)
}

func postPicture(context *gin.Context) {
	//_, header, err := context.Request.FormFile("file")
	//if err != nil {
	//	log.Printf("form file error %v", err)
	//	return
	//}
	access_token := context.GetHeader("access_token")
	file_name := context.GetHeader("file_name")
	res, err := model.IsValidToken(access_token)
	if err != nil {
		log.Println(err)
		return
	}
//	dst := "/home/ccrr/Desktop/mission/week2/oauth/picture_store/" + header.Filename
//	err = context.SaveUploadedFile(header, dst)
//	if err != nil {
//		log.Println(err)
//		return
//	}
	returnType := model.IsValidUploadOperation(access_token)
	if returnType == false {
		log.Println("no access to upload operation!")
		return
	}
	model.Upload(file_name, res.UserName)
	context.JSON(http.StatusOK, nil)
}

func deletePicture(context *gin.Context) {
	access_token := context.GetHeader("access_token")
	picture_name := context.GetHeader("picture_name")
	res, err := model.IsValidToken(access_token)
	if err != nil {
		log.Println(err)
		return
	}
	returnType := model.IsValidDeleteOperation(access_token)
	if returnType == false {
		log.Println("no access to delete operation!")
		return
	}
	err = model.DeletePicture(picture_name, res.UserName)
	if err != nil {
		log.Println(err)
		return
	}
}

func checkPicture(context *gin.Context) {
	access_token := context.GetHeader("access_token")
	res, err := model.IsValidToken(access_token)
	if err != nil {
		log.Println(err)
		return
	}
	returnType := model.IsValidGetOperation(access_token)
	if returnType == false {
		log.Println("no access to get operation!")
		return
	}
	_ = model.SearchPicture(res.UserName)
}