package main

import "fmt"

// declare a custom type called deck which extends all the functionality
// of built-in Slice type
type deck []string

/*
Create a receiver function called print to print all cards.
Receiver function are function which attach to custom types.
In our case any variable of type deck has an access of print function.
To declare a receiver function need to add "(variable_name type_name)"

In our case (d deck) where "d" is the actual copy of the deck
and "deck" indicates that every variable of type deck can call this function.
*/
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValue := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
