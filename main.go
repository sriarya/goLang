package main

// go is not a Object Oriented Program, so there's no concept of classes.

func main() {
	//go is a static language like  C++ and Java
	//var card string = "Ace of Spades"
	//card := "Ace of Spades" // here the type is being inferred by go
	//and also use : for the very first initialisation
	//No need to use : for reassignment
	// card := newCard()
	// fmt.Println(card)
	//fmt.Println(printState())
	//arrays()

	cards := newDeck()
	// cards.saveToFile("my_cards")
	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println([]byte(cards.toString()))

}

// When retutning something from a function do mention the type of variable being returned
// func newCard() string {
// 	return "Five of Diamonds"
// }

func newCard() int {
	return 32
}
