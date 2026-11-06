func evalRPN(tokens []string) int {
	stack := []int{}
	p := -1
	for i := range tokens {
		switch tokens[i] {
		case "+":
			p--
			stack[p] += stack[p+1]
		case "-":
			p--
			stack[p] -= stack[p+1]
		case "*":
			p--
			stack[p] *= stack[p+1]
		case "/":
			p--
			stack[p] /= stack[p+1]
		default:
			p++
			val := parseInt(tokens[i])
			if p >= len(stack) {
				stack = append(stack, val)
			} else {
				stack[p] = val
			}
		}
	}
	return stack[0]
}

func parseInt(s string) int {
	sign := 1
	val := 0
	for i := range s {
		switch s[i] {
		case '-':
			sign = -1
		default:
			val = val*10 + int(s[i]-'0')
		}
	}
	return val * sign
}