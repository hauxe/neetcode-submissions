func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	if n == 0 {
		return 1
	}
	half := myPow(x, n/2)
	if n&1 == 1 {
		return half * half * x
	}
	return half * half
}