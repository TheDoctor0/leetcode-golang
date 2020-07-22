package solutions

import (
    "strconv"
)

func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return nil
    }

    var result []string
    head := 0

    for i := range nums {
        if i < len(nums) - 1 && nums[i] + 1 == nums[i + 1] {
            continue
        }

        if head == i {
            result = append(result, strconv.Itoa(nums[i]))
        } else {
            current := strconv.Itoa(nums[head]) + "->" + strconv.Itoa(nums[i])
            result = append(result, current)
        }

        head = i + 1
    }

    return result
}
