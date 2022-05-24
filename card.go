package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type card struct {
	Number int
	Suit   string
}

func (c card) toString() string {
	return strconv.Itoa(c.Number) + " of " + c.Suit
}

func toCard(str string) card {
	elements := strings.Split(str, " ")
	number, err := strconv.Atoi(elements[0])
	if err != nil {
		printError(err)
		os.Exit(adapterError)
	}
	suit := elements[2]
	if !suits.contains(suit) {
		os.Exit(parseError)
	}
	return card{number, suit}
}

func (c card) display() {
	fmt.Println(c.toString())
}
