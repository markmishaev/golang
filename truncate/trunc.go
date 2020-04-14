package main

import (
	"fmt"
	"os"
)

func main() {

	/*
		Write a program which prompts the user to enter a floating point number and prints the integer,
		which is a truncated version of the floating point number that was entered.
		Truncation is the process of removing the digits to the right of the decimal place.
	*/

	fmt.Print("Please enter real number: \n")
	var numberToTruncate float64

	_, err := fmt.Scanln(&numberToTruncate)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var truncatedNumber int = int(numberToTruncate)
	fmt.Println(truncatedNumber)
}
