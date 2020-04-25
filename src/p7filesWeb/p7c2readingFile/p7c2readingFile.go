package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "./hello.txt"
	content, err := ioutil.ReadFile(filename)
	checkError(err)

	result := string(content) //cf doc src code pkg/io/ioutil/ioutil.go func ReadFile.....readAll (more in src code of Go)
	fmt.Println("Read from file : ", result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}

}
