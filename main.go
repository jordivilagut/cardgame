package enum

func main() {
	deck := newDeck()
	deck.display()
	deck.shuffle()
	deck.deal()
}
