package main

import (
	"fmt"

	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	//one two -> one
	/* var s string
	fmt.Scanln(&s)
	fmt.Println(s) */

	//one two -> one two
	reader := bufio.NewReader(os.Stdin)
	/*fmt.Print("enter text")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)*/

	//int
	fmt.Print("enter num")
	str, _ := reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value of number =", f)
	}
	fmt.Println(str)
}
