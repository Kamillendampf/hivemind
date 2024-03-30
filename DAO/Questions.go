package DAO

import (
	"awesomeProject/structs"
	"log"
)

func InsertQuestion(questions *structs.Questions) {
	db.Create(questions)
}

func GetQuestion() *structs.Questions {
	question := new(structs.Questions)
	db.Find(&question)
	return question
}

func InsertAnswer(answers *structs.Answers) {
	db.Create(answers)
}

func GetAnswers(answer *structs.Answers) []structs.Answers {
	var answers []structs.Answers
	if err := db.Where("root=?", answer.Root).Find(&answers); err.Error != nil {
		log.Println(err)
	}
	return answers
}

func UpdateRating(answer *structs.Answers, rating int) {
	if err := db.Model(answer).Where("UserIdentificationKey=?", answer.ID).Update("score", rating); err.Error != nil {
		log.Println(err)
	}
}

func GetRating(answers *structs.Answers) *structs.Answers {
	answer := new(structs.Answers)
	db.First(&answer, answers.ID)
	return answer
}
