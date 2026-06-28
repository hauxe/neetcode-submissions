func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	begin, end := 0, len(nums1)
	m, n := len(nums1), len(nums2)
	for begin <= end {
		middle := (begin + end) / 2
		left1 := math.MinInt64
		if middle > 0 {
			left1 = nums1[middle-1]
		}
		right1 := math.MaxInt64
		if middle < m {
			right1 = nums1[middle]
		}
		nums1Count := middle
		nums2Count := (m+n)/2 - nums1Count
		if (m+n)%2 == 1 {
			nums2Count++
		}
		left2 := math.MinInt64
		if nums2Count > 0 {
			left2 = nums2[nums2Count-1]
		}
		right2 := math.MaxInt64
		if nums2Count < n {
			right2 = nums2[nums2Count]
		}
		// fmt.Println(left1, right1, left2, right2)
		if left1 <= right2 && left2 <= right1 {
			if (m+n)%2 == 0 {
				return float64(max(left1, left2)+min(right1, right2)) / 2.0
			}
			return float64(max(left1, left2))
		} else if left1 > right2 {
			end = middle - 1
		} else {
			begin = middle + 1
		}
		// fmt.Println(begin, end)

	}
	return 0.0
}