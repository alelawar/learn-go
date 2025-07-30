package main

import (
	"fmt"
)

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "eko"
	// person["adress"] = "jabaru"

	person := map[string]string{
		"name":   "ale",
		"adress": "jabaru",
	}

	fmt.Println(person["name"])
	fmt.Println(person["adress"])
	fmt.Println(person)

	book := map[string]string{
		"name":   "buku nasgor",
		"author": "ale",
		"ups":    "wrong",
	}

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
