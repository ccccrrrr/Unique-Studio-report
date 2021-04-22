package model

import "log"

type Picture struct {
	Name string `json:"picture_name"`
	Path string `json:"picture_path"`
//	Size string `json:"picture_size"`
}

func Upload(base_path string, picture Picture){
	// different directory path need to add
	var p1 Picture
	if err := dbPicture.Where("name = ?", picture.Name).First(&p1).Error; err == nil {
		log.Println("not found!")
		return
	}else {
		dbPicture.Create(&picture)
	}
	return
}