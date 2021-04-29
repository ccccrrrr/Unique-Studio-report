package model

import (
	"github.com/gin-gonic/gin"
	"log"
)

func CheckCookie(c *gin.Context, name string) (string, bool) {
	body, err := c.Request.Cookie(name)
	if err != nil {
		log.Println(err)
		return "", false
	}
	return body.Value, true
}

func IsValidUserName(username string, c *gin.Context) bool {
	_username, returnType := CheckCookie(c, "username")
	if returnType == false {
		return false
	}
	if _username == username {
		return true
	}
	return false
}