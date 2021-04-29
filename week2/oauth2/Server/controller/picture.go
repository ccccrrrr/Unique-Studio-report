package controller

import "github.com/gin-gonic/gin"

func PictureOperation(server *gin.Engine) {
	server.GET("/picture_store", checkPicture)
	server.DELETE("/picture_store", deletePicture)
	server.PUT("/picture_store", postPicture)
}

func postPicture(context *gin.Context) {

}

func deletePicture(context *gin.Context) {

}

func checkPicture(context *gin.Context) {

}