package stack

func removeStars(s string) string {
	res := []byte{}
	chars := []byte(s)

	for _, r := range chars {
		if r == '*' {
			res = res[0 : len(res)-1]
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}
