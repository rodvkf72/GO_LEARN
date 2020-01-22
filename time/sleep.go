package main

import (
	"fmt"
	"time"
)

func CountDown(seconds int) {
	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}
}

func main() {
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("I am so excited!!")
	})
	/*
		timer := time.AfterFunc(3*time.Second, func() {

		})
		timer.Stop()
	*/
	fmt.Println("Ladies and gentlemen!")
	CountDown(5)
}
