package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	root, _ := filepath.Abs(".") //absolute
	fmt.Println("Processing path", root)

	err := filepath.Walk(root, processPath)
	if err != nil {

		fmt.Println("error:", err)
	}
}

//doc golang.org/pkg/path/filepath.....index....type Walkfunc
func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)

		}
	}
	return nil
}
