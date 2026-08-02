func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	return doCombinationSum2(candidates, target)
}

func doCombinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	validIdx := len(candidates)
	for i := range candidates {
		if candidates[i] > target {
			validIdx = i
			break
		}
	}
	candidates = candidates[:validIdx]
	var result [][]int
	last := len(candidates)
	for i := last - 1; i >= 0; i-- {
		if last == len(candidates) || candidates[i] < candidates[last] {
			last = i
		}
		val := candidates[i] * (last - i + 1)
		if val == target {
			vals := []int{}
			for k := 0; k < last-i+1; k++ {
				vals = append(vals, candidates[i])
			}
			result = append(result, vals)
		} else if val < target {
			
			sums := doCombinationSum2(candidates[last+1:], target-val)
			if len(sums) > 0 {
				for j := range sums {
					for k := 0; k < last-i+1; k++ {
						sums[j] = append(sums[j], candidates[i])
					}
				}
			}
			result = append(result, sums...)
		}
	}
	return result
}