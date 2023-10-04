package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ListAllItems() {
	db, err := sql.Open("sqlite3", "db/invDB.db")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db) // Make sure to close the database when you're done with it.

	querySQL := "SELECT * FROM item"
	rows, err := db.Query(querySQL)
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var category string
		var name string
		var desc string
		var price float64
		err := rows.Scan(&category, &name, &desc, &price)
		if err != nil {
			panic(err)
		}

		println("ITEM:")
		fmt.Println("Category: ", category)
		fmt.Println("Name: ", name)
		fmt.Println("Description: ", desc)
		fmt.Println("Price: ", price)

	}
}

func AddItem(item Food) {
	db, err := sql.Open("sqlite3", "db/invDB.db")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db) // Make sure to close the database when you're done with it.

	querySQL := fmt.Sprintf(
		"INSERT INTO Item VALUES  ('%s', '%s', '%s', %d)",
		item.category,
		item.name,
		item.desc,
		item.price,
	)
	_, err = db.Exec(querySQL)
	if err != nil {
		return
	}
	if err != nil {
		panic(err)
	}
}

func ExecuteSqlFile(file string) {
	db, err := sql.Open("sqlite3", "db/invDB.db")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	// Read the SQL file
	sqlFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlFile.Close()

	sqlBytes, err := ioutil.ReadAll(sqlFile)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the SQL statements
	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal(err)
	}
}
