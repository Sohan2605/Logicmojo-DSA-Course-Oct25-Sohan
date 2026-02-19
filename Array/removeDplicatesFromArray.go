func removeDuplicates(nums []int) int {
    	if len(nums) == 0 {
		return 0
	}

	// slow pointer
	i := 0

	// fast pointer
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}