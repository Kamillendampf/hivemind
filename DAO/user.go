package DAO

import (
	"awesomeProject/structs"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func InsertUser(user *structs.User) {
	db.Create(user)
}

func FindUserCredentials(credentials *structs.Credentials) bool {
	user := new(structs.User)

	fmt.Println("Adresse der user Variable:", &db)

	err := db.First(&user, "useridentificationkey=?", credentials.UserIdentificationKey)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return false
	} else if err.Error != nil {
		log.Println(err.Error.Error())
		return false
	}
	return true
}

func DeleteUser(credentials *structs.Credentials) {
	user := new(structs.User)
	db.Where("user_identification_key=?", credentials.UserIdentificationKey).Delete(&user)
}

func FindUserPublicKeyFromCredentials(credentials *structs.Credentials) string {
	var userCompany structs.User
	err := db.Where("useridentificationkey = ?", credentials.UserIdentificationKey).First(&userCompany)
	if err.Error != nil {
		log.Fatal(err)
		return ""
	}
	return userCompany.PublicUserKey
}
