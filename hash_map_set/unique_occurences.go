package hashmapset

func uniqueOccurrences(arr []int) bool {
	set := make(map[int]int, 0)
	occurences := make(map[int]interface{}, 0)

	for _, v := range arr {
		if _, ok := set[v]; ok {
			set[v] += 1
		} else {
			set[v] = 0
		}
	}

	for _, v := range set {
		if _, ok := occurences[v]; ok {
			return false
		} else {
			occurences[v] = struct{}{}
		}
	}
	return true
}
