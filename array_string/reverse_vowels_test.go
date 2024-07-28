package arraystring

import "testing"

func TestReverseVowels(t *testing.T) {
	input := "leetcode"
	expected := "leotcede"

	got := reverseVowels(input)
	if expected != got {
		t.Errorf("expected: %s, got %s", expected, got)
	}
}

func TestReverseVowelsShortString(t *testing.T) {
	input := "ai"
	expected := "ia"

	got := reverseVowels(input)
	if expected != got {
		t.Errorf("expected: %s, got %s", expected, got)
	}
}

func TestReverseVowelsWithUppercase(t *testing.T) {
	input := "aA"
	expected := "Aa"

	got := reverseVowels(input)
	if expected != got {
		t.Errorf("expected: %s, got %s", expected, got)
	}
}
