package main

// Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.
// Note that elements beyond the length of the original array are not written.
// Do the above modifications to the input array in place, do not return anything from your function.

// Example 1:
// Input: [1,0,2,3,0,4,5,0]
// Output: null
// Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

// [1,0,2,3,0,4,5,0]
// [1,0, ,2,3,0,4,5]
// [1,0,0,2,3,0,4,5]

// Example 2:
// Input: [1,2,3]
// Output: null
// Explanation: After calling your function, the input array is modified to: [1,2,3]

var arr = [8]int{1,0,2,3,0,4,5,0}

func duplicateZeros(arr [8]int) {

	// Make sure it loop only range of array prevent bug.
	for i := 0; i < len(arr); i++ {

		//Check if it found '0' in array
		if (arr[i] == 0) {

			// To insert array, we need to operate from end of array
			// Move array to Right side by copy Left to Right 
			for j := (len(arr)-1); j > i; j-- {
				// j = start from last index of array
				// i = where we found the index of '0'
				// loop until j greater than i (not equal)
				
				//Point to end of array = copy val from last array - 1 (left to right)
				arr[j] = arr[j-1]
				// at the last point will copy 0 from left to right
				// so we will have '0, 0' (double zero)
			}
			
			// since current i point to '0' already.
			// and we added '0' next to it, now we have double zero
			// we have to skip loop check by 1
			// to prevent from found '0' infinity
			i++
		}
	}
}


func main() {
	duplicateZeros(arr)
}