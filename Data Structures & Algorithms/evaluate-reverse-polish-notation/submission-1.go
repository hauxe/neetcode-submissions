func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	p := 0
	for i := range tokens {
		switch tokens[i] {
		case "+":
			stack[p-2] += stack[p-1]
			p -= 1
		case "-":
			stack[p-2] -= stack[p-1]
			p -= 1
		case "*":
			stack[p-2] *= stack[p-1]
			p -= 1
		case "/":
			stack[p-2] /= stack[p-1]
			p -= 1
		default:
			stack[p] = parseInt(tokens[i])
			p++
		}

	}
	return stack[0]
}

func parseInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	val := 0
	for i := range s {
		if i == 0 && s[0] == '-' {
			sign = -1
			continue
		}
		val = 10*val + int(s[i]-'0')
	}
	return sign * val
}