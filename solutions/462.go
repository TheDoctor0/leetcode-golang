package solutions

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var moves = 0
	middle := nums[len(nums) / 2]

	for _, value := range nums {
		moves += int(math.Abs(float64(middle - value)))
	}

	return moves
}
