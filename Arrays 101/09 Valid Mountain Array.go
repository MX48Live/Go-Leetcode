package main

import "fmt"

// false
// var arr = []int{2,1}
// var arr = []int{0,3,2,1}
// var arr = []int{3,5,5}
// var arr = []int{3,2,1}
// var arr = []int{4,5,4,3,2,1}
// var arr = []int{4,5,6,4,3,2}
// var arr = []int{3,7,6,4,3,0,1,0}
// var arr = []int{3,7,6,4,3,0,1,0}

// true
var arr = []int{0,3,2,1}

func validMountainArray(arr []int) bool {
  
	// Array should be greather than 3 and 2 must > than 1
	if (len(arr) < 3) || (arr[0] > arr[1]) { return false }

	stillGoUp := true

	for i := 1; i < len(arr); i++ {
		// if it equal should be done
		if arr[i - 1] == arr[i] { return false }

		if (stillGoUp) {

			if arr[i-1] > arr[i] { stillGoUp = false } 

		} else {

			if arr[i-1] < arr[i] { return false }
		}
	}
	return !stillGoUp
}	

func main() {
	fmt.Println(validMountainArray(arr))
}