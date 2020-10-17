package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v5"
	"os"
	"strings"
)

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
		fn := genMap[xs[1]]
		m[xs[0]] = fn()
	}

	data, _ := json.MarshalIndent(&m, "", "    ")
	fmt.Println(string(data))
}
