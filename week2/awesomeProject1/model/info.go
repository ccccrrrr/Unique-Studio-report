package model

type RequestPermission struct {
	LoginId  string `json:"login_id"`
	Password string `json:"password"`
	// tbc
}
