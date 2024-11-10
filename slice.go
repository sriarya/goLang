package main

import "fmt"

func arrays() {
	cards := []string{"Five of Spades", "Ace of Diamonds"}
	//A slice/array in Go can only have one type of value in it.
	//Meaning either all of them have to be string or int, not both
	cards = append(cards, "Seven of Diamonds")
	// here the append is not using the exisiting cards to add.
	//It is creating new array and adding the values we provided, and then assigns it to cards var

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
