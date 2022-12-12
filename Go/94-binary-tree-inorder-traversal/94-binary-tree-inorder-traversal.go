package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
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
