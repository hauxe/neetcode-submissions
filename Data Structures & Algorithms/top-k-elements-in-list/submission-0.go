func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	buckets := make([][]int, len(nums))
	for k, v := range m {
		buckets[v-1] = append(buckets[v-1], k)
	}
	result := make([]int, 0, k)
	for i := len(nums) - 1; i >= 0; i-- {
		if len(buckets[i]) == 0 {
			continue
		}
		if len(buckets[i]) > k-len(result) {
			n := k - len(result)
			result = append(result, buckets[i][:n]...)
		} else {
			result = append(result, buckets[i]...)
		}
		if len(result) >= k {
			return result
		}
	}
	return nil
}