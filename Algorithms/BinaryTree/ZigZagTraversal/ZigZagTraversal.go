package ZigZagTraversal

type TreeNode struct {
	Value bool
	Left   *TreeNode
	Right  *TreeNode
}

type stack struct {
	nodes []*TreeNode
}

func ZigZagTraversal(root *TreeNode) []bool {
	if root == nil {
		return []bool{}
	}

	s1 := new(stack)
	s2 := new(stack)
	s2.push(root)

	answer := []bool{}
	flag := false

	for !s1.isEmpty() || !s2.isEmpty() {
		if flag {
			for !s1.isEmpty() {
				node := s1.pop()
				answer = append(answer, node.Value)

				if node.Left != nil {
					s2.push(node.Left)
				}

				if node.Right != nil {
					s2.push(node.Right)
				}
			}
		} else {
			for !s2.isEmpty() {
				node := s2.pop()
				answer = append(answer, node.Value)
				if node.Right != nil {
					s1.push(node.Right)
				}
				if node.Left != nil {
					s1.push(node.Left)
				}
			}
		}
		flag = !flag
	}
	return answer
}

func reverse(slice []bool) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func (s *stack) push(node *TreeNode) {
	s.nodes = append(s.nodes, node)
}

func (s *stack) pop() *TreeNode {
	if s.isEmpty() {
		return nil
	}
	n := len(s.nodes) - 1
	node := s.nodes[n]
	s.nodes = s.nodes[:n]
	return node
}

func (s *stack) isEmpty() bool {
	return len(s.nodes) == 0
}
