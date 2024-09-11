package arraystring

import (
	"testing"
)

func stringsToChars(strs []string) []byte {
	bytes := make([]byte, len(strs))
	for i, s := range strs {
		bytes[i] = s[0]
	}
	return bytes
}

func TestCompress(t *testing.T) {
	chars := stringsToChars([]string{"a", "a", "b", "b", "c", "c", "c"})
	out := 6
	if got := compress(chars); got != out {
		t.Errorf("fail: expected %d got %d", out, got)
	}
}

func TestCompressShortString(t *testing.T) {
	chars := stringsToChars([]string{"a"})
	out := 1
	if got := compress(chars); got != out {
		t.Errorf("fail: expected %d got %d", out, got)
	}
}

func TestCompressBigCount(t *testing.T) {
	chars := stringsToChars([]string{"a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"})
	out := 14
	if got := compress(chars); got != out {
		t.Errorf("fail: expected %d got %d", out, got)
	}
}
