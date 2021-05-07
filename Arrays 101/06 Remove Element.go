package main

import (
	"fmt"
)

// Example 1:
// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2]
// Explanation: Your function should return length = 2, with the first two elements of nums being 2.
// It doesn't matter what you leave beyond the returned length. For example if you return 2 with nums = [2,2,3,3] or nums = [2,2,0,0], your answer will be accepted.

// Example 2:
// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3]
// Explanation: Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4. Note that the order of those five elements can be arbitrary. It doesn't matter what values are set beyond the returned length.

// var nums = []int{3,2,2,3}
// var val int = 3

var nums = []int{0,1,2,2,3,0,4,2}
var val int = 2

// var nums = []int{3,3,3,3,3,3,3,3}
// var val int = 3

// var nums = []int{0,1,2,3,4,5,4,3}
// var val int = 3

// var nums = []int{4,5}
// var val int = 4

// var nums = []int{3,3,3,3,3,3,3,4}
// var val int = 3


func removeElement(nums []int, val int) int {	
	// IMPORTANT !!!!!!
	// IMPORTANT !!!!!!
	// This Test REQUIRED to 'value in-place' mean you can't copy, decrease, increase array.
	// I took 1 day to do this but not success.
	// So I founded Answer in Stack Overflow and try to understand
	// I turn out. You just need to remove 'Val' number from array.
	// and return how new length of array by ignore 'Val' number
	// So to meet the requirement you just need to change 'Val' number
	// to any number.
	
	// Only my stupid here is. I have read instruction not clear enught before try to slove this. 
	
	newLenght := 0;
	fmt.Printf("Started Array: %v\n", nums)
	for _, v := range nums {

		if val != v {
			nums[newLenght] = v
			fmt.Printf("----New Array: %v\n", nums)
			newLenght++
		}
	}
	return newLenght
}

func main() {
	fmt.Println(removeElement(nums, val))
}
