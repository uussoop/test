package main

import (
	"fmt"
	o "os"
	"runtime"
)

func main() {
	fmt.Println("Hello from Go!")
	os := runtime.GOOS

	if os == "js" {

		path := "./folder"

		files, err := o.ReadDir(path)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
			if file.IsDir() {
				fmt.Println("[DIR]")
			}
		}
	}
	// Print output
}
