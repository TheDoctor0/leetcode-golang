package solutions

func maxArea(height []int) int {
	maxArea := 0

	i, j := 0, len(height)-1

	for {
		if i == j {
			break
		}

		minHeight := height[i]

		if height[j] < height[i] {
			minHeight = height[j]
		}

		area := (j - i) * minHeight

		if area > maxArea {
			maxArea = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}
