package main

import (
	"fmt"
)

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_deck")

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
	cards := newDeckFromfile("my_deck")
	cards.print()
	fmt.Println("-------------------------")
	cards.shuffle()
	cards.print()
}
