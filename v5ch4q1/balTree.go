package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

	nums := []int{1, 2, 3, 4}

	fmt.Println(buildTreeFromSortedList(nums))
}

func buildTreeFromSortedList(nums []int) *TreeNode {
	sort.Ints(nums)
	var buildTree func(arr []int, start, end int) *TreeNode
	buildTree = func(arr []int, start, end int) *TreeNode {
		if start < 0 || end < 0 || start > end || end >= len(arr) {
			return nil
		}
		mid := start + (end-start)/2
		return &TreeNode{Val: arr[mid], Left: buildTree(arr, start, mid-1), Right: buildTree(arr, mid+1, end)}
	}
	return buildTree(nums, 0, len(nums)-1)
}
