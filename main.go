package main

func main() {

	deck := newDeck()
	deck.saveToFile("cards")
	deckFromFile := newDeckFromFile("cards")
	deckFromFile.display()
	deckFromFile.shuffle()
	deckFromFile.deal(5)
	deleteFile("cards")
}
