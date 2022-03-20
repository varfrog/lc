package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var nodes = flattenAndClearChildren(root)
	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	return buildNode(nodes)
}

func buildNode(nodes []*TreeNode) *TreeNode {
	if nodes == nil || len(nodes) == 0 {
		return nil
	}
	rootNode := nodes[len(nodes)/2]
	nodesLeft := nodes[:len(nodes)/2]
	nodesRight := nodes[len(nodes)/2+1:]

	return &TreeNode{
		Val:   rootNode.Val,
		Left:  buildNode(nodesLeft),
		Right: buildNode(nodesRight),
	}
}

func flattenAndClearChildren(node *TreeNode) []*TreeNode {
	if node == nil {
		return nil
	}
	var nodes []*TreeNode
	if node.Left != nil {
		nodes = append(nodes, flattenAndClearChildren(node.Left)...)
	}
	nodes = append(nodes, &TreeNode{Val: node.Val})
	if node.Right != nil {
		nodes = append(nodes, flattenAndClearChildren(node.Right)...)
	}
	node.Left = nil
	node.Right = nil

	return nodes
}
