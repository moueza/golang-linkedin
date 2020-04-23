package main

import (
	"fmt"
)

func main() {
	//doc   pkg/runtime
	//https://talks.golang.org/2015/go-gc.pdf

	//new (allocate, not initialize) vs make(allocate,  initialize)

	//panic assign
	/*var m map[string]int
	m["key"] = 42*/

	m := make(map[string]int)
	m["key"] = 42

	fmt.Println(m)

	//GC desallocate
}
