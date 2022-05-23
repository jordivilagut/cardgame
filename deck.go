package main

import (
	"fmt"
	"math/rand"
)

type deck []card

func newDeck() deck {
	var d deck
	var suits = [...]string{"Clubs", "Diamonds", "Hearts", "Spades"}

	for _, suit := range suits {
		for j := 1; j < 13; j++ {
			card := card{j, suit}
			d = append(d, card)
		}
	}

	return d
}

func (d deck) display() {
	for _, card := range d {
		card.display()
	}
}

func (d deck) shuffle() {
	fmt.Println("Shuffle the cards.")
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	d.display()
}

func (d deck) deal() {
	fmt.Println("Deal 4 cards.")
	for i := 0; i < 4; i++ {
		d[rand.Intn(len(d))].display()
	}
}
