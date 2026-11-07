func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for i := range asteroids {
		if asteroids[i] < 0 {
			j := len(stack) - 1
			pop := false
			// possible of collision
			for ; j >= 0; j-- {
				if stack[j] < 0 {
					break
				}
				if stack[j] == -1*asteroids[i] {
					pop = true
					stack = stack[:j]
					break
				}
				if stack[j] > -1*asteroids[i] {
					pop = true
					break
				}
				stack = stack[:j]
			}
			if !pop {
				stack = append(stack, asteroids[i])
			}
			continue
		}
		stack = append(stack, asteroids[i])
	}
	return stack
}