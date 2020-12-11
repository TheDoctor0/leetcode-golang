package solutions

func numberOfBoomerangs(points [][]int) int {
    result := 0

    for i := 0; i < len(points); i++ {
        dictionary := make(map[int]int)

        for j := 0; j < len(points); j++ {
            distance := calculateDistance(points[i], points[j])

            if value, ok := dictionary[distance]; !ok {
                dictionary[distance]++
            } else {
                result += 2 * value
                dictionary[distance]++
            }
        }
    }

    return result
}

func calculateDistance(a []int, b []int) int {
    x := a[0] - b[0]
    y := a[1] - b[1]

    return x * x + y * y
}
