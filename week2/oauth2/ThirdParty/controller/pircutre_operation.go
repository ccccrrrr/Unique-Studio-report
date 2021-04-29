package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"oauth2/ThirdParty/model"
)

func PictureOperation(app *gin.Engine) {
	app.GET("/home/picture", PicturePage)
	app.POST("/home/picture", PictureCRUD)
}

func PicturePage(c *gin.Context) {
	c.HTML(http.StatusOK, "picture.html", nil)
}

func PictureCRUD(context *gin.Context) {
	if context.PostForm("submit") == "show" {
		accessTokenInfo, err := model.GetNewestAccessTokenInfo()
		if err != nil {
			log.Println(err)
			return
		}

		req, _ := http.NewRequest(http.MethodGet, "http://localhost:9090/picture", nil)

		req.Header.Add("access_token", accessTokenInfo.AccessToken)

		http.DefaultClient.Do(req)
		context.HTML(http.StatusOK, "picture.html", nil)
		return
	}

	if context.PostForm("submit") == "delete" {
		picture_name := context.PostForm("picture_name")
		accessTokenInfo, err := model.GetNewestAccessTokenInfo()
		if err != nil {
			log.Println(err)
			return
		}

		req, _ := http.NewRequest(http.MethodDelete, "http://localhost:9090/picture", nil)

		req.Header.Add("access_token", accessTokenInfo.AccessToken)
		req.Header.Add("picture_name", picture_name)
		http.DefaultClient.Do(req)
		context.HTML(http.StatusOK, "picture.html", nil)
		return
	}

	if context.PostForm("submit") == "add" {
		accessTokenInfo, err := model.GetNewestAccessTokenInfo()
		if err != nil {
			log.Println(err)
			return
		}

		f, header, err := context.Request.FormFile("picture")
		if err != nil {
			log.Println(err)
			return
		}

		bodyBuffer := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuffer)
		fileWriter, _ := bodyWriter.CreateFormFile("file", header.Filename)
		_, _ = io.Copy(fileWriter, f)

		content_type := bodyWriter.FormDataContentType()

		req, _ := http.NewRequest(http.MethodPut, "http://localhost:9090/picture", nil)
		req.Header.Add("access_token", accessTokenInfo.AccessToken)
		req.Header.Add("file_name", header.Filename)
		req.Header.Set("Content-Type", content_type)
		client := http.Client{}
		_, _ = client.Do(req)

		context.HTML(http.StatusOK, "picture.html", nil)
	}
}


