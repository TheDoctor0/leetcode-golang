package solutions

func largestRectangleArea(heights []int) int {
    area := 0

    for i := 1; i < len(heights); i++ {
        if heights[i] >= heights[i-1] {
            continue
        }

        for j := i - 1; j >= 0; j-- {
            if heights[j] <= heights[i] {
                break
            }

            currentArea := heights[j] * (i - j)

            if currentArea > area {
                area = currentArea
            }

            heights[j] = heights[i]
        }
    }

    for i, height := range heights {
        currentArea := (len(heights) - i) * height

        if currentArea > area {
            area = currentArea
        }
    }

    return area
}