package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	left, supermax, water := 0, 0, 0
	right := len(height) - 1
	for left < right {
		if height[left] < height[right] {
			supermax = max(height[left], supermax)
			water += supermax - height[left]
			left++
		} else {
			supermax = max(height[right], supermax)
			water += supermax - height[right]
			right--
		}
	}
	return water
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
