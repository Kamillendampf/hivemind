package structs

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	ID      uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TagName string `gorm:"column:tag_name;primaryKey;autoIncrement" json:"tag_name"`
	TagUser string `gorm:"column:tag_user;primaryKey;autoIncrement" json:"tag_user"`
}
