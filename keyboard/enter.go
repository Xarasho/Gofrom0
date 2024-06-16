package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n1 int
var n2 int
var text string
var err error

func EnterNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter first number: ")
	if scanner.Scan() {
		n1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Data is incorrect " + err.Error())
		}
	}

	fmt.Println("Enter second number: ")
	if scanner.Scan() {
		n2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Data is incorrect " + err.Error())
		}
	}

	fmt.Println("Enter text: ")
	if scanner.Scan() {
		text = scanner.Text()
	}

	fmt.Println(text, n1*n2)
}
