package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).
Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

//Person type
type Person struct {
	fname string
	lname string
}

func main() {

	var personSlice []Person

	in := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the name of file: ")
	in.Scan()

	fileName := in.Text()
	fileName = strings.Trim(fileName, "\"")

	fmt.Println(fileName)

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {

		fmt.Println(eachline)

		fields := strings.Fields(eachline)
		fmt.Println(fields, len(fields))

		p := Person{fname: fields[0], lname: fields[1]}

		personSlice = append(personSlice, p)

		printSlice(personSlice)
	}
}

func printSlice(s []Person) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
