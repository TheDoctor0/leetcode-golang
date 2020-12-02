package solutions

func countSegments(s string) int {
    count := 0
    previous := ' '

    for _, current := range s {
        if previous == ' ' && current != ' ' {
            count++
        }

        previous = current
    }

    return count
}
