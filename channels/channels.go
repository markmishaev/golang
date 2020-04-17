package main

import "fmt"

func main() {

	c := make(chan int)

	go prod(1, 2, c)
	go prod(3, 4, c)

	a := <-c
	b := <-c

	fmt.Println(a * b)
}

func prod(n1 int, n2 int, c chan int) {
	c <- n1 * n2
}
