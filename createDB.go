package main

import (
	"database/sql"
	"os"
)
import "log"

func CreateDB() {
	db, _ := sql.Open("sqlite3", "db/invDB.db")

	if db != nil {
		log.Print("Existing database discovered, deleting....")
		err := os.Remove("db/invDB.db")
		if err != nil {
			log.Print(err.Error())
			return
		}
	}

	executeSqlFile("sql/item.sql")

	log.Print("New Database created")

}
