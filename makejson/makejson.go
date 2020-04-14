package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
	Write a program which prompts the user to first enter a name, and then enter an address.
	Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
	Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

func main() {

	personMap := make(map[string]string)

	in := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter name: ")
	in.Scan()
	name := in.Text()

	fmt.Print("Enter address: ")
	in.Scan()
	address := in.Text()

	personMap[name] = address

	barr, err := json.Marshal(personMap)

	if err == nil {
		jsonStr := string(barr)
		fmt.Println(jsonStr)
	}
}
