/*
In cards workspace there are two files with package main.
In such scenarios run <go run *.go> instead of just <go run main.go>
go run *.go will load all go files in the package and run the main().

*** It's not recommended to have more than one file with package main
In same workspace. Only main.go should have package main.
*/

/*
	Difference between array and slice
	Array --> Fixed length list
	Slice --> An Array that can grow or shrink.

	Every record inside and array or slice must be in same type.
*/

package main

// import "fmt"

func main() {
	// type deck []string
	// var card string = "ace of spades"
	// card := "ace of spades"   //declare a variable and assign a value to it
	// card = "Five of diamonds" //variable assignment
	// card := newCard()
	// fmt.Println(card)

	// Create a Slice of type string and assign value
	// cards := []string{newCard(), "ace of spades"}

	// loop over card slice
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	//Create a Slice of type deck and assign value
	//creating cards using custom deck type
	// cards := deck{newCard(), "ace of spades"}
	// fmt.Println(cards)
	cards := newDeck()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

// func newCard() string {
// 	return "Five of diamonds"
// }
