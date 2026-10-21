func topKFrequent(nums []int, k int) []int {
	buckets := make([][]int, len(nums)+1)
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for num, count := range m {
		buckets[count] = append(buckets[count], num)
	}
	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		for j := range buckets[i] {
			result = append(result, buckets[i][j])
			if len(result) >= k {
				return result
			}
		}
	}
	return nil
}