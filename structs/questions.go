package structs

import "gorm.io/gorm"

type Questions struct {
	gorm.Model
	ID             uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	QuestionID     string `gorm:"column:question_id;" json:"question_id"`
	UserName       string `gorm:"column:user_name" json:"username"`
	PublicUserKey  string `gorm:"column:public_user_key" json:"public_user_key"`
	QuestionHeader string `gorm:"column:questions_header" json:"question_header"`
	QuestionText   string `gorm:"column:question_text" json:"question_text"`
	Tags           string `gorm:"column: tags"`
}

type Answers struct {
	gorm.Model
	ID            uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Root          string `gorm:"column:question_id;" json:"question_id"` //Falls du dich entscheidest dass es moeglich sein wird auf Anworten zu Antowrten
	UserName      string `gorm:"column:user_name" json:"username"`
	PublicUserKey string `gorm:"column:public_user_key" json:"public_user_key"`
	Answer        string `gorm:"column:question_text" json:"question_text"`
	Score         string `gorm:"column:score" json:"score"`
}
