package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return countGoodNodes(root, root.Val)
}

func countGoodNodes(root *TreeNode, greatestOnPath int) int {
	res := 0
	if root.Val >= greatestOnPath {
		res += 1
	}
	var newGreatest = greatestOnPath
	if root.Val > greatestOnPath {
		newGreatest = root.Val
	}
	if root.Left != nil {
		res += countGoodNodes(root.Left, newGreatest)
	}
	if root.Right != nil {
		res += countGoodNodes(root.Right, newGreatest)
	}
	return res
}
