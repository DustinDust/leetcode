package arraystring

/*
Given a string s, reverse only all the vowels in
the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u',
and they can appear in both lower and upper cases,
more than once.
*/

// SOLUTION: 2 pointers, on running from the beginning
// and one running from the end, towards each other
func reverseVowels(s string) string {
	isVowels := func(s byte) bool {
		switch s {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}
	result := []byte(s)
	p1, p2 := 0, len(s)-1
	for p1 < p2 {
		for p1 < p2 && !isVowels(result[p1]) {
			p1++
		}

		for p1 < p2 && !isVowels(result[p2]) {
			p2--
		}

		result[p1], result[p2] = result[p2], result[p1]
		p1++
		p2--
	}
	return string(result)
}
