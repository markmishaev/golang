package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal interface represents different animal types
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow type to implement Animal interfaces
type Cow struct {
	name string
}

//Bird type to implement Animal interfaces
type Bird struct {
	name string
}

//Snake type to implement Animal interfaces
type Snake struct {
	name string
}

//Eat method of Cow type
func (c Cow) Eat() {
	fmt.Println("grass")
}

//Move method of Cow type
func (c Cow) Move() {
	fmt.Println("walk")
}

//Speak method of Cow type
func (c Cow) Speak() {
	fmt.Println("moo")
}

//Eat method of Bird type
func (b Bird) Eat() {
	fmt.Println("worms")
}

//Move method of Bird type
func (b Bird) Move() {
	fmt.Println("fly")
}

//Speak method of Bird type
func (b Bird) Speak() {
	fmt.Println("peep")
}

//Eat method of Snake type
func (s Snake) Eat() {
	fmt.Println("mice")
}

//Move method of Snake type
func (s Snake) Move() {
	fmt.Println("slither")
}

//Speak method of Snake type
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	var a Animal

	for {

		in := bufio.NewScanner(os.Stdin)
		fmt.Print(">")

		in.Scan()
		input := in.Text()

		if strings.HasPrefix(input, "X") {
			fmt.Println("Good bye!")
			os.Exit(0)
		}

		//  The first string is “newanimal”.
		//  The second string is an arbitrary string which will be the name of the new animal.
		// The third string is the type of the new animal, either “cow”, “bird”, or “snake”.

		//The first string is “query”.
		//The second string is the name of the animal.
		//The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.

		arguments := strings.Fields(input)
		command := arguments[0]
		animalName := arguments[1]
		commandType := arguments[2]

		if command == "newanimal" {
			if commandType == "cow" {
				a = Cow{name: animalName}
				fmt.Println("Created it")
			} else if commandType == "bird" {
				a = Bird{name: animalName}
				fmt.Println("Created it")
			} else if commandType == "sneak" {
				a = Snake{name: animalName}
				fmt.Println("Created it")
			}
		} else if command == "query" {
			if animalName == "cow" {
				a = Cow{name: animalName}
			} else if animalName == "bird" {
				a = Bird{name: animalName}
			} else if animalName == "sneak" {
				a = Snake{name: animalName}
			}

			if commandType == "eat" {
				a.Eat()
			} else if commandType == "move" {
				a.Move()
			} else if commandType == "speak" {
				a.Speak()
			}
		}

	}
}
