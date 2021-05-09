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
	"math"
)

var nums = []int{2,2,2,3,4}
// var nums = []int{5,2,2}
// var nums = []int{3,2,1}
// var nums = []int{1,2}
// var nums = []int{3}
// var nums = []int{2,2,3,1}
func thirdMax(nums []int) int {
	max, second, third := math.MinInt64, math.MinInt64, math.MinInt64
    
    for _, v := range nums {        
		if v == max || v == second || v == third {
			continue
		}
        
        switch {
        case v > max:
            max, second, third = v, max, second
        case v > second:
            second, third = v, second
        case v > third:
            third = v
        }
	}
    
	if third == math.MinInt64 {
		return max
	}
    
	return third
}

func main() {
	fmt.Println(thirdMax(nums))
}