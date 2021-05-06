package main

import "fmt"

// Given a binary array nums, return the maximum number of consecutive 1's in the array.
// var nums = []int{1,1,0,1,1,1}
var nums = []int{1,0,1,1,0,1}

func findMaxConsecutiveOnes(nums []int) int {
	top := 0
	current := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			current++
		} else {			
			current = 0
		}
		if top < current {
			top = current
		}
	}
    return top
}

func main() {
	fmt.Println(findMaxConsecutiveOnes(nums))
}

