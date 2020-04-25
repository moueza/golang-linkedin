package main

import "fmt"

func main() {
	doSomething()
	sum := addValues(23, 54)
	fmt.Println("sum:", sum)
	sum = addAllValues(12, 54, 79)
	fmt.Println("New sum:", sum)

}
func doSomething() {
	fmt.Println("Doing something")
}

//'inference
func addValues(value1, value2 int) int {
	return value1 + value2
}

//varrangs
//no function overloading
//uppercase=public out package
//lowercase
func addAllValues(values ...int) int {
	sum := 0
	for i := range values { //slice
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
