package arraystring

import "testing"

func TestGcdOfStrings(t *testing.T) {
	str1 := "ABABAB"
	str2 := "ABAB"
	expected := "AB"

	result := gcdOfStrings(str1, str2)
	if expected != result {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
