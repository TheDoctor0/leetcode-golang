package solutions

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    length := len(queries)
    result := make([]float64, length)
    visited := make(map[int]bool)

    for i := 0; i < length; i++ {
        result[i] = calculateEquation(equations, values, visited, queries[i][0], queries[i][1], 1.0)
    }

    return result
}

func calculateEquation(equations [][]string, values []float64, visited map[int]bool, a, b string, answer float64) float64 {
    for i, value := range equations {
        if visited[i] {
            continue
        }

        if value[0] == a && a == b || value[1] == a && a == b {
            return 1.0
        } else if value[0] == a {
            if value[1] == b {
                return answer * values[i]
            }

            visited[i] = true
            k := calculateEquation(equations, values, visited, value[1], b, answer * values[i])
            visited[i] = false

            if k != -1.0 {
                return k
            }
        } else if value[1] == a && values[i] != 0.0 {
            if value[0] == b {
                return answer / values[i]
            }

            visited[i] = true
            k := calculateEquation(equations, values, visited, value[0], b, answer / values[i])
            visited[i] = false

            if k != -1.0 {
                return k
            }
        }
    }

    return -1.0
}
