package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	firstdeck, seconddeck := deal(cards,5)
	fmt.Println()
	fmt.Println("first deck")
	
	firstdeck.print()
	fmt.Println()
	fmt.Println("Second deck")
	seconddeck.print()


}

func newcard() string {
	return "club"
}