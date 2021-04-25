package model

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"gorm.io/gorm"

	//	"strconv"
	"time"
)

type AccessTokenInfo struct {
	UserName    string    `json:"user_name"`
	AuthCode    string    `json:"code"`
	Scope       string    `json:"scope"`
	AccessToken string    `json:"access_token"`
	ExpireTime  time.Time `json:"expire_time"`
	CreateTime  time.Time `json:"create_time"`
}

type AuthInfo struct {
	gorm.Model
	UserName   string    `json:"user_name"`
	Code       string    `json:"code"`
	Scope      string    `json:"scope"`
	ExpireTime string    `json:"expire_time"`
	TimeNow    time.Time `json:"time_now"`
}

type GetInfo1 struct {
	Scope     string `json:"scope"`
	UseName   string `json:"user_name"`
	TimeLimit string `json:"time_limit"`
}

type SendInfo1 struct {
	Scope        string
	TimeLimit    string
	ClientId     string
	ResponseType string
	RedirectUri  string
}

func GenerateInfo(temp GetInfo1, c *gin.Context) SendInfo1 {
	var res SendInfo1
	res.TimeLimit = temp.TimeLimit
	res.Scope = temp.Scope

	// let's temporary assume that clientid=username
	res.ClientId = generateClientId(temp.UseName)
	res.RedirectUri = "https://localhost:9090/login"
	//res.ResponseType = c.con
	return res
}

func generateClientId(s string) string {
	return s
}

func GenerateTempAuthCode(clientId string) string {
	w := randstr.Hex(10)
	return w
}

func AuthCodeToUser(client_id string, scope string, expire string, code string) {
	var authInfo AuthInfo
	authInfo.UserName = client_id
	authInfo.Scope = scope
	authInfo.Code = code
	authInfo.ExpireTime = expire
	authInfo.TimeNow = time.Now()
	InsertAuthCodeInfo(authInfo)
}

func FindInfoThroughAuthCode(authCode string) (AuthInfo, bool) {
	var authInfo AuthInfo
	if err := Db.Where("code = ?", authCode).First(&authInfo).Error; err == nil {
		return authInfo, true
	} else {
		return authInfo, false
	}
}

func GenerateAccessToken(authCode string, expireTime int) AccessTokenInfo {
	var accessTokenInfo AccessTokenInfo
	authCodeInfo, _ := FindInfoThroughAuthCode(authCode)
	accessTokenInfo.Scope = authCodeInfo.Scope
	accessTokenInfo.CreateTime = time.Now()
	accessTokenInfo.UserName = authCodeInfo.UserName
	accessTokenInfo.ExpireTime = authCodeInfo.TimeNow.Add(time.Minute * time.Duration(expireTime))
	accessTokenInfo.AuthCode = authCodeInfo.Code
	accessTokenInfo.AccessToken = randstr.Hex(20)
	Db.Create(&accessTokenInfo)
	return accessTokenInfo
}
