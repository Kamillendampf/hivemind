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

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var _credentials structs.Credentials
		err := json.NewDecoder(r.Body).Decode(&_credentials)
		if err != nil {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				var _errorMessage = "read body from " + ip + "failed\n"
				http.Error(w, _errorMessage, http.StatusBadRequest)
			} else {
				http.Error(w, " unknown host", http.StatusInternalServerError)
			}
			return
		}

		if Service.CredentialCheck(&_credentials) {
			w.WriteHeader(http.StatusAccepted)
			mapingOutput := make(map[string]string)
			mapingOutput["company"] = Service.FindUserPublicKeyFromCredentials(&_credentials)
			jsonizedOutput, err := json.Marshal(mapingOutput)
			if err != nil {
				log.Print("JSON faild to read with: ", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(jsonizedOutput)
			return
		} else {
			log.Printf("login %s was unsuccessful\n", _credentials.UserIdentificationKey)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var _user structs.User
		err := json.NewDecoder(r.Body).Decode(&_user)
		if err != nil && err != io.EOF {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				var _errorMessage = "read body from " + ip + "failed\n"
				http.Error(w, _errorMessage, http.StatusBadRequest)
			} else {
				http.Error(w, "unknown host", http.StatusInternalServerError)
			}
			return
		}

		Service.InsertUser(&_user)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

}
