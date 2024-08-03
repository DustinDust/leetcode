package datastructures

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, expected: []int{5, 4, 3, 2, 1}},
		{input: []int{1, 2}, expected: []int{2, 1}},
		{input: []int{}, expected: []int{}},
		{input: []int{1}, expected: []int{1}},
	}

	for _, tc := range testCases {
		list := sliceToList(tc.input)
		reversedList := reverseList(list)
		result := listToSlice(reversedList)
		if !reflect.DeepEqual(result, tc.expected) {
            if len(result) == 0 && len(tc.expected) == 0 {
                continue
            }
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

// Helper function to convert a slice to a linked list.
func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Helper function to convert a linked list to a slice.
func listToSlice(head *ListNode) []int {
	var nums []int
	current := head
	for current != nil {
		nums = append(nums, current.Val)
		current = current.Next
	}
	return nums
}
