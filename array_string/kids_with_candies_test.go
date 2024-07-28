package arraystring

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	input := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	expect := []bool{true, true, true, false, true}

	result := kidsWithCandies(input, extraCandies)
	if !reflect.DeepEqual(expect, result) {
		t.Errorf("expected: %v got %v", expect, result)
	}
}
