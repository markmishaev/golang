package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	/*
		Write a program which prompts the user to enter a string.
		The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
		The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
		The program should print “Not Found!” otherwise.
		The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
		Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
		The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
	*/

	fmt.Print("Enter text: \n")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var iFound, aFound, nFound bool

	for i, c := range input {

		c = unicode.ToLower(c)

		if i == 0 && c == 'i' {
			iFound = true
		}

		if i == len(input)-1 && c == 'n' {
			nFound = true
		}

		if i != 0 && i != len(input)-1 && c == 'a' {
			aFound = true
		}
	}

	if iFound && aFound && nFound {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
