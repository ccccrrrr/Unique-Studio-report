package main

import (
	"math/rand"
	"time"
)
var letters = []byte("abcdefghjkmnpqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~")

func init() {
	rand.Seed(time.Now().Unix())
}

func generateCodeVerifier() string{
	length := rand.Intn(86) + 43
	var codeVerifier string
	for i := 0; i < length ; i++ {
		codeVerifier += string(letters[rand.Intn(43)])
	}
	return codeVerifier
}
