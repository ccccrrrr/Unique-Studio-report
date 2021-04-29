package model

import (
	"github.com/gin-gonic/gin"
	"log"
)

func CheckCookie(c *gin.Context, name string) (string, error) {
	body, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	return body.Value, nil
}

func IsValidUserName(username string, c *gin.Context) bool {
	_username, err := CheckCookie(c, "username")
	if err != nil {
		log.Println(err)
		return false
	}
	if _username == username {
		return true
	}
	return false
}
