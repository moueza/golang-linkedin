package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	////dow := rand.Intn(6) + 1 //range 0 to x
	//fmt.Println("Day:", dow)

	result := ""

	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "it s sunday"
	case 7:
		result = "it s saturday"
	default:
		result = "it s weekday"
	}

	//fmt.Println("Day:", dow, ",", result)
	fmt.Println(result)

	x := -42

	//if like, no break
	switch {
	case x < 0:
		result = "less than 0"
		//fallthrough reset and replace by following
	case x == 0:
		result = "= 0"
	case x > 0:
		result = "greater than 0"
	}
	fmt.Println(result)

}
