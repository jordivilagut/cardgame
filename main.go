package main

import (
	"os"
)

func main() {

	deck := newDeck()

	saveError := deck.saveToFile("cards")
	if saveError != nil {
		printError(saveError)
	}

	deckFromFile := newDeckFromFile("cards")
	deckFromFile.display()
	deckFromFile.shuffle()
	deckFromFile.deal(5)

	deleteError := os.Remove("cards")
	if deleteError != nil {
		printError(deleteError)
	}
}
