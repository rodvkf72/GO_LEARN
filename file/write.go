package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("C:/Users/rodvk/go/src/github.com/rodvkf72/file/writetest.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	num := 10
	if _, err := fmt.Fprintf(f, "%d\n", num); err != nil {
		fmt.Println(err)
	}
}
