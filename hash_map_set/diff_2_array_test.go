package hashmapset

import (
	"reflect"
	"testing"
)

func TestFindDifference(t *testing.T) {
	input1 := []int{1, 2, 3}
	input2 := []int{2, 4, 6}

	expected := [][]int{
		{1, 3},
		{4, 6},
	}
	got := findDifference(input1, input2)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected: %v, got %v", expected, got)
	}
}
