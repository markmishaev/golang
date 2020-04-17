package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go printMessage(&wg)

	wg.Wait()

	fmt.Println("Main Routine")
}

func printMessage(wg *sync.WaitGroup) {
	fmt.Println("New Routine")
	wg.Done()
}
