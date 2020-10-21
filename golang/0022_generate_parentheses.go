package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	return helper(n, n, "", nil)
}

func helper(left, right int, queue string, res []string) []string {
	if left == 0 {
		return append(res, queue+strings.Repeat(")", right))
	}
	if right > left {
		res = helper(left, right-1, queue+")", res)
	}
	return helper(left-1, right, queue+"(", res)
}
