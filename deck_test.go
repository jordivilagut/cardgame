package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 48 {
		t.Errorf("Invalid amount of cards: %d", len(deck))
	}
}

func TestNewDeckFromFile(t *testing.T) {
	deck := newDeck()
	deck.saveToFile("cards")
	deckFromFile := newDeckFromFile("cards")

	if len(deckFromFile) != len(deck) {
		t.Error("Deck from file has different length.")
	}

	if deck[0] != deckFromFile[0] {
		t.Error("First element differs in deck from file.")
	}

	deleteFile("cards")
}
