func decodeString(s string) string {
	stack1 := []byte{}
	stack2 := []int{}
	stack3 := []int{}
	num := 0
	for i := range s {
		c := s[i]
		switch {
		case '0' <= c && c <= '9':
			num = num*10 + int(c-'0')
		case c == '[':
			stack2 = append(stack2, len(stack1))
			stack3 = append(stack3, num)
			num = 0
		case c == ']':
			idx, num := stack2[len(stack2)-1], stack3[len(stack3)-1]
			stack2, stack3 = stack2[:len(stack2)-1], stack3[:len(stack3)-1]
			loop := stack1[idx:]
			for range num - 1 {
				stack1 = append(stack1, loop...)
			}
		default:
			stack1 = append(stack1, c)
		}
	}
	return string(stack1)
}