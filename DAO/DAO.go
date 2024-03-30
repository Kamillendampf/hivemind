package DAO

import (
	"awesomeProject/structs"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	database()
	_connectGormDataBase()

	err := db.AutoMigrate(structs.User{}, structs.Questions{}, structs.Answers{})
	if err != nil {
		log.Fatal("Until auto migrate form database something went wrong Error: " + err.Error())
		return
	}
	log.Print(&db)
}

func _connectGormDataBase() {
	dbUser := "postgres"
	dbPass := "admin"
	dbName := "hivemind"
	dbHost := "localhost"
	dbPort := "5432"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s  port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Exec("SET client_encoding = 'UTF8'")
	if err != nil {
		panic("Something went wrong until connecting with gorm to Database. Error: " + err.Error())
	}
}
