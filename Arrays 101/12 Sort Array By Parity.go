// Given an array A of non-negative integers, return an array consisting of all the even elements of A,
// followed by all the odd elements of A.
// You may return any answer array that satisfies this condition.

// Example 1:
// Input: [3,1,2,4]
// Output: [2,4,3,1]
// The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

// Note:
// 1 <= A.length <= 5000
// 0 <= A[i] <= 5000

package main

import "fmt"

// var A = []int{3,1,0,2,4,1}
// var A = []int{3,1,2,0,3,4}
var A = []int{3,1,0,2,4,1}
// var A = []int{3,1,2,4}
// ==========[2,4,3,1]

func sortArrayByParity(A []int) []int {
	if len(A) <= 1 { return A }

	// Sort Array First
	for i := 0; i < len(A)-1; i++ {
		for j := i+1; j < len(A); j++ {
			if A[i] > A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}

	for i := 0; i < len(A); i++ {
		for j := i+1; j < len(A); j++ {
			if (A[i] % 2) != 0 {

				if (A[j] % 2) == 0 {
					A[i], A[j] = A[j], A[i]	
				}

			}
			// fmt.Println(j)
		}
	}

	return A
}

func main() {
	fmt.Println(sortArrayByParity(A))
}