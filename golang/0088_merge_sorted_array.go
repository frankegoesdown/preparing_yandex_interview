package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i < m+n && j < n {
		if nums1[i] > nums2[j] {
			copy(nums1[i+1:], nums1[i:])
			nums1[i] = nums2[j]
			j++
			continue
		}
		i++
	}
	if j < n {
		copy(nums1[m+j:], nums2[j:])
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) []int {
	if n == 0 {
		return nums1
	}
	i := 0
	j := 0
	for i != m {
		if nums1[i] > nums2[j] {
			tmp := nums2[j]
			nums2[j] = nums1[i]
			nums1[i] = tmp
			sort.Ints(nums2)
		}
		i++
	}
	for i := 0; i < n; i++ {
		nums1[m] = nums2[i]
		m++
	}
	return nums1
}
