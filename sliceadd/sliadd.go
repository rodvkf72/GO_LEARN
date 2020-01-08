package sliceadd

import (
	"fmt"
)

func Example_append() {
	f1 := []string{"사과", "바나나", "배"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)
	f4 := append(f1[:2], f2...)

	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
}

//f3 := append(f1, f2) (x)
//append is variable factor so you can not use slice(f2)
//slice -> append = ... ex) f1 -> f1...
