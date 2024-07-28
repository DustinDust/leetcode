package arraystring

import (
	"strings"
)

func reverseWords(s string) string {
	result := make([]string, 0)
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != " " && words[i] != "" {
			result = append(result, words[i])
		}
	}
	return strings.Join(result, " ")
}
