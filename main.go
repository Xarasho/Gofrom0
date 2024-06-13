package main

import (
	"fmt"

	"github.com/Xarasho/Gofrom0/variables"
)

func main() {
	// variables.ShowIntegers()
	// variables.ShowRest()
	state, text := variables.ConverToText(1588)
	fmt.Println("Text is: ", text, " and State is: ", state)
}
