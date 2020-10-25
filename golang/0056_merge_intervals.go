package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i, interval := range intervals {
		if i == 0 {
			ans = append(ans, interval)
			continue
		}

		top := ans[len(ans)-1]
		if interval[0] > top[1] {
			ans = append(ans, interval)
		} else if interval[1] > top[1] {
			top[1] = interval[1]
		}
	}

	return ans
}
