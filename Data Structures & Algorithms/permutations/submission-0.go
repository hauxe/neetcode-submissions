
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var result [][]int
	for i := range nums {
		newNums := make([]int, 0, len(nums)-1)
		newNums = append(newNums, nums[0:i]...)
		newNums = append(newNums, nums[i+1:]...)
		
		perms := permute(newNums)
		for j := range perms {
			result = append(result, append([]int{nums[i]}, perms[j]...))
		}
	}
	return result
}