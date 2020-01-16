package solutions

import (
    "sort"
)

func threeSum(nums []int) [][]int {
    length := len(nums)
    var results [][]int

    if length < 3 {
        return results
    }

    sort.Ints(nums)

    for i := 0; i < length - 2; i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        target, left, right := -nums[i], i + 1, length - 1

        for left < right {
            sum := nums[left] + nums[right]

            if sum == target {
                results = append(results, []int{nums[i], nums[left], nums[right]})

                left++
                right--

                for left < right && nums[left] == nums[left - 1] {
                    left++
                }

                for left < right && nums[right] == nums[right + 1] {
                    right--
                }
            } else if sum > target {
                right--
            } else if sum < target {
                left++
            }
        }
    }

    return results
}
