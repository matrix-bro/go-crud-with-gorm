package utils

import (
	"fmt"
	"strconv"
)

func ConvertStringToUint(s string) (uint, error) {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("please enter valid ID")
	}
	return uint(result), nil
}
