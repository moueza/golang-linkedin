package main

import (
	"fmt"
	"stringutil"
)

//cf doc fmt ...error
func main() {
	n1, l1 := stringutil.FullName("Zaphod", "Beeblerox")
	fmt.Printf("Fullname: %v,number of chars: %v\n\n", n1, l1)

	n2, l2 := stringutil.FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("Fullname: %v,number of chars: %v", n2, l2)
}
