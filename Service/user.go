package Service

import (
	"awesomeProject/DAO"
	"awesomeProject/structs"
)

// insert, update, delete

func InsertUser(user *structs.User) {
	DAO.InsertUser(user)
}

func DeleteUser(credentials *structs.Credentials) {
	DAO.DeleteUser(credentials)
}

func CredentialCheck(credentials *structs.Credentials) bool {
	return DAO.FindUserCredentials(credentials)
}

func FindUserPublicKeyFromCredentials(credentials *structs.Credentials) string {
	return DAO.FindUserPublicKeyFromCredentials(credentials)
}
