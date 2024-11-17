package main

import "fmt"

type pack []string

func arrays() {
	//cards := []string{"Five of Spades", "Ace of Diamonds"}
	// We can either define the slice as shown above, or prefne the type and then resuse
	cards := pack{"Five of Spades", "Ace of Diamonds"}
	//A slice/array in Go can only have one type of value in it.
	//Meaning either all of them have to be string or int, not both
	cards = append(cards, "Seven of Diamonds")
	// here the append is not using the exisiting cards to add.
	//It is creating new array and adding the values we provided, and then assigns it to cards var

	cards.print()
}

// So what we have below is a receiver function.
// Here its possible to define methods that can be applied to slice, only if we define a type
// for the elements in the slice.
func (p pack) print() {
	for i, card := range p {
		fmt.Println(i, card)
	}
}
