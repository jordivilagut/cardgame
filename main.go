package main

func main() {

	deck := newDeck()
	deck.saveToFile("cards")
	deckFromFile := newDeckFromFile("cards")
	deckFromFile.shuffle()
	deckFromFile.deal(2)
	deckFromFile.shuffle()
	deckFromFile.deal(2)
	deckFromFile.shuffle()
	deckFromFile.deal(2)
	deleteFile("cards")
}
