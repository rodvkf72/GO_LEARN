package slice

import "fmt"

func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	//Output:
	//[1 2 3 4 5]
	//[2 3]
	//[3 4 5]
	//[1 2 3]
}

//minus can not be used
//fruits[:-1] (X)
//fruits[:len(fruits)-1] (O)
