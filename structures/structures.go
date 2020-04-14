package main

import (
	"fmt"
)

//Person type
type Person struct {
	name  string
	addr  string
	phone string
}

func main() {

	//var idMap map[string]int
	//idMap = make(map[string]int)

	//Map literal
	idMap1 := map[string]int{"Joe": 123, "Mark": 1234}
	idMap1["Liza"] = 3456

	fmt.Println(idMap1["Mark"])
	fmt.Println(idMap1["Liza"])

	//delete(idMap1, "Joe")

	for key, value := range idMap1 {
		fmt.Println(key, value)
	}

	var p1 Person
	p1.name = "Mark"
	p1.addr = "Molcho 4"
	p1.phone = "123-2323"

	fmt.Println(p1)

	//p2 := new(Person)
	p2 := Person{name: "Liza", addr: "Lala", phone: "123456"}
	fmt.Println(p2)
}
