func generateParenthesis(n int) []string {
	return doGenerateParenthesis(n, 0)
}

func doGenerateParenthesis(left int, right int) []string {
	var result []string
	if left == 0 && right == 0 {
		return []string{""}
	}
	if left > 0 {
		list := doGenerateParenthesis(left-1, right+1)
		for i := range list {
			result = append(result, "("+list[i])
		}
	}
	if right > 0 {
		list := doGenerateParenthesis(left, right-1)
		for i := range list {
			result = append(result, ")"+list[i])
		}
	}
	return result
}