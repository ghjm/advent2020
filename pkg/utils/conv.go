package utils

import "strconv"

func AtoiOrPanic(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}

func AtoiOrPanic64(s string) int64 {
	r, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return r
}
