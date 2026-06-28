type Solution struct{}

var length = 3

func (s *Solution) Encode(strs []string) string {
	sum := 0
	for i := range strs {
		sum += (length + len(strs[i]))
	}
	encoded := make([]byte, sum)
	idx := 0
	for i := range strs {
		n := len(strs[i])
		for j := 0; j < length; j++ {
			encoded[idx+length-j-1] = byte(n%10) + '0'
			n = n / 10
		}
		idx += length
		copy(encoded[idx:idx+len(strs[i])], []byte(strs[i]))
		idx += len(strs[i])
	}
	return string(encoded)
}

func (s *Solution) Decode(encoded string) []string {
	idx := 0
	result := make([]string, 0)
	for idx < len(encoded) {
		l := 0
		for j := 0; j < length; j++ {
			n := int(encoded[idx+j] - '0')
			l = (10*l + n)
		}
		idx += length
		result = append(result, string(encoded[idx:idx+l]))
		idx += l
	}
	return result
}