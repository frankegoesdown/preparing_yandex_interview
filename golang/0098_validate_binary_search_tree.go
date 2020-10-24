package main

import "math"

func main() {

}

func isValidBST(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return valid(root.Left, math.MinInt64, root.Val) && valid(root.Right, root.Val, math.MaxInt64)

}

func valid(root *TreeNode, min, max int) bool {
	if nil == root {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return valid(root.Left, min, root.Val) && valid(root.Right, root.Val, max)
}
