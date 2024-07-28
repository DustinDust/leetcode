package arraystring

func kidsWithCandies(candies []int, extraCandies int) []bool {
	curMax := 0
	result := make([]bool, len(candies))
	for _, v := range candies {
		if curMax < v {
			curMax = v
		}
	}
	for i, v := range candies {
		if v+extraCandies >= curMax {
			result[i] = true
		}
	}
	return result
}
