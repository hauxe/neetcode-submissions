func sortColors(nums []int) {
	red, white, blue := 0, 0, len(nums)-1 // 0-red, 1-white, 2-blue
	for white <= blue {
		switch nums[white] {
		case 0:
			nums[red], nums[white] = nums[white], nums[red]
			red++
			white++
		case 1:
			white++
		case 2:
			nums[white], nums[blue] = nums[blue], nums[white]
			blue--
		}
	}
}