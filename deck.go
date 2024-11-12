package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"of"+suit)
		}
	}

	return cards
}

// creating a receiver function on cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal is not a receiver function because,
// it doesn't define that the function call needs to have a type
func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}
