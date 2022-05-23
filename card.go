package main

import (
	"fmt"
	"strconv"
)

type card struct {
	Number int
	Suit   string
}

func (c card) toString() string {
	return strconv.Itoa(c.Number) + " of " + c.Suit
}

func (c card) display() {
	fmt.Println(c.toString())
}
