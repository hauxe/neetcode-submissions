func calPoints(operations []string) int {
	result := []int{}
	for i := range operations {
		switch operations[i] {
		case "+":
			result = append(result, result[len(result)-1]+result[len(result)-2])
		case "C":
			result = result[:len(result)-1]
		case "D":
			result = append(result, result[len(result)-1]*2)
		default:
			v, _ := strconv.Atoi(operations[i])
			result = append(result, v)
		}
	}
	sum := 0
	for i := range result {
		sum += result[i]
	}
	return sum
}