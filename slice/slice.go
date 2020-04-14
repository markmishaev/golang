package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	/*
		Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
		The program should be written as a loop.
		Before entering the loop, the program should create an empty integer slice of size (length) 3.
		During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
		The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
		The slice must grow in size to accommodate any number of integers which the user decides to enter.
		The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
	*/

	fmt.Println("Hello dear user. Enter numbers to be added to slice, X to exit")

	userInputSlice := make([]int, 3)
	inputSize := 0

	for {

		in := bufio.NewScanner(os.Stdin)
		fmt.Print(">")

		in.Scan()
		input := in.Text()

		if strings.HasPrefix(input, "X") {
			fmt.Println("Good bye!")
			os.Exit(0)
		}

		numInput, err := strconv.Atoi(input)

		if err == nil {
			if inputSize < len(userInputSlice) {
				userInputSlice[inputSize] = numInput
			} else {
				userInputSlice = append(userInputSlice, numInput)
			}
			if inputSize > 1 {
				sort.Ints(userInputSlice)
			}
			printSlice(userInputSlice)
			inputSize++
		} else {
			fmt.Println(err)
		}
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
