package main

import (
	"fmt"
	"strconv"
)

// Given an array nums of integers, return how many of them contain an even number of digits.
// Input: nums = [555,901,482,1771]
// Output: 1
// Explanation:
// Only 1771 contains an even number of digits.

var nums = []int{12,345,2,6,7896}

func findNumbers(nums []int) int {
	evenNumber := 0

	for i := 0; i < len(nums); i++ {

		evenDigit := len(strconv.Itoa(nums[i]))
		if evenDigit % 2 == 0 {
			evenNumber++
		} 
	}
	return evenNumber
}

func main() {
	fmt.Println(findNumbers(nums))
}
