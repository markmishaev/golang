package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Person class represents Person entity
type Person struct {
	FirstName string
	LastName  string
}

func main() {

	p1 := Person{FirstName: "Mark", LastName: "Mishaev"}

	barr, err := json.Marshal(p1)

	if err == nil {
		fmt.Println(barr)
	}

	var p2 Person
	err1 := json.Unmarshal(barr, &p2)
	if err1 == nil {
		fmt.Println(p2)
	}

	dat, e := ioutil.ReadFile("C:\\temp\\notes1.txt")
	if e == nil {
		fmt.Println(dat)
	}

	e1 := ioutil.WriteFile("C:\\temp\\notes_new.txt", dat, 0777)
	if e1 != nil {
		fmt.Println(e1)
	}

	f, e2 := os.Open("C:\\temp\\notes1.txt")
	if e2 == nil {
		barr1 := make([]byte, 10)
		nb, e3 := f.Read(barr1)
		if e3 == nil {
			fmt.Println(nb)
		}
		f.Close()
	}

	f1, e4 := os.Create("C:\\temp\\notes3.txt")
	if e4 == nil {
		f1.WriteString("Hi Marko Polo")
	}
	f1.Close()
}
