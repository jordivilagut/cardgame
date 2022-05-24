package main

import (
	"fmt"
	"io/fs"
	"os"
)

var (
	suits             = slice{"Clubs", "Diamonds", "Hearts", "Spades"}
	separator         = ","
	defaultPermission = fs.FileMode(0666)
	readError         = 1
	adapterError      = 2
	parseError        = 3
)

func printError(err error) {
	fmt.Println("Error: ", err)
}

func printErrorAndQuit(err error, code int) {
	printError(err)
	os.Exit(code)
}

func deleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		printError(err)
	}
}
