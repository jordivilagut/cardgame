package main

func main() {
	deck := newDeck()
	deck.display()
	deck.shuffle()
	deck.deal(5)
}
