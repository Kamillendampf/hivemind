package DAO

import (
	"awesomeProject/structs"
	"errors"
	"gorm.io/gorm"
	"log"
)

func InsertDefaultTag(t structs.Tag) {
	db.Create(&t)
}

func InsertTag(t *structs.Tag) {
	err := db.Where("tag_name=?", t.TagName).First(t)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			db.Create(&t)
		} else {
			log.Println(err)
		}
	} else {
		log.Println("Tag is already existing, so ignore it.")
	}
}

func GetTags() []structs.Tag {
	var tags []structs.Tag
	if err := db.Find(&tags); err != nil {
		log.Println(err)
	}
	return tags
}
