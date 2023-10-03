package main

import (
	_ "github.com/mattn/go-sqlite3"
	"strconv"
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
	ListAllItems()
	item := enterItem()
	AddItem(item)
	ListAllItems()

}

func enterItem() (item Food) {

	item.name = Input("Enter item name >> ")
	item.desc = Input("Enter item description >> ")
	price := Input("Enter item price >> ")
	item.category = Input("Enter item category >> ")
	convertedPrice, err := strconv.ParseInt(price, 10, 64)
	item.price = convertedPrice
	if err != nil {
		return
	}

	return item
}
