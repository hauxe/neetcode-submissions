func merge(nums1 []int, m int, nums2 []int, n int) {
	// move 0s to begin

	for p1, p2 := m-1, m+n-1; p1 >= 0; p1, p2 = p1-1, p2-1 {
		nums1[p1], nums1[p2] = nums1[p2], nums1[p1]
	}

	// merge
	p1, p2, p3 := 0, n, 0
	for p3 < n {
		switch {
		case p2 < m+n && nums1[p2] < nums2[p3]:
			nums1[p1], nums1[p2] = nums1[p2], nums1[p1]
			p1++
			p2++
		default:
			nums1[p1] = nums2[p3]
			p1++
			p3++
		}
	}
}