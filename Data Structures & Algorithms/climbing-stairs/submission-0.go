func climbStairs(n int) int {
	stepCount := make(map[int]int)
	var doClimb func(n int) int
	doClimb = func(n int) int {
		if n == 0 {
			return 1
		}
		if n == 1 {
			return 1
		}
		if count, exist := stepCount[n]; exist {
			return count
		}
		stepN2 := doClimb(n - 2)
		stepCount[n-2] = stepN2
		stepN1 := doClimb(n - 1)
		stepCount[n-1] = stepN1
		stepCount[n] = stepN1 + stepN2
		return stepN1 + stepN2
	}
	return doClimb(n)
}