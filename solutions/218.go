package solutions

func getSkyline(buildings [][]int) [][]int {
    length := len(buildings)

    if length == 0 {
        return nil
    }

    if length == 1 {
        building := buildings[0]

        return [][]int{
            {building[0], building[2]},
            {building[1], 0},
        }
    }

    middle := length / 2

    return mergeSkyline(getSkyline(buildings[: middle]), getSkyline(buildings[middle:]))
}

func mergeSkyline(skyline1, skyline2 [][]int) [][]int {
    var result [][]int
    var i, j, x, y1, y2 int

    for i < len(skyline1) && j < len(skyline2) {
        if skyline1[i][0] == skyline2[j][0] {
            x = skyline1[i][0]
            y1 = skyline1[i][1]
            y2 = skyline2[j][1]

            i++
            j++
        } else if skyline1[i][0] < skyline2[j][0] {
            x = skyline1[i][0]
            y1 = skyline1[i][1]

            i++
        } else {
            x = skyline2[j][0]
            y2 = skyline2[j][1]

            j++
        }

        result = appendSkyline(result, x, max(y1, y2))
    }

    for ; i < len(skyline1); i++ {
        result = appendSkyline(result, skyline1[i][0], skyline1[i][1])
    }

    for ; j < len(skyline2); j++ {
        result = appendSkyline(result, skyline2[j][0], skyline2[j][1])
    }

    return result
}

func appendSkyline(skyline [][]int, x, y int) [][]int {
    length := len(skyline)

    if length == 0 || skyline[length - 1][1] != y {
        return append(skyline, []int{x, y})
    }

    return skyline
}
