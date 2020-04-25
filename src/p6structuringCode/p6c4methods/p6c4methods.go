package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

//attached
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v! \n", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func (d *Dog) PointerThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 17, "Woof"} //curly braces for struct initialization https://stackoverflow.com/questions/30358313/compiler-too-many-arguments-given-despite-that-all-are-given , not parenthesis
	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound = "Arf"
	poodle.Speak()

	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes() //does not modify : as value, not pointer

	poodle.PointerThreeTimes()
	poodle.PointerThreeTimes()
}

//no inheritance...unique name
