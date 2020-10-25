package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}))
}

func findMaxConsecutiveOnes(nums []int) int {
	zIndices := make([]int, len(nums)+2)
	zIndex := 1

	// keep track of the indices for each zero, i.e. [0,1,0] makes zIndices=[0,2]
	for i, num := range nums {
		if num == 0 {
			zIndices[zIndex] = i
			zIndex++
		}
	}

	// add -1, and length of array before and after so [0,2] becomes [-1,0,2,3]
	zIndices[0] = -1
	zIndices[zIndex] = len(nums)

	// if no zeros were found, return the length of the array
	if zIndex == 1 {
		return len(nums)
	}

	// for each zero in zIndices add how many 1's are in between each zero
	maxLen := 0
	for i := 1; i < zIndex; i++ {
		length := zIndices[i+1] - zIndices[i-1] - 1
		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}
