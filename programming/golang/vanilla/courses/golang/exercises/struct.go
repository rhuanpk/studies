//go:build struct

/*
	Create some program that:
		- create a struct house with variables noRooms, price and city.
		- how does a struct differ from a class?
			R: classes are a concept of object oriented programming. A class can be used to create bjects. Classes can inherit from another.

	Source: https://golangr.com/struct/
*/

package main

import (
	"fmt"
)

func main() {

	type House struct {
		Rooms int64
		Price float64
		City  string
	}

	house := House{
		Rooms: 2,
		Price: 237937.99,
		City:  "ottawa",
	}

	text := "*** HOUSE ***\n\n"
	text += "- Rooms: %d unities.\n"
	text += "- Price: $ %.2f.\n"
	text += "- City: %s.\n"

	fmt.Printf(text, house.Rooms, house.Price, house.City)

}
