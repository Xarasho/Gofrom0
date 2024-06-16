package exercises

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)


func Ex02() {
	scanner := bufio.NewScanner(os.Stdin)

	var num int
	var err error

	for {
		fmt.Println("Please enter a number: ")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Incorrect Data, please try again")
				continue
			} else {
				break
			}
		}
	}

	for i:=1; i<=10; i++ {
		fmt.Printf("%d x %d = %d \n", num, i, i*num)
	}
}