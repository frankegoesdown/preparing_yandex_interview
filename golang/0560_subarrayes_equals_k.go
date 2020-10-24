package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{3, 4, 7, 2, -3, 1, 4, 2}, 7))
}
func subarraySum(nums []int, k int) int {
	sums := make(map[int]int)
	res := 0
	sum := 0
	sums[0] = 1
	fmt.Println(sum)
	fmt.Println(res)
	fmt.Println(sums)
	fmt.Println("=====")
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res += sums[sum-k]
		sums[sum]++
		fmt.Println(sum)
		fmt.Println(res)
		fmt.Println(sums)
		fmt.Println("=====")
	}
	return res
}
