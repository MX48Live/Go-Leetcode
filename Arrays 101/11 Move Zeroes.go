// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Example 2:
// Input: nums = [0]
// Output: [0]

// Constraints:
// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1

// Follow up: Could you minimize the total number of operations done?

package main

import "fmt"

// var nums = []int{0,1}
var nums = []int{0,1,0,3,0,13,14,-1}
// var nums = []int{0}

func moveZeroes(nums []int)  {
	// Edge Case
	if len(nums) <= 1 { fmt.Println(nums) }
	
	for i := 0; i < len(nums)-1; i++ {

		for j := i+1; j < len(nums); j++ {

			if nums[i] == 0 {
				nums[i] = nums[j]
				nums[j] = 0
			}
		}
	}
	fmt.Println(nums)
}

func main() {
	moveZeroes(nums)
}