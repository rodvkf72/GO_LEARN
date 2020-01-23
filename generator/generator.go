package main

import "fmt"

func NewIntGeneratior() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func main() {
	gen := NewIntGeneratior()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
}
