func minCostClimbingStairs(cost []int) int {
	costMap := make(map[int]int)
	var findCost func(cost []int, start int) int
	findCost = func(cost []int, start int) int {
		if start >= len(cost) {
			return 0
		}
		if start == len(cost)-1 {
			return cost[start]
		}
		if val, exist := costMap[start]; exist {
			return val
		}
		cost1 := findCost(cost, start+2) + cost[start]
		cost2 := findCost(cost, start+1) + cost[start]

		costMap[start] = min(cost1, cost2)

		return costMap[start]
	}
	return min(findCost(cost, 1), findCost(cost, 0))
}