package main

import (
	"fmt"
	"runtime"
)



func main() {
	fmt.Println("Go runs on")
	// swtich clause only execute first case!
	// Other languages need break for each cases, Go doesn' need it.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
