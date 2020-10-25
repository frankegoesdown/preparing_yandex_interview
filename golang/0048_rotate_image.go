package main

import "fmt"

func main() {
	fmt.Println(rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func rotate(m [][]int) [][]int {
	n := len(m)
	for i := 0; i < n/2; i++ {
		fmt.Println(i)
		for j := i; j < n-i-1; j++ {
			fmt.Println(j)
			temp := m[i][j]
			fmt.Println(temp)
			fmt.Println(m[n-j-1][i])
			fmt.Println(m[n-i-1][n-j-1])
			fmt.Println(m[j][n-i-1])
			m[i][j] = m[n-j-1][i]
			m[n-j-1][i] = m[n-i-1][n-j-1]
			m[n-i-1][n-j-1] = m[j][n-i-1]
			m[j][n-i-1] = temp
			fmt.Println("--------------")
		}
		fmt.Println("==============")
	}
	return m
}
