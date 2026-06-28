func isPalindrome(s string) bool {
	left, right := 0, len(s) - 1
	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}
		if !byteEqualFold(s[left], s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(b byte) bool {
	return ('a' <= b && b <= 'z') ||
	('A' <= b && b <= 'Z') ||
	('0' <= b && b <= '9')
}

func byteEqualFold(a byte, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a = a + 'a' - 'A'
	}
	if b >= 'A' && b <= 'Z' {
		b = b + 'a' - 'A'
	}
	return a == b
}
