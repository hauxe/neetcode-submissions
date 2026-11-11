func simplifyPath(path string) string {
	stack := []string{}
	new := true
	for i := range path {
		switch path[i] {
		case '/':
			if len(stack) > 0 {
				switch stack[len(stack)-1] {
				case ".":
					stack = stack[:len(stack)-1]
				case "..":
					if len(stack) > 1 {
						stack = stack[:len(stack)-2]
					} else {
						stack = stack[:len(stack)-1]
					}
				}
			}
			new = true

		default:
			if new {
				stack = append(stack, string(path[i]))
				new = false
			} else {
				stack[len(stack)-1] += string(path[i])
			}
		}
	}
	if len(stack) > 0 {
		switch stack[len(stack)-1] {
		case ".":
			stack = stack[:len(stack)-1]
		case "..":
			if len(stack) > 1 {
				stack = stack[:len(stack)-2]
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return "/" + strings.Join(stack, "/")
}