package main

var seq = 0

func Sequence() interface{} {
	seq = seq + 1
	return seq
}