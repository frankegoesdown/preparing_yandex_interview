package main

func main() {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lCombined := append(nums1, nums2...)
	sort.Ints(lCombined)
	n := len(lCombined)
	if n%2 == 0 {
		return (float64(lCombined[n/2] + lCombined[(n/2)-1])) / 2
	} else {
		return float64(lCombined[(n-1)/2])
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)

	i := 0
	j := 0

	which := func() int {
		tmp := 0
		if i < l1 && j < l2 {
			if nums1[i] > nums2[j] {
				tmp = nums2[j]
				j++
			} else {
				tmp = nums1[i]
				i++
			}
		} else {
			if j < l2 {
				tmp = nums2[j]
				j++
			}

			if i < l1 {
				tmp = nums1[i]
				i++
			}
		}

		return tmp
	}

	t := (l1 + l2) / 2
	pre := 0
	for p := 0; p < t; p++ {
		pre = which()
	}

	tmp := which()

	if (l1+l2)%2 == 0 {
		return float64(tmp+pre) / 2
	}

	return float64(tmp)
}
