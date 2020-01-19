package solutions

import (
    "sort"
)

func fourSum(nums []int, target int) [][]int {
    var result [][]int
    length := len(nums)

    if length < 4 {
        return result
    }

    sort.Ints(nums)

    for i := 0; i < length; i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        for j := i + 1; j < length; j++ {
            if j > i + 1 && nums[j] == nums[j - 1] {
                continue
            }

            k := j + 1
            l := length - 1

            for k < l {
                sum := nums[i] + nums[j] + nums[k] + nums[l]

                if sum == target {
                    result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})

                    k++
                    for ; k < length - 1 && nums[k] == nums[k - 1]; k++ {}
                } else if sum < target {
                    k++
                } else {
                    l--
                }
            }
        }
    }

    return result
}
