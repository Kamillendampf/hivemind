package Controller

import (
	"awesomeProject/Service"
	"awesomeProject/structs"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
)

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	if SecuredMiddleware(r) && r.Method == "POST" {
		var _question structs.Questions
		err := json.NewDecoder(r.Body).Decode(&_question)
		if err != nil && err != io.EOF {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				var _errorMessage = "read body from " + ip + "failed\n"
				http.Error(w, _errorMessage, http.StatusBadRequest)
			} else {
				http.Error(w, "unknown host", http.StatusInternalServerError)
			}
			Service.InsertQuestion(&_question)
			return
		}
	}
}

func GetQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		jsonized, err := json.Marshal(Service.GetQuestion())
		if err != nil {
			log.Println(err)
		}
		_, _ = w.Write(jsonized)
	}
}

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	if SecuredMiddleware(r) && r.Method == "POST" {
		var _answer structs.Answers
		err := json.NewDecoder(r.Body).Decode(&_answer)
		if err != nil && err != io.EOF {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				var _errorMessage = "read body from " + ip + "failed\n"
				http.Error(w, _errorMessage, http.StatusBadRequest)
			} else {
				http.Error(w, "unknown host", http.StatusInternalServerError)
			}
			Service.InsertAnswer(&_answer)
			return
		}
	}
}

func GetAnswers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req structs.Answers
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			ErrorMessage := err
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				var _errorMessage = "read body from " + ip + "failed\n"
				http.Error(w, _errorMessage, http.StatusInternalServerError)
			} else {

				_errorMessage := "read body from " + ip + " failed with error message: " + ErrorMessage.Error() + "\n"
				http.Error(w, _errorMessage, http.StatusBadRequest)
			}
			return
		}
		response, err := json.Marshal(Service.GetAnswers(&req))
		if err != nil {
			log.Println(err)
		}
		_, _ = w.Write(response)
	}
}

type rating struct {
	Vote   bool            `json:"vote"`
	Answer structs.Answers `json:"answer"`
}

func RateAnswers(w http.ResponseWriter, r *http.Request) {
	if SecuredMiddleware(r) && r.Method == "POST" {
		var rat rating
		err := json.NewDecoder(r.Body).Decode(&rat)
		if err != nil {
			ErrorMessage := err
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				var _errorMessage = "read body from " + ip + "failed\n"
				http.Error(w, _errorMessage, http.StatusInternalServerError)
			} else {

				_errorMessage := "read body from " + ip + " failed with error message: " + ErrorMessage.Error() + "\n"
				http.Error(w, _errorMessage, http.StatusBadRequest)
			}
			Service.RatingAnswer(rat.Vote, &rat.Answer)
			return
		}
	}
}
