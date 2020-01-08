package slicevolume

import "fmt"

func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	slice1 := nums[:3]
	fmt.Println(slice1)
	fmt.Println("len:", len(slice1))
	fmt.Println("cap:", cap(slice1))
	fmt.Println()

	slice2 := nums[2:]
	fmt.Println(slice2)
	fmt.Println("len:", len(slice2))
	fmt.Println("cap:", cap(slice2))
	fmt.Println()

	slice3 := nums[:4]
	fmt.Println(slice3)
	fmt.Println("len:", len(slice3))
	fmt.Println("cap:", cap(slice3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, slice1, slice2, slice3)

	//Output:
	//[1 2 3 4 5]
	//len:5
	//cap:5

	//[1 2 3]
	//len:3
	//cap:5

	//[3 4 5]
	//len:3
	//cap:3

	//[1 2 3 4]
	//len:4
	//cap:5

	//[1 2 100 4 5] [1 2 100] [100 4 5] [1 2 100 4]
}
