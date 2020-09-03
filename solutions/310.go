package solutions

func findMinHeightTrees(n int, edges [][]int) []int {
    graph := make([][]int, n)

    for _, value := range edges {
        graph[value[0]] = append(graph[value[0]], value[1])
        graph[value[1]] = append(graph[value[1]], value[0])
    }

    firstVisited := make([]bool, n)
    firstRoute := findMinHeightTree(graph, 0, firstVisited)

    secondVisited := make([]bool, n)
    secondRoute := findMinHeightTree(graph, firstRoute[0], secondVisited)

    return secondRoute[(len(secondRoute) - 1) / 2 : len(secondRoute) / 2 + 1]
}

func findMinHeightTree(graph [][]int, n int, visited []bool) []int {
    var result []int
    visited[n] = true

    for _, value := range graph[n] {
        if !visited[value] {
            route := findMinHeightTree(graph, value, visited)

            if len(route) > len(result) {
                result = route
            }
        }
    }

    return append(result, n)
}
