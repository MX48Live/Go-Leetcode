package main

import (
	"fmt"
	"sort"
)

// var heights = []int{1,1,4,2,1,3}
//3
// var heights = []int{5,1,2,3,4}
//5
var heights = []int{1,2,3,4,5} 
//0

func heightChecker(heights []int) int {
	if len(heights) < 2 { return 0 }
	new := make([]int, len(heights))
	copy(new, heights)
	sort.Ints(new)
	count := 0


	for i, _ := range heights {
		fmt.Printf("new: %v | heights: %v\n", new[i], heights[i] )
		if new[i] != heights[i] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(heightChecker(heights))
}