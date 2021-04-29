package main

import (
	"github.com/gin-gonic/gin"
	"oauth2/ThirdParty/controller"
	"oauth2/ThirdParty/model"
)

func init() {
	model.InitDataBase()
}

func main() {
	app := gin.Default()

	app.LoadHTMLGlob("./ThirdParty/static/html/*")

	controller.PictureOperation(app)
	controller.HomePage(app)
	controller.AuthzCode(app)
	controller.AccessToken(app)

	_ = app.Run(":9001")
}
