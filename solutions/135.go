package solutions

func candy(ratings []int) int {
    length := len(ratings)

    if length == 0 {
        return 0
    }

    totalCandies := 1
    leftNeighbour := 1

    for i := 1; i < length; {
        if ratings[i] > ratings[i - 1] {
            leftNeighbour = leftNeighbour + 1
            totalCandies += leftNeighbour

            i++
        } else if ratings[i] == ratings[i-1] {
            leftNeighbour = 1
            totalCandies += 1

            i++
        } else {
            count := 0

            for ; i < length && ratings[i - 1] > ratings[i]; i++ {
                count++

                totalCandies += count
            }

            if count >= leftNeighbour {
                totalCandies += count - leftNeighbour + 1
            }

            leftNeighbour = 1
        }
    }

    return totalCandies
}
