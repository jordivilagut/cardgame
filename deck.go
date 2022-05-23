package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
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

func (d deck) toString() string {

	var cards []string
	for _, card := range d {
		cards = append(cards, card.toString())
	}

	return strings.Join(cards, " ")
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

func (d deck) deal(handSize int) {
	fmt.Printf("Deal %d cards.\n", handSize)
	for i := 0; i < handSize; i++ {
		d[rand.Intn(len(d))].display()
	}
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
