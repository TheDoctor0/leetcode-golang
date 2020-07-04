package solutions

func findOrder(numCourses int, prerequisites [][]int) []int {
    if numCourses == 0 {
        return []int{}
    }

    graph := make(map[int][]int)

    for _, prerequisite := range prerequisites {
        want, need := prerequisite[0], prerequisite[1]
        graph[want] = append(graph[want], need)
    }

    visited := make(map[int]bool)
    using := make(map[int]bool)
    order := make([]int, 0)

    for i := 0; i < numCourses; i++ {
        if visited[i] {
            continue
        }

        if hasCycles(i, graph, visited, using, &order) {
            return []int{}
        }
    }

    return order
}

func hasCycles(subject int, graph map[int][]int, visited, using map[int]bool, order *[]int) bool {
    if using[subject] {
        return true
    }

    if visited[subject] {
        return false
    }

    using[subject] = true

    for _, need := range graph[subject] {
        if hasCycles(need, graph, visited, using, order) {
            return true
        }
    }

    visited[subject] = true
    using[subject] = false
    *order = append(*order, subject)

    return false
}
