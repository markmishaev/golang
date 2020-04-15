package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal supertype for holding information about different animal types
type Animal struct {
	food       string
	locomotion string
	noise      string
}

//Eat method return the food of the corresponding animal
func (a *Animal) Eat() string {
	return "Food: " + a.food
}

//Move method return the locomotion of the corresponding animal
func (a *Animal) Move() string {
	return "Locomotion: " + a.locomotion
}

//Speak method return the noise of the corresponding animal
func (a *Animal) Speak() string {
	return "Sound: " + a.noise
}

//Cow function returns new Cow object initialized
func Cow() *Animal {
	return &Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
}

//Bird function returns new Bird object initialized
func Bird() *Animal {
	return &Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

}

//Snake function returns new Snake object initialized
func Snake() *Animal {
	return &Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

}

func main() {

	var animal *Animal

	fmt.Println("Enter the name and action of an animal. ")

	for {

		in := bufio.NewScanner(os.Stdin)
		fmt.Print(">")

		in.Scan()
		input := in.Text()

		if strings.HasPrefix(input, "X") {
			fmt.Println("Good bye!")
			os.Exit(0)
		}

		// “cow”, “bird”, or “snake” ; “eat”, “move”, or “speak”

		arguments := strings.Fields(input)
		animalName := arguments[0]
		animalAction := arguments[1]

		if animalName == "cow" {
			animal = Cow()
		} else if animalName == "bird" {
			animal = Bird()
		} else if animalName == "snake" {
			animal = Snake()
		}

		if animalAction == "eat" {
			fmt.Println(animal.Eat())
		} else if animalAction == "move" {
			fmt.Println(animal.Move())
		} else if animalAction == "speak" {
			fmt.Println(animal.Speak())
		}
	}
}
