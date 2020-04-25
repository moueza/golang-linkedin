package main

import (
	"fmt"
)

//doc fmt ....interface cf Errorf
//every type implements at least an interface, be it empty
//diamon resolution
type Animal interface {
	Speak() string
	//Speak(volume int) string //will fail because no matching
}

//interface : no keyword implements
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof"
}

type Cat struct {
}

func (d Cat) Speak() string {
	return "Miaou"
}

type Cow struct {
}

func (d Cow) Speak() string {
	return "Meuh"
}

func main() {
	poodle := Animal(Dog{}) //cast just for sure
	fmt.Println(poodle)

	//slice
	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

//no inheritance...unique name
