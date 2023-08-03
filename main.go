package main

import (
	"fmt"
	"runtime"
	"syscall/js"
)

func main() {
	fmt.Println("Hello from Go!")
	os := runtime.GOOS

	if os == "js" {

		// WebAssembly - no system commands

		// Could use JavaScript interop to get process info
		out := js.Global().Get("navigator").Call("getProcessList")

		fmt.Println(out)
	}
	// Print output
}
