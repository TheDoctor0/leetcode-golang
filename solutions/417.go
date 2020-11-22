package solutions

func pacificAtlantic(matrix [][]int) [][]int {
    if matrix == nil || len(matrix) < 1 {
        return nil
    }

    var result [][]int
    m, n := len(matrix), len(matrix[0])
    pacific, atlantic := make([][] bool, m), make([][] bool, m)

    for i := 0; i < m; i++ {
        pacific[i] = make([]bool, n)
        atlantic[i] = make([]bool, n)
    }

    for i := 0; i < m; i++ {
        dfs(i, 0, matrix, pacific)
    }

    for j := 0; j < n; j++ {
        dfs(0, j, matrix, pacific)
    }

    for i := 0; i < m; i++ {
        dfs(i, n - 1, matrix, atlantic)
    }

    for j := 0; j < n; j++ {
        dfs(m - 1, j, matrix, atlantic)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if pacific[i][j] && atlantic[i][j] {
                result = append(result, []int{i, j})
            }
        }
    }

    return result

}

func dfs(i int, j int, matrix [][]int, visited [][]bool) {
    direction := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
    m, n := len(matrix), len(matrix[0])
    visited[i][j] = true

    for k := 0; k < 4; k++ {
        x, y := i + direction[k][0], j + direction[k][1]
        if x < 0 || x >= m || y < 0 || y >= n {
            continue
        }

        if visited[x][y] {
            continue
        }

        if matrix[x][y] < matrix[i][j] {
            continue
        }

        dfs(x, y, matrix, visited)
    }
}
