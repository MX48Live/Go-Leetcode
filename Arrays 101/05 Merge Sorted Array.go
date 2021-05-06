package main

import (
	"fmt"
	"sort"
)

// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has a size equal to m + n such that it has enough space to hold
// additional elements from nums2.

// Example 1:
// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]

// Example 2:
// Input: nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]

var nums1 = []int{1,2,3,0,0,0}
var m int = 3
var nums2 = []int{2,5,6}
var n int = 3

func merge(nums1 []int, m int, nums2 []int, n int)  {

		//Slice it down to m , n
		nums2 = append(nums2[:n])
		nums1 = append(nums1[:m], nums2[:n]... )
		
		// Sorting
		sorting := nums1
		sort.Ints(sorting)
		fmt.Println(sorting)
}

func main() {
	merge(nums1, m, nums2, n)
}