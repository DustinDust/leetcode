package arraystring

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := true
	output := increasingTriplet(input)
	if output != expected {
		t.Errorf("fail: expected %v, got %v", expected, output)
	}
}
