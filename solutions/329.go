package solutions

import (
    "math"
)

func longestIncreasingPath(matrix [][]int) int {
    max := 0

    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return max
    }

    visited := make([][]int, len(matrix))

    for i := range visited{
        visited[i] = make([]int , len(matrix[0]))
    }

    for i, row := range matrix {
        for j := range row {
            if visited[i][j] > 0 {
                continue
            }

            previous := math.MinInt32
            path := getMaxPath(matrix, visited, i, j, previous)

            if path > max {
                max = path
            }
        }
    }
    return max
}

func getMaxPath(matrix, visited [][]int, i, j, prevVal int) int {
    if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || matrix[i][j] <= prevVal {
        return 0
    }

    if visited[i][j] > 0 {
        return visited[i][j]
    }

    var directions = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    max := 0

    for _, dir := range directions {
        x, y := i + dir[0], j + dir[1]
        path := getMaxPath(matrix, visited, x, y, matrix[i][j])

        if path > max {
            max = path
        }
    }

    visited[i][j] = max + 1

    return visited[i][j]
}
