package solutions

func circularArrayLoop(nums []int) bool {
    for i, value := range nums {
        nums[i] = value % len(nums)
    }

    for i := range nums {
        if nums[i] % len(nums) <= 0 {
            continue
        }

        index := i

        for nums[index] % len(nums) > 0 {
            nextIndex := (index + nums[index]) % len(nums)
            nums[index] = (i + 1) * len(nums)
            index = nextIndex
        }

        if nums[index] == (i + 1) * len(nums) {
            return true
        }
    }

    for i := range nums {
        if nums[i] % len(nums) >= 0 {
            continue
        }

        index := i
        feet := map[int]interface{}{}

        for nums[index] % len(nums) < 0 {
            feet[index] = nil
            nextIndex := (index + nums[index] + len(nums)) % len(nums)
            nums[index] = -1 * (i + 1) * len(nums)
            index = nextIndex
        }

        if nums[index] == -1 * (i + 1) * len(nums) {
            return true
        }
    }

    return false
}
