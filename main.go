package main

import (
	"fmt"
	"runtime"
	// "github.com/Xarasho/Gofrom0/variables"
)

func main() {
	// variables.ShowIntegers()
	// variables.ShowRest()
	// state, text := variables.ConverToText(1588)
	// fmt.Println("Text is: ", text, " and State is: ", state)

	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("This isn't Windows, it's: ", os)
	} else {
		fmt.Println("This is Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("This is linux")
	case "darwin":
		fmt.Println("This is darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
