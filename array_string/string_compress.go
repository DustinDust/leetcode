package arraystring

import "strconv"

func compress(chars []byte) int {
	var curI = 1
	var compressI = 0
	currentChar := chars[0]
	curCharCount := 1
	leng := 0

	for curI < len(chars) {
		currentChar = chars[curI]
		if chars[curI] == currentChar {
			curCharCount++
		} else {
			leng = leng + 1 + len(strconv.Itoa(curCharCount))
			currentChar = chars[curI]
			curCharCount = 1
		}
		curI++
	}
	return len(chars)
}
