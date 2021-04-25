package model

import (
	"time"
)


type User struct {
	UserName     string `json:"username" gorm:"user_name"`
	UserPassword string `json:"userpassword" gorm:"user_password"`
	LastLoginTime time.Time
}

var (
	duration = time.Minute * 20
)

func InsertUser(user User) error {
	var u User
	if err := Db.Table("users").Where("user_name = ?", user.UserName).First(&u).Error; err != nil {
		//log.Println(err)
		Db.Table("users").Table("users").Create(&user)
		return err
	}
	return nil
}

func Login(u User) bool {
	var user User
	if err := Db.Table("users").Where("user_name = ?", u.UserName).First(&user); err == nil {
		return false
	}else {
		if user.UserPassword == u.UserPassword {
			// need to update state
			Db.Table("users").Model(&user).Where("user_name = ?", u.UserName).Update("last_login_time", time.Now())
			return true
		}else {
			return false
		}
	}
}

func CheckState(name string) bool {
	var u User
	if err := Db.Table("users").Where("user_name = ?", name).First(&u).Error; err == nil {
		if time.Now().Sub(u.LastLoginTime) > duration {
			return false
		}else {
			return true
		}
	}
	return false
}

func IsValid(name string) bool {
	return CheckState(name)
}

//func LogOut(u User) bool {
//	// update state
//}