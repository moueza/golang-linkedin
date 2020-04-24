package main

import (
	"fmt"
)

func main() {
	var x float64 = 0 //+
	var result string
	if x < 0 {
		result = "less than 0"
	} else if x == 0 {
		result = "= 0"

	} else { //: curl linked+++
		result = "greater than 0"
	}
	fmt.Println("Result:", result)

	var result2 string
	if y := 42; y < 0 {
		result = "y less than 0"
	} else if y == 0 {
		result2 = "y = 0"

	} else { //: curl linked+++
		result2 = " y greater than 0"

	}
	fmt.Println("Result:", result2)
	//fmt.Println("y:", y)//error not defined

}
