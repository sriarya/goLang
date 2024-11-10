package main

import "fmt"

func main() {
	//go is a static language like  C++ and Java
	//var card string = "Ace of Spades"
	//card := "Ace of Spades" // here the type is being inferred by go
	//and also use : for the very first initialisation
	//No need to use : for reassignment
	card := newCard()
	fmt.Println(card)
}

// When retutning something from a function do mention the type of variable being returned
// func newCard() string {
// 	return "Five of Diamonds"
// }

func newCard() int {
	return 32
}
