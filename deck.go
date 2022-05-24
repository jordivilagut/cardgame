package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []card

func newDeck() deck {
	var d deck

	for _, suit := range suits {
		for j := 1; j < 13; j++ {
			card := card{j, suit}
			d = append(d, card)
		}
	}

	return d
}

func newDeckFromFile(filename string) deck {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		printError(err)
		printErrorAndQuit(err, readError)
	}

	cardStrings := strings.Split(string(bytes), separator)
	var d deck
	for _, str := range cardStrings {
		d = append(d, toCard(str))
	}

	return d
}

func (d deck) toString() string {

	var cards []string
	for _, card := range d {
		cards = append(cards, card.toString())
	}

	return strings.Join(cards, separator)
}

func (d deck) toByteArray() []byte {
	return []byte(d.toString())
}

func (d deck) display() {
	for _, card := range d {
		card.display()
	}
}

func (d deck) shuffle() {
	fmt.Println("Shuffle the cards.")
	random().Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	d.display()
}

func (d deck) deal(handSize int) {
	fmt.Printf("Deal %d cards.\n", handSize)
	for i := 0; i < handSize; i++ {
		d[random().Intn(len(d)-1)].display()
	}
}

func (d deck) saveToFile(filename string) {
	err := ioutil.WriteFile(filename, d.toByteArray(), defaultPermission)
	if err != nil {
		printError(err)
	}
}
