package solutions

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	first, second := nums1, nums2
	lengthFirst, lengthSecond := len(first), len(second)

	if lengthFirst > lengthSecond {
		first, second, lengthFirst, lengthSecond = second, first, lengthSecond, lengthFirst
	}

	min, max := 0, lengthFirst
	var i, j int
	var leftMax, rightMin float64

	for min <= max {
		i = (min + max) / 2
		j = (lengthFirst + lengthSecond + 1) / 2 - i

		switch {
		case i < lengthFirst && second[j - 1] > first[i]:
			min = i + 1
		case i > 0 && first[i - 1] > second[j]:
			max = i - 1
		default:
			if i == 0 {
				leftMax = float64(second[j - 1])
			} else if j == 0 {
				leftMax = float64(first[i - 1])
			} else {
				leftMax = math.Max(float64(first[i - 1]), float64(second[j - 1]))
			}

			if (lengthFirst + lengthSecond) % 2 == 1 {
				return leftMax
			}

			if i == lengthFirst {
				rightMin = float64(second[j])
			} else if j == lengthSecond {
				rightMin = float64(first[i])
			} else {
				rightMin = math.Min(float64(first[i]), float64(second[j]))
			}

			return (leftMax + rightMin) / 2
		}
	}

	return 0
}