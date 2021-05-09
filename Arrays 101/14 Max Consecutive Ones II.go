// Max Consecutive Ones II

// Solution
// Given a binary array nums, return the maximum number of consecutive 1's in the array if you can flip at most one 0.

// Example 1:
// Input: nums = [1,0,1,1,0]
// Output: 4
// Explanation: Flip the first zero will get the maximum number of consecutive 1s. After flipping,
// the maximum number of consecutive 1s is 4.

// Example 2:
// Input: nums = [1,0,1,1,0,1]
// Output: 4

// Constraints:
// 1 <= nums.length <= 105
// nums[i] is either 0 or 1.

// Follow up: What if the input numbers come in one by one as an infinite stream? In other words,
// you can't store all numbers coming from the stream as it's too large to hold in memory.
// Could you solve it efficiently?

package main

import "fmt"

// var nums = []int{1,0,0,1,1,1}
// var nums = []int{1,1,0,1}
var nums = []int{0}
// var nums = []int{1}

func findMaxConsecutiveOnes(nums []int) int {

	if len(nums) == 0 { return 0 }
	if len(nums) == 1 { return 1 }

	left := 0
	right := 0
	zeros := 0
	zeroIndex := -1
	maxOnes := 0

	for right < len(nums) {
		if nums[right] == 1 {
			right++
		} else {
			if zeros >= 1 {
				left = zeroIndex + 1
				zeros = 0
			}
			zeroIndex = right
			right++
			zeros++
		}
		maxOnes = max(maxOnes, right - left)

	}
  return maxOnes
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
	fmt.Println(findMaxConsecutiveOnes(nums))
}
