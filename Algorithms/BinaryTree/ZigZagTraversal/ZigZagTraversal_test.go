package ZigZagTraversal

import (
	"testing"
)

func TestZigZagTraversalEmpty(t *testing.T) {
	var root *TreeNode
	want := []bool{}
	got := ZigZagTraversal(root)
	assertSliceEqual(t, want, got)
}

func TestZigZagTraversalSingle(t *testing.T) {
	root := &TreeNode{Value: true}
	want := []bool{true}
	got := ZigZagTraversal(root)
	assertSliceEqual(t, want, got)
}

func TestZigZagTraversalFirstCase(t *testing.T) {
	root := &TreeNode{Value: true}
	root.Left = &TreeNode{Value: true}
	root.Right = &TreeNode{Value: false}
	root.Left.Left = &TreeNode{Value: true}
	root.Left.Right = &TreeNode{Value: false}
	root.Right.Left = &TreeNode{Value: true}
	root.Right.Right = &TreeNode{Value: true}

	want := []bool{true, true, false, true, true, false, true}
	got := ZigZagTraversal(root)
	assertSliceEqual(t, want, got)
}

func TestZigZagTraversalSecondCase(t *testing.T) {
	root := &TreeNode{Value: true}
	root.Left = &TreeNode{Value: false}
	root.Right = &TreeNode{Value: false}
	root.Left.Left = &TreeNode{Value: true}
	root.Left.Right = &TreeNode{Value: true}
	root.Right.Left = &TreeNode{Value: false}
	root.Right.Right = &TreeNode{Value: false}

	want := []bool{true, false, false, false, false, true, true}
	got := ZigZagTraversal(root)
	assertSliceEqual(t, want, got)
}

func assertSliceEqual(t *testing.T, want, got []bool) {
	t.Helper()
	if len(want) != len(got) {
		t.Errorf("Expected slice length %d, but got %d", len(want), len(got))
	}
	for i := range want {
		if want[i] != got[i] {
			t.Errorf("Expected element %d to be %v, but got %v", i, want[i], got[i])
		}
	}
}
