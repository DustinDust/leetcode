package stack

import "testing"

func TestRemoveStart(t *testing.T) {
	input := "leet**cod*e"
	expected := "lecoe"

	got := removeStars(input)

	if got != expected {
		t.Errorf("Fail: expected %s got %s", expected, got)
	}
}
