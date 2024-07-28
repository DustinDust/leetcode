package arraystring

import (
	"testing"
)

func TestMerggStringsAlternatively(t *testing.T) {
	word1 := "abc"
	word2 := "pqr"
	expectedResullt := "apbqcr"
	result := mergeAlternately(word1, word2)
	if result != expectedResullt {
		t.Errorf("Failed: expected '%s', got '%s'", expectedResullt, result)
	}
}

func TestMergeStringAlternativelyDiffLength(t *testing.T) {
	word1 := "abcd"
	word2 := "pq"
	expectedResullt := "apbqcd"
	result := mergeAlternately(word1, word2)
	if result != expectedResullt {
		t.Errorf("Failed: expected '%s', got '%s'", expectedResullt, result)
	}
}
