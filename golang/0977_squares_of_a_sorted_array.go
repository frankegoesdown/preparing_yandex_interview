package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(A []int) (b []int) {
	size := len(A)
	res := make([]int, size)
	for l, r, i := 0, size-1, size-1; l <= r; i-- {
		if A[l]+A[r] < 0 {
			res[i] = A[l] * A[l]
			l++
		} else {
			res[i] = A[r] * A[r]
			r--
		}
	}
	return res
}
