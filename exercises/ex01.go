package exercises

import (
	"strconv"
)

func Ex01(s string) (int, string) {
	int_s, err := strconv.Atoi(s)

	if err != nil {
		return 0, "Error: " + err.Error()
	}

	if int_s > 100 {
		return int_s, "Bigger than 100"
	} else if int_s < 100 {
		return int_s, "Smaller than 100"
	} else {
		return int_s, "num is 100"
	}

}
