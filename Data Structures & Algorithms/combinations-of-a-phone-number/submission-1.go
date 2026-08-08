var m = [8][]string{
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return m[int(digits[0]-'0')-2]
	}
	c := digits[0]
	letters := letterCombinations(digits[1:])
	var result []string
	for i := range letters {
		for _, s := range m[int(c-'0')-2] {
			result = append(result, s+letters[i])
		}
	}
	return result
}