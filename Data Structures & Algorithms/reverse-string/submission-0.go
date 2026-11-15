func reverseString(s []byte) {
	for begin, end := 0, len(s)-1; begin < end; begin, end = begin+1, end-1 {
		s[begin], s[end] = s[end], s[begin]
	}
}