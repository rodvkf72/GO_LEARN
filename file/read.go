package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("C:/Users/rodvk/go/src/github.com/rodvkf72/file/readtest.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var num int
	if _, err := fmt.Fscanf(f, "%d\n", &num); err == nil {
		fmt.Println(num)
	}
}
