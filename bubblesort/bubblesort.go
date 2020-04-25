package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a Bubble Sort program in Go.
The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.
As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.
A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
*/

func main() {

	in := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a sequence of up to 10 integers: ")
	in.Scan()

	var numbersSlice []int

	integersSequence := in.Text()
	ints := strings.Fields(integersSequence)

	for _, numStr := range ints {
		numInt, err := strconv.Atoi(numStr)
		if err == nil {
			numbersSlice = append(numbersSlice, numInt)
		}
	}

	fmt.Println("Numbers before sorting: ")
	fmt.Println(numbersSlice)

	BubbleSort(numbersSlice)

	fmt.Println("Numbers after sorting: ")
	fmt.Println(numbersSlice)
}

func BubbleSort(tosort []int) {
	size := len(tosort)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if tosort[j] < tosort[j-1] {
				Swap(tosort, j)
			}
		}
	}
}

//Swap method swaps the position of two adjacent elements in the slice
func Swap(tosort []int, i int) {
	tosort[i], tosort[i-1] = tosort[i-1], tosort[i]
}
