package main

import (
	"fmt"
	"strings"
)

//https://golang.org/pkg/strings/
func main() {
	str1 := "implicit string"
	fmt.Println(str1)
	//%T upper
	fmt.Printf("str1: %v:%T\n", str1, str1)

	var str2 string = "explicit string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lvalue := "hello"
	uvalue := "HELLO"
	fmt.Println("Equal?", (lvalue == uvalue))
	fmt.Println("Equal?", (lvalue == lvalue))
	fmt.Println("Equal?", strings.EqualFold(lvalue, uvalue))

	fmt.Println("Equal?", strings.Contains("I am a priest", "am"))
	fmt.Println("Equal?", strings.Contains("I am a priest", "amkjlg"))

}


