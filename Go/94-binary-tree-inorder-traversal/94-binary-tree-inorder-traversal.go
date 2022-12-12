package leetcode

import "fmt"

// func main() {
// 	fmt.Println("Hello")
// 	root := NewNode(1, nil, NewNode(2, NewNode(3, nil, nil), nil)) // 1, 3, 2
// 	fmt.Println(inorderTraversalIterative(root))
// }

type TreeNode struct {
	Val     int
	Left    *TreeNode
	Right   *TreeNode
	Visited bool
}

func (r *TreeNode) String() string {
	return fmt.Sprintf("%v %v", r.Val, r.Visited)
}

func NewNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func inorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}

	var pop func() *TreeNode = func() *TreeNode {
		rt := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return rt
	}
	var push func(*TreeNode) = func(r *TreeNode) {
		stack = append(stack, r)
	}

	push(root)

	for len(stack) != 0 {
		cur := pop()
		if cur == nil {
			continue
		}

		// move cur to most left of the tree
		for cur.Left != nil && !cur.Left.Visited {
			push(cur)
			cur = cur.Left
		}

		cur.Visited = true
		result = append(result, cur.Val)

		if cur.Right != nil {
			push(cur.Right)
		}

	}
	return result

}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	var inorder func(r *TreeNode)
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		// left self and right
		if r.Left != nil {
			inorder(r.Left)
		}

		result = append(result, r.Val)

		if r.Right != nil {
			inorder(r.Right)
		}
	}
	inorder(root)

	return result
}
