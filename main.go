package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
)

func main() {
	fmt.Println("Hello from Go!")
	os := runtime.GOOS

	if os == "js" {

		// Make request
		resp, err := http.Get("https://raw.githubusercontent.com/uussoop/test/master/main.go")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// Print status and response
		fmt.Println(resp.Status)
		fmt.Println(string(body))
	}
	// Print output
}
