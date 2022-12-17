package main

func main() {

	pathSum(&TreeNode{}, 0)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int

	if root == nil {
		return ans
	}

	// if this is leaf
	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			ans = append(ans, []int{root.Val})
		}
		return ans
	}

	value := targetSum - root.Val

	// if has left node
	if root.Left != nil {
		for _, v := range pathSum(root.Left, value) {
			// push root.Val to this path
			path := []int{root.Val}
			path = append(path, v...)
			ans = append(ans, path)
		}

	}

	if root.Right != nil {
		for _, v := range pathSum(root.Right, value) {
			// push root.Val to this path
			path := []int{root.Val}
			path = append(path, v...)
			ans = append(ans, path)
		}
	}

	return ans

}
