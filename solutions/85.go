package solutions

func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }

    area := 0
    width, height := len(matrix), len(matrix[0])
    left, right, maxHeight := make([]int, height), make([]int, height), make([]int, height)

    for i := range right {
        right[i] = height - 1
    }

    for i := 0; i < width; i++ {
        for j := 0; j < height; j++ {
            if matrix[i][j] == '1' {
                maxHeight[j] = maxHeight[j] + 1
            } else {
                maxHeight[j] = 0
            }
        }

        leftBottom := 0

        for j := 0; j < height; j++ {
            if matrix[i][j] == '1' {
                left[j] = max(left[j], leftBottom)
            } else {
                left[j] = 0
                leftBottom = j + 1
            }
        }

        rightBottom := height - 1

        for j := height - 1; j >= 0; j-- {
            if matrix[i][j] == '1' {
                right[j] = min(right[j], rightBottom)
            } else {
                right[j] = height - 1
                rightBottom = j - 1
            }
        }

        for j := 0; j < height; j++ {
            area = max(area, maxHeight[j] * (right[j] - left[j] + 1))
        }
    }

    return area
}
