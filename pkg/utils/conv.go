package utils

import "strconv"

func AtoiOrPanic(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}
