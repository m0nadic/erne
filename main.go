package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v5"
	"os"
	"strings"
)
const arrayIndicator = "["

func main() {
	count := flag.Int("n", 1, "number of records to generate")

	flag.Parse()

	gofakeit.Seed(0)
	for i := 0; i < *count; i++ {
		generateRecord(os.Args[1:])
	}
}

func generateRecord(args []string) {
	m := make(map[string]interface{})
	for _, a := range args {

		xs := strings.Split(a, "=")
		if len(xs) != 2 {
			continue
		}

		key := xs[0]
		value := xs[1]

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

func expand(value string) interface{} {
	if strings.HasPrefix(value, arrayIndicator) {
		return expandArray(value)
	}

	return value
}


