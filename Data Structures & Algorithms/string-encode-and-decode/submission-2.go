type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result []byte
	for _, str := range strs {
		for n := len(str); n > 0; n /= 10 {
			result = append(result, byte(n%10)+'0')
		}
		result = append(result, '#')
		result = append(result, []byte(str)...)
	}
	return string(result)
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	idx := 0
	n := len(encoded)
	for idx < n {
		length := 0
		inc := 1
		for ; idx < n; idx++ {
			if encoded[idx] == '#' {
				break
			}
			length += int(encoded[idx]-'0') * inc
			inc *= 10
		}
		idx++
		result = append(result, string(encoded[idx:idx+length]))
		idx += length
	}
	return result
}