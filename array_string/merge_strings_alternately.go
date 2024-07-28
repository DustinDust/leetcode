package arraystring

import "math"

/*
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.
Return the merged string

Example 1:

Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
*/

func mergeAlternately(word1 string, word2 string) string {
	var i int
	diff := int(math.Abs(float64(len(word1) - len(word2))))
	var longer *string
	if len(word1) > len(word2) {
		longer = &word1
	} else {
		longer = &word2
	}

	result := make([]byte, len(word1)+len(word2))
	for k := 0; k < len(result)-diff; k += 2 {
		result[k] = word1[i]
		result[k+1] = word2[i]
		i++
	}
	for j := len(result) - diff; j < len(result); j++ {
		result[j] = (*longer)[i]
		i++
	}
	return string(result)
}
