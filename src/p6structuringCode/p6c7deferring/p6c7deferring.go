package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Close the file")
	fmt.Println("Open the file")

	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")

	myFunc()

	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	fmt.Println("Undeferred statement")

	x := 1000
	defer fmt.Println("val of x", x)
	x++
	defer fmt.Println("val of x after incrementing", x)

}

func myFunc() {
	defer fmt.Println("deferred in the function")
	fmt.Println("not deferred in the function")

}
