package Controller

import (
	"awesomeProject/Service"
	"awesomeProject/structs"
	"net/http"
)

/*
SecuredMiddleware()
Controlling access to the secured area of the application.
This function returns a bool value and could deny the access.
*/
func SecuredMiddleware(r *http.Request) bool {
	var credentials = structs.Credentials{UserIdentificationKey: r.Header.Get("Authorization")}
	if Service.CredentialCheck(&credentials) {
		return true
	} else {
		return false
	}

}
