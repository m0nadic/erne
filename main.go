package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v5"
	"strings"
)

const arrayIndicator = "["

func main() {
	count := flag.Int("n", 1, "number of records to generate")

	flag.Parse()

	args := flag.Args()
	gofakeit.Seed(0)
	for i := 0; i < *count; i++ {

		generateRecord(args)
	}
}

func generateRecord(args []string) {
	m := make(map[string]interface{})
	for _, a := range args {

		key, value := splitKeyVal(a)

		fn, ok := genMap[value]
		if ok {
			m[key] = fn()
		} else {
			m[key] = expand(value)
		}

	}

	data, _ := json.MarshalIndent(&m, "", "    ")
	fmt.Println(string(data))
}

func splitKeyVal(a string) (string, string) {
	xs := strings.Split(a, "=")

	switch len(a) {
	case 0:
		return "", ""
	case 1:
		return xs[0], ""
	default:
		return xs[0], xs[1]
	}
}

func expand(value string) interface{} {
	if strings.HasPrefix(value, arrayIndicator) {
		return expandArray(value)
	}

	return value
}
