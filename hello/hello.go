package main

import "fmt"

func main() {

	//this is my first comment in GO
	//Hello GO!!!

	/*

	   	const x1 = 1.3
	   	const y1 = "234234"

	   	const (
	   		l1 = 15
	   		z1 = 1.7
	   	)

	   	ptr := new(int)
	   	*ptr = 3

	   	type Grades int
	   	const (
	   		A Grades = iota
	   		B
	   		C
	   		D
	   		F
	   	)

	   	for i := 0; i < 10; i++ {
	   		fmt.Printf("hi ")
	   	}

	   	var i = 0

	   	for i < 20 {
	   		fmt.Printf("hey ")
	   		i++
	   	}


	   	for {
	    		fmt.Printf("hey1 ")
	   	}

	*/

	var appleNum int
	fmt.Printf("Please enter number of apples: ")

	num, err := fmt.Scan(&appleNum)

	if num > 0 && err != nil {
		fmt.Printf("%d", appleNum)
	}

}
