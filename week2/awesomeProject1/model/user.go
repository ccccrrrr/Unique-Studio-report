package model

type User struct {
	UserName string `json:"username" gorm:"username"`
	UserPassword string `json:"userpassword" gorm:"userpassword"`
}

func InsertUser(user User) error {
	var u User
	if err := db.Where("username = ?", user.UserName).First(&u).Error; err != nil {
		return err
	}
	db.Create(&user)
	return nil
}

func SearchUser(name string) User {
	var u User
	if err := db.Where("username = ?", name).First(&u).Error; err != nil {
		return User{}
	}else {
		return u
	}
}

func Login(u User) bool {
	var user User
	if err := db.Where("username = ?", u.UserName).First(&user); err != nil {
		return false
	}else {
		if user.UserPassword == u.UserPassword {
			return true
		}else {
			return false
		}
	}
}