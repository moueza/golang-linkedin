package main

import (
	"fmt"
	"sort"
)

func main() {
	//map : often String to other than String
	states := make(map[string]string)
	states["OR"] = "Oregon"
	states["CA"] = "California"
	states["WA"] = "Washington"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	fmt.Println(states)

	//delete(states,'OR') //quotes because string declared ; KO
	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"
	//iterate for
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)

	}
	//slice with no capacity here
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted")
	for i := range keys {
		fmt.Println(states[keys[i]])

	}

}
