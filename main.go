package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v5"
	"os"
	"strings"
)

func main() {
	gofakeit.Seed(0)

	m := make(map[string]string)
	for _, a := range os.Args[1:] {
		xs := strings.Split(a, "=")
		fn := genMap[xs[1]]
		m[xs[0]] = fn()
	}

	data, _ := json.MarshalIndent(&m, "", "    ")
	fmt.Println(string(data))
}
