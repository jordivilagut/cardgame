package main

import "fmt"

type card struct {
	Number int
	Suit   string
}

func (c card) display() {
	fmt.Println(c)
}
