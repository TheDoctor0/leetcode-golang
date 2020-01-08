package solutions

func removeDuplicates(nums []int) int {
	length := len(nums)

	if length <= 1 {
		return length
	}

	unique, previous := 1, nums[0]

	for i := 1; i < length; i++ {
		if nums[i] != previous {
			nums[unique] = nums[i]
			previous = nums[i]

			unique++
		}
	}

	return unique
}
