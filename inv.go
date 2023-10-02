package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// Item /**
// Using an interface for now but wondering if everything should just be an Item struct
// */
type Item interface {
	category() string
	name() string
	price() int64
	desc() string
}

type Food struct {
	//implements Item
	category string
	name     string
	price    int64
	desc     string
}

type Drink struct {
	//implements Item
	category string
	name     string
	price    int64
	desc     string
}

func main() {
	apple := Food{name: "Apple"}
	orange := Food{name: "Orange"}
	coffee := Drink{name: "Coffee"}

	println(apple.name, orange.name, coffee.name)
	readDB()

}

func readDB() {
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
		var price int64
		err := rows.Scan(&category, &name, &desc, &price)
		if err != nil {
			panic(err)
		}

		println("ITEM:")
		fmt.Println("Category: ", category)
		fmt.Println("Name: ", name)
		fmt.Println("Description: ", desc)
		fmt.Println("Price: ", price)

		// Process the retrieved data
	}

}
