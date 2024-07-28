package arraystring

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	input := "the sky is blue"
	expected := "blue is sky the"

	got := reverseWords(input)
	if expected != got {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}
