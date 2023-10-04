package main

import (
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strconv"
)

// Item /**
// Using an interface for now but wondering if everything should just be an Item struct
// */
type Item interface {
	category() string
	name() string
	price() float64
	desc() string
}

type Food struct {
	//implements Item
	category string
	name     string
	price    float64
	desc     string
}

type Drink struct {
	//implements Item
	category string
	name     string
	price    float64
	desc     string
}

func main() {
	CreateDB()
	//ListAllItems()
	selectAction()
	item := enterItem()
	AddItem(item)
	//ListAllItems()

}

func selectAction() {
	for {
		println("0.) Quit")
		println("1.) Create New Item")
		println("2.) List all existing Items")

		input := Input("Please make a selection >>  ")

		if input == "0" {
			os.Exit(0)
		} else if input == "1" {
			enterItem()
		} else if input == "2" {
			ListAllItems()
		} else {
			println("INVALID SELECTION")
		}
	}
}

func enterItem() (item Food) {
	item.name = Input("Enter item name >> ")
	item.desc = Input("Enter item description >> ")
	price := Input("Enter item price >> ")
	item.category = Input("Enter item category >> ")
	convertedPrice, err := strconv.ParseFloat(price, 64)
	item.price = convertedPrice
	if err != nil {
		return
	}

	return item
}
