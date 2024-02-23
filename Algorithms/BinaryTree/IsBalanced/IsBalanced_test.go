package main

import (
	"testing"
)

func TestIsTreeBalancedEmpty(t *testing.T) {
	var root *TreeNode

	want := true
	got := isTreeBalanced(root)
	assertEqual(t, want, got)
}

func TestIsTreeBalancedSingle(t *testing.T) {
	root := &TreeNode{Value: true}

	want := true
	got := isTreeBalanced(root)
	assertEqual(t, want, got)
}

func TestIsTreeBalancedBalanced(t *testing.T) {
	root := &TreeNode{Value: true}
	root.Left = &TreeNode{Value: true}
	root.Right = &TreeNode{Value: false}
	root.Left.Left = &TreeNode{Value: true}
	root.Left.Right = &TreeNode{Value: false}
	root.Right.Left = &TreeNode{Value: true}
	root.Right.Right = &TreeNode{Value: true}

	want := true
	got := isTreeBalanced(root)
	assertEqual(t, want, got)
}

func TestIsTreeBalancedUnbalanced(t *testing.T) {
	root := &TreeNode{Value: true}
	root.Left = &TreeNode{Value: true}
	root.Right = &TreeNode{Value: true}
	root.Left.Left = &TreeNode{Value: true}

	want := false
	got := isTreeBalanced(root)
	assertEqual(t, want, got)
}

func assertEqual(t *testing.T, want, got bool) {
	t.Helper()
	if want != got {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
