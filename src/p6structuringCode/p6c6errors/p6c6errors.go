package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//error=interface
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		//fmt.Println(err)
		fmt.Println(err.Error())
	}
	myError := errors.New("My error")
	fmt.Println(myError)

	//comma OK syntax
	attendance := map[string]bool{
		"Ann":  true,
		"Mike": true}

	attended, ok := attendance["Mike"] //try M...no info
	if ok {
		fmt.Println("Mike attented?", attended)
	} else {
		fmt.Println("no info for Mike")
	}
}

//no inheritance...unique name
