package solutions

import (
    "sort"
    "strconv"
)

func largestNumber(nums []int) string {
    var result string

    sort.Slice(nums, func(i, j int) bool {
        first, second := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])

        return first + second >= second + first
    })

    for _, number := range nums {
        result += strconv.Itoa(number)
    }

    if result[0] == '0' {
        return result[: 1]
    }

    return result
}