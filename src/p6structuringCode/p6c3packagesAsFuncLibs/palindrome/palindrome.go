package main

import "fmt"

//cf doc fmt ...error
func main() {
	n1, l1 := FullName("Zaphod", "Beeblerox")
	fmt.Printf("Fullname: %v,number of chars: %v\n\n", n1, l1)

	n2, l2 := FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("Fullname: %v,number of chars: %v", n2, l2)
}

//uppercase
func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length //order important
}

//order not important
func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l //=because already in signature
	length = len(full)
	return
}
