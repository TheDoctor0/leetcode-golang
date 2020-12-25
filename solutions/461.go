package solutions

func hammingDistance(x int, y int) int {
	distance, result := x^y, 0

	for distance != 0 {
		if distance & 1 == 1 {
			result++
		}

		distance >>= 1
	}

	return result
}
