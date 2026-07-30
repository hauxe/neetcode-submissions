func checkValidString(s string) bool {
	left := []int{}
	star := []int{}
	for i := range s {
		switch s[i] {
		case '(':
			left = append(left, i)
		case ')':
			if len(left) > 0 {
				left = left[:len(left)-1]
			} else if len(star) > 0 {
				star = star[:len(star)-1]
			} else {
				return false
			}
		case '*':
			star = append(star, i)
		}
	}
	fmt.Println(left)
	fmt.Println(star)
	if len(star) < len(left) {
		return false
	}
	if len(left) == 0 {
		return true
	}
	for i := len(left) - 1; i >= 0; i-- {
		if star[len(star)-1] < left[i] {
			return false
		}
		star = star[:len(star)-1]
	}
	return true
}