package main

import (
	"awesomeproject1/thirdparty/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("./thirdparty/static/html/*")

	controller.ThirdParty(app)

	_ = app.Run(":9001")
}
