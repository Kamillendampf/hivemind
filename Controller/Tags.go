package Controller

import (
	"awesomeProject/Service"
	"awesomeProject/structs"
	"encoding/json"
	"log"
	"net"
	"net/http"
)

func GetTags(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		response, err := json.Marshal(Service.GetTags())
		if err != nil {
			log.Println(err)
		}
		_, _ = w.Write(response)
	}
}

func CreateTag(w http.ResponseWriter, r *http.Request) {
	if SecuredMiddleware(r) && r.Method == "POST" {
		var t structs.Tag
		err := json.NewDecoder(r.Body).Decode(&t)
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
		Service.CreateTags(&t)
	}
}

func CreateDefaultSystemTags() {
	Service.CreateDefaultTags()
}
