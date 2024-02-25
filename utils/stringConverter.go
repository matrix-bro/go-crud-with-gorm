package utils

import (
	"fmt"
	"strconv"
)

func ConvertStringToUint(s string) (uint, string) {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err.Error())
		return 0, "Invalid data"
	}
	return uint(result), ""
}
