package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasChildren(node *TreeNode) bool {
	return node.Left != nil || node.Right != nil
}

func calcHeight(node *TreeNode, heightCache *map[*TreeNode]int8) int8 {
	if node == nil {
		(*heightCache)[node] = 0
		return 0
	}
	cachedHeight, ok := (*heightCache)[node]
	if ok {
		return cachedHeight
	}

	if !hasChildren(node) {
		(*heightCache)[node] = 1
		return 1
	}

	var leftHeight int8
	var rightHeight int8

	if node.Left != nil {
		leftHeight += 1 + calcHeight(node.Left, heightCache)
	}
	if node.Right != nil {
		rightHeight += 1 + calcHeight(node.Right, heightCache)
	}

	var result int8
	if leftHeight > rightHeight {
		result = leftHeight
	} else {
		result = rightHeight
	}
	(*heightCache)[node] = result
	return result
}

func isBalanced(root *TreeNode) bool {
	return isBalancedCached(root, &map[*TreeNode]int8{})
}

func isBalancedCached(root *TreeNode, heightCache *map[*TreeNode]int8) bool {
	if root == nil {
		return true
	}
	var leftNode = root.Left
	var rightNode = root.Right

	leftHeight := calcHeight(leftNode, heightCache)
	rightHeight := calcHeight(rightNode, heightCache)
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return false
	}

	leftNodeBalanced := true
	if leftNode != nil {
		leftNodeBalanced = isBalancedCached(leftNode, heightCache)
	}
	rightNodeBalanced := true
	if rightNode != nil {
		rightNodeBalanced = isBalancedCached(rightNode, heightCache)
	}

	return leftNodeBalanced && rightNodeBalanced

}
