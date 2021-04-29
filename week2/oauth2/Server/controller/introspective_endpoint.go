package controller

import "github.com/gin-gonic/gin"

func IntrospectiveEndpoint(server *gin.Engine) {
	server.GET("/introspective", _checkPicture)
	server.DELETE("/introspective", _deletePicture)
	server.PUT("/introspective", _postPicture)
}

func _postPicture(context *gin.Context) {

}

func _deletePicture(context *gin.Context) {

}

func _checkPicture(context *gin.Context) {

}
