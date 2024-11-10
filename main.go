package main

import "fmt"

func main() {
	//go is a static language like  C++ and Java
	//var card string = "Ace of Spades"
	card := "Ace of Spades" // here the type is being inferred by go
	//and also use : for the very first initialisation
	//No need to use : for reassignment
	card = "Five of Diamonds"
	fmt.Println(card)
}
