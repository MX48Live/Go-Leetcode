package main

import (
	"fmt"
	"math"
	"sort"
)

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].

var nums = []int{-7,-3,2,3,11}

func sortedSquares(nums []int) []int {
    for i, val := range nums {

			//math.Abs Function reuqire type float64 
			f64 := math.Abs(float64(val))
			
			//convert f64 to int because nums[] is int type
 			i32 := int(f64)

			nums[i] = i32*i32
		}
		
		sorting := nums
		sort.Ints(sorting)
		return sorting
}

func main() {
	fmt.Println(sortedSquares(nums))
}