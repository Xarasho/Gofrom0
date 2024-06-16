package iterators

import (
	"fmt"
)

func Iter() {
	for i:=10; i>0; i--{
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}