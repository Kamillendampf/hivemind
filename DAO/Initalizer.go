package DAO

import (
	"github.com/go-pg/pg"
	"log"
)

func database() {

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "admin",
		Database: "hivemind",
		Addr:     "localhost:5432",
	})
	var exists bool
	_, err := db.QueryOne(pg.Scan(&exists), "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = ?)", "hivemind")
	if err != nil {
		if !exists {
			log.Println("Fehler beim Abfragen der Datenbankexistenz: " + err.Error())
			// Die Datenbank existiert nicht, also erstellen wir sie
			_, errc := db.Exec("CREATE DATABASE hivemind")
			if errc != nil {
				panic("Fehler beim Erstellen der Datenbank: " + err.Error())
			}
			log.Println("Die Datenbank wurde erfolgreich erstellt.")
		} else {
			log.Println("Die Datenbank existiert bereits.")
		}

	}

}
