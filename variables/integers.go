package variables

import "fmt"

func ShowIntegers() {
	var common_int int
	int_32 := int32(10)
	int_64 := int64(100)

	fmt.Println("common_int = ", common_int)
	fmt.Println("int_32 = ", int_32)
	fmt.Println("int_64 = ", int_64)
}
