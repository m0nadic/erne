package main

import "fmt"

var seq = 0

func Sequence() string {
	seq = seq + 1
	return fmt.Sprintf("%d", seq)
}