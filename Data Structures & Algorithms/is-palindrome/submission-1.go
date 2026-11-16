func isPalindrome(s string) bool {
	begin, end := 0, len(s)-1
	var a, b byte
	var ok bool
	for begin < end {
		if a, ok = toLower(s[begin]); !ok {
			begin++
			continue
		}
		if b, ok = toLower(s[end]); !ok {
			end--
			continue
		}
		if a != b {
			return false
		}
		begin++
		end--
	}
	return true
}

func toLower(b byte) (byte, bool) {
	if '0' <= b && b <= '9' {
		return b, true
	}
	if 'a' <= b && b <= 'z' {
		return b, true
	}
	if 'A' <= b && b <= 'Z' {
		return b - 'A' + 'a', true
	}
	return 0, false
}