package solutions

func canCross(stones []int) bool {
    cache := make(map[int]map[int]struct{})

    for _, stone := range stones {
        cache[stone] = make(map[int]struct{})
    }

    cache[0][1] = struct{}{}

    for i := 0; i < len(stones) - 1; i++ {
        for step := range cache[stones[i]] {
            reach := step + stones[i]

            if reach == stones[len(stones) - 1] {
                return true
            }

            if set, ok := cache[reach]; ok {
                set[step] = struct{}{}

                if step > 1 {
                    set[step - 1] = struct{}{}
                }

                set[step + 1] = struct{}{}
            }
        }
    }

    return false
}
