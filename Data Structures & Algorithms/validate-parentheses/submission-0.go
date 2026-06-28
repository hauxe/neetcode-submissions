func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := range s {
		if openChar, exist := m[s[i]]; exist {
			if len(stack) == 0 || stack[len(stack)-1] != openChar {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}