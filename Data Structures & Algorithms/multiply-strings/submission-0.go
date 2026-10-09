func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	res := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			sum := mul + res[i+j+1] // add what's already in the ones slot
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10 // carry into the next higher slot
		}
	}

	// skip leading zeros
	start := 0
	for start < len(res)-1 && res[start] == 0 {
		start++
	}

	b := make([]byte, 0, len(res)-start)
	for ; start < len(res); start++ {
		b = append(b, byte(res[start])+'0')
	}
	return string(b)
}