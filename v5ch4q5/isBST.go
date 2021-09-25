package main

import "fmt"

type TreeNode struct {
	val                 int
	left_ptr, right_ptr *TreeNode
}

const MAX_INT = int(^(uint(0)) >> 1)
const MIN_INT = -MAX_INT - 1

func main() {
	var root *TreeNode
	fmt.Println(isBST(root))
}

func isBST(root *TreeNode) bool {
	return isBSTHelper(root, MIN_INT, MAX_INT)
}

func isBSTHelper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if min > root.val || max < root.val {
		return false
	}
	return isBSTHelper(root.left_ptr, min, root.val) && isBSTHelper(root.right_ptr, root.val, max)
}
