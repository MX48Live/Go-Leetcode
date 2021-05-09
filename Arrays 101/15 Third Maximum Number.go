// Given integer array nums, return the third maximum number in this array. If the third maximum does not exist, return the maximum number.

// Example 1:
// Input: nums = [3,2,1]
// Output: 1
// Explanation: The third maximum is 1.

// Example 2:
// Input: nums = [1,2]
// Output: 2
// Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

// Example 3:
// Input: nums = [2,2,3,1]
// Output: 1
// Explanation: Note that the third maximum here means the third maximum distinct number.
// Both numbers with value 2 are both considered as second maximum.

// Constraints:
// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1

// Follow up: Can you find an O(n) solution?

package main

import (
	"fmt"
)

var nums = []int{2,2,2,3,4}
// var nums = []int{5,2,2}
// var nums = []int{3,2,1}
// var nums = []int{1,2}
// var nums = []int{3}
// var nums = []int{2,2,3,1}

func thirdMax(nums []int) int {
	max1 := 0
	max2 := 0
	max3 := 0

	for _, num := range nums {

		if(num > max1) {
			max3 = max2
			max2 = max1
			max1 = num
		} else if (num == max1) {
			continue
		} else if(num > max2) {
			max3 = max2
			max2 = num
		} else if (num == max2) {
			continue
		} else if (num > max3) {
			max3 = num
		} else if(num == max3) {
			continue
		}
	}
	if(max3 == 0) {
		return max1
	}
	return max3
}

func main() {
	fmt.Println(thirdMax(nums))
}