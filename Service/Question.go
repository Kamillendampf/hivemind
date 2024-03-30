package Service

import (
	"awesomeProject/DAO"
	"awesomeProject/structs"
	"log"
	"strconv"
)

func InsertQuestion(questions *structs.Questions) {
	DAO.InsertQuestion(questions)
}

func GetQuestion() *structs.Questions {
	return DAO.GetQuestion()
}

func InsertAnswer(answers *structs.Answers) {
	DAO.InsertAnswer(answers)
}

func GetAnswers(answers *structs.Answers) []structs.Answers {
	return DAO.GetAnswers(answers)
}

func RatingAnswer(vote bool, answer *structs.Answers) {
	if vote {
		rating, err := strconv.Atoi(DAO.GetRating(answer).Score)
		if err != nil {
			log.Println(err)
		}
		DAO.UpdateRating(
			answer,
			rating+1,
		)

	}
}
