package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+"of"+" "+suit)
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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		//If there is anything wrong while reading the file, log the error and exit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return []string(s)
}

func (d deck) shuffle() {
	for i := range d {
		//The randomisation that we are trying to acheive is not that good with the below code,
		//as most of the times the same seed is used by go in rand, which generates very similar
		//random pattern of numbers every time we call the function
		//newPosition := rand.Intn(len(d) - 1)

		//So we create a new source , which generates a seed based on current time and the we use it.
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		newPosition := r.Intn(len(d) - 1)
		d[newPosition], d[i] = d[i], d[newPosition]
	}
}
