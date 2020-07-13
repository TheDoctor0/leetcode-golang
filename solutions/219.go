package solutions

func containsNearbyDuplicate(nums []int, k int) bool {
    index := make(map[int]int, k)

    for i, n := range nums {
        if i > k {
            obsoleteIndex := i - k - 1

            if index[nums[obsoleteIndex]] == obsoleteIndex {
                delete(index, nums[obsoleteIndex])
            }
        }

        if previous, ok := index[n]; ok {
            if i - previous <= k {
                return true
            }
        }

        index[n] = i
    }

    return false
}
