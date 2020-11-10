package solutions

import (
    "strconv"
)

func readBinaryWatch(num int) []string {
    hours, minutes := generateValueMap(12), generateValueMap(60)
    var result []string

    for  i := 0; i <= num; i++ {
        for _, hour := range hours[i] {
            for _, minute := range minutes[num - i] {
                current := strconv.Itoa(hour) + ":"

                if minute < 10 {
                    current += "0"
                }

                current += strconv.Itoa(minute)

                result = append(result, current)
            }
        }
    }

    return result
}

func generateValueMap(limit int) map[int][]int{
    valueMap := map[int][]int{}

    for i := 0; i < limit; i++ {
        count :=0
        k := i

        for k > 0 {
            if k % 2 == 1 {
                count++
            }

            k /= 2
        }

        if _, ok := valueMap[count]; !ok {
            valueMap[count] = []int{}
        }

        valueMap[count] = append(valueMap[count], i)
    }

    return valueMap
}
