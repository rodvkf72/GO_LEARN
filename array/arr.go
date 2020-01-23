package array

import (
	"fmt"
)

func Example_array() {
	fruits := [3]string{"사과", "바나나", "배"}
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruit)
		//for i, fruit := range fruits {
		//	fmt.Println(i, fruit)
	}
	//Output:
	//사과는 맛있다.
	//바나나는 맛있다.
	//배는 맛있다.
}
