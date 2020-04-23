package main

import (
	"fmt"
)

//Class classes
type Dog struct {
	Breed  string //uppercase
	Weight int
}

func main() {
	//no super , no sub, no heritage
	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nweight: %v", poodle.Breed, poodle.Weight) //dotted

}
