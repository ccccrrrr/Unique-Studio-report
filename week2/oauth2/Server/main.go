package main

import (
	"github.com/gin-gonic/gin"
	"oauth2/Server/controller"
	"oauth2/Server/model"
)

func init() {
	model.InitServerDatabase()
}

func main() {
	server := gin.Default()

	server.LoadHTMLGlob("./Server/static/html/*")

	controller.Login(server)
	controller.PictureOperation(server)
	controller.Register(server)
	controller.AccessTokenEndpoint(server)
	controller.IntrospectiveEndpoint(server)
	controller.AuthAndLogin(server)
//	controller.AuthzEndpoint(server)

	_ = server.Run(":9090")
}
