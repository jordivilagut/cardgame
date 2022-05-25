package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 48 {
		t.Errorf("Invalid amount of cards: %d", len(deck))
	}

	if deck[0].Suit != "Clubs" {
		t.Errorf("First element suit should be Clubs and got: %s", deck[0].Suit)
	}

	if deck[0].Number != 1 {
		t.Errorf("First element suit should be 1 and got: %d", deck[0].Number)
	}
}

func TestNewDeckFromFile(t *testing.T) {

	deleteFile("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	deckFromFile := newDeckFromFile("_decktesting")

	if len(deckFromFile) != len(deck) {
		t.Error("Deck from file has different length.")
	}

	if deck[0] != deckFromFile[0] {
		t.Error("First element differs in deck from file.")
	}

	deleteFile("_decktesting")
}
