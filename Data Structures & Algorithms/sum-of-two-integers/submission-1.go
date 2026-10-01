func getSum(a int, b int) int {
	for b != 0 {
		remain := (a & b) << 1
		a ^= b
		b = remain
	}
	return a
}