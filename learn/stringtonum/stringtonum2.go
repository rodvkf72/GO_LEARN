package main

import "fmt"

func main() {
	var num int
	var s string

	fmt.Sscanf("57", "%d", &num)
	fmt.Println(num)
	s = fmt.Sprint(3.14)
	fmt.Println(s)
	s = fmt.Sprintf("%x", 13402077)
	fmt.Println(s)
}
