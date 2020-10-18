package main

import (
	"strconv"
	"strings"
)

func expandArray(value string) interface{} {
	t := strings.Replace(value, "[", "", -1)
	v := strings.Replace(t, "]", "", -1)
	cmpts := strings.Split(v, ":")
	size, cmd := computeSizeAndGen(cmpts)
	genFn := genMap[cmd]
	retVal := make([]interface{}, 0)

	for i := 0; i < size; i++ {
		retVal = append(retVal, genFn())
	}

	return retVal
}

func computeSizeAndGen(cmpts []string) (int, string) {
	switch len(cmpts) {
	case 0:
		return 0, ""
	case 1:
		return 5, cmpts[0]
	case 2:
		n, _ := strconv.Atoi(cmpts[1])
		return n, cmpts[0]
	default:
		return 5, cmpts[1]
	}
}
