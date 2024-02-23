package main

type TreeNode struct {
	Value bool
	Left   *TreeNode
	Right  *TreeNode
}

func isTreeBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftCount := countTrueValues(root.Left)
	rightCount := countTrueValues(root.Right)
	return leftCount == rightCount
}

func countTrueValues(node *TreeNode) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.Value {
		count = 1
	}
	return count + countTrueValues(node.Left) + countTrueValues(node.Right)
}
