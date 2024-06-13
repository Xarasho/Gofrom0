package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Salary float64
var State bool
var Date time.Time

func ShowRest() {
	Name = "Kenny"
	Salary = 1577.66
	State = true
	Date = time.Now()
	fmt.Println(Name)
	fmt.Println(Salary)
	fmt.Println(State)
	fmt.Println(Date)
}

func ConverToText(n int) (bool, string) {
	var text string
	text = strconv.Itoa(n)
	return true, text
}
