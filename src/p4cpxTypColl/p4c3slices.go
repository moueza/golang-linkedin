package main

import (
	"fmt"
	"sort"
)

func main() {
	//var colors=[]
	//var colors=[5]
	var colors = []string{"red", "green", "blue"}

	//fmt.Println(len(numbers))

	colors = append(colors, "added")

	//remove
	//colors = append(colors[1:len(colors)])
	//colors = append(colors[1:])
	colors = []string{"red", "green", "blue"}
	colors = append(colors[:len(colors)-1])

	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 0
	numbers[1] = 10

	numbers[2] = 2
	numbers[3] = 3
	numbers[4] = 4

	numbers = append(numbers, 235)
	fmt.Println(cap(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)

}
