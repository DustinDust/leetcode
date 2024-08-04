package binarytree

import "testing"

func TestGoodNodes(t *testing.T) {
	input := arrayToTree([](*int){ptr(3), ptr(1), ptr(4), ptr(3), nil, ptr(1), ptr(5)})
	got := goodNodes(input)
	expected := 4

	if got != expected {
		t.Errorf("Fail: expects %v but got %v", expected, got)
	}
}

func arrayToTree(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}
	queue := []*TreeNode{root}
	index := 1

	for index < len(arr) {
		current := queue[0]
		queue = queue[1:]

		// Check if left child exists
		if index < len(arr) && arr[index] != nil {
			current.Left = &TreeNode{Val: *arr[index]}
			queue = append(queue, current.Left)
		}
		index++

		// Check if right child exists
		if index < len(arr) && arr[index] != nil {
			current.Right = &TreeNode{Val: *arr[index]}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}

// treeToArray converts a binary tree to an array
// The array is represented in level-order traversal and includes `nil` for missing nodes
func treeToArray(root *TreeNode) []*int {
	if root == nil {
		return []*int{}
	}

	var result []*int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current != nil {
			val := current.Val
			result = append(result, &val)
			queue = append(queue, current.Left, current.Right)
		} else {
			result = append(result, nil)
		}
	}

	// Trim trailing nils which are not needed
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func ptr(x int) *int {
	return &x
}
