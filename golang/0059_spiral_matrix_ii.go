package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	top, bottom, left, right := 0, n-1, 0, n-1
	a, b := 1, n*n
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	for a <= b {
		for i := left; i <= right && a <= b; i++ {
			res[top][i] = a
			a++
		}
		top++
		for i := top; i <= bottom && a <= b; i++ {
			res[i][right] = a
			a++
		}
		right--
		for i := right; i >= left && a <= b; i-- {
			res[bottom][i] = a
			a++
		}
		bottom--
		for i := bottom; i >= top && a <= b; i-- {
			res[i][left] = a
			a++
		}
		left++
	}
	return res
}
