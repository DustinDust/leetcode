package arraystring

import (
	"reflect"
	"testing"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	output := productExceptSelf(input)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Fail: expected %v, got %v", expected, output)
	}
}

func TestProductOfArrayExceptSelfWithZeroes(t *testing.T) {
	input := []int{-1, 1, 0, -3, 3}
	expected := []int{0, 0, 9, 0, 0}

	output := productExceptSelf(input)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Fail: expected %v, got %v", expected, output)
	}
}
