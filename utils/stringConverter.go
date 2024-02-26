package utils

import (
	"strconv"
)

func ConvertStringToUint(s string) (uint, error) {
	result, err := strconv.Atoi(s)
	if err != nil {
		// fmt.Println(err.Error())
		return 0, err
	}
	return uint(result), nil
}
