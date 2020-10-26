package main

import (
	"fmt"
	"strings"
)

// time O(4n/sqrt(n))
// space O(4n/sqrt(n)) + O(n) store sequence
func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	return helper(n, n, "", nil)
}

func helper(left, right int, queue string, res []string) []string {
	fmt.Println(left)
	fmt.Println(right)
	fmt.Println(queue)
	fmt.Println(res)
	if left == 0 {
		return append(res, queue+strings.Repeat(")", right))
	}
	if right > left {
		res = helper(left, right-1, queue+")", res)
	}
	return helper(left-1, right, queue+"(", res)
}

func generateParenthesis2(n int) []string {
	res := []string{}
	helper2(n, 0, 0, "", &res)
	return res
}
func helper2(n, left, right int, temp string, res *[]string) {
	if n == right {
		*res = append(*res, temp)
		return
	}

	if left != n {
		helper2(n, left+1, right, temp+"(", res)
	}

	if left > right {
		helper2(n, left, right+1, temp+")", res)
	}
}
