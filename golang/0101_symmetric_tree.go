package main

func main() {

}

func isSymmetric(root *TreeNode) bool {
	return (root == nil) || symmetric(root.Left, root.Right)
}

func symmetric(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val == r.Val {
		return symmetric(l.Left, r.Right) && symmetric(l.Right, r.Left)
	}
	return false
}
