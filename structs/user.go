package structs

import "gorm.io/gorm"

type Credentials struct {
	UserIdentificationKey string `json:"UserIdentificationKey"`
}

type User struct {
	gorm.Model
	ID                    uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserIdentificationKey string `gorm:"column:user_identification_key" json:"UserIdentificationKey"`
	PublicUserKey         string `gorm:"column:public_user_key" json:"public_user_key"`
	Score                 int    `gorm:"column:score" json:"score"`
}
