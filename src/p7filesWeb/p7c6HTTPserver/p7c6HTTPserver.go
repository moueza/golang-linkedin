package main

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	//"math/big"
	"net/http"
	//"strings"
)

type Hello struct{}

func (h Hello) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	//my comment
	//F writer obj
	fmt.Fprint(w, "<h1>Hello from the GO server</h1>")

}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	//TODO more : concurrent request
	//Go built with the web in mind
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
