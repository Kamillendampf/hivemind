package main

import (
	"awesomeProject/Controller"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

//const file = "../Assents/server.config"
const serverError string = "Server failed with failure code: "

type Config struct {
	Port string
	Cert string
	Key  string
}

var config Config

func init() {
	//loadConfig(file)
}

func main() {
	log.Printf("Server is listening on Port %s \n", config.Port)

	//Core-API
	http.HandleFunc("/", Controller.LandingPageHandler)

	http.HandleFunc("/login", Controller.Login)
	http.HandleFunc("/registerUser", Controller.RegisterUser)
	http.HandleFunc("/createQuestion", Controller.CreateQuestion)
	http.HandleFunc("/getQuestions", Controller.GetQuestion)
	http.HandleFunc("/createAnswer", Controller.CreateAnswer)
	http.HandleFunc("/getAnswers", Controller.GetAnswers)
	http.HandleFunc("/rateAnswers", Controller.RateAnswers)
	http.HandleFunc("/getTags", Controller.GetTags)
	http.HandleFunc("/createTag", Controller.CreateTag)

	//creating server
	//err := http.ListenAndServeTLS(config.Port, config.Cert, config.Key, nil)
	err := http.ListenAndServe("localhost:8080", nil) //This should be only used in development mode!! Dont use it in for production
	if err != nil {
		log.Fatal(serverError, err)
		os.Exit(1)
	}
}

/*
loadingConfig()

Type: private

This function loads the configuration file for tls connections.
The path of the configuration file is default set on "/etc/hivemind/server.config". Change that if it is necessary.
*/
/*func loadConfig(file string) {
	_file, err := os.Open(file)
	if err != nil {
		log.Fatal("Error until reading Configuration file: error code: " + err.Error())
	}
	defer _file.Close()

	decoder := json.NewDecoder(_file)
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatal("error ocured until decodig file to struct Config with error code: " + err.Error())
	}
}*/
