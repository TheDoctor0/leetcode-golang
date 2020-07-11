package solutions

func containsDuplicate(nums []int) bool {
    set := make(map[int]struct{})

    for _, value := range nums {
        if _, ok := set[value]; ok {
            return true
        } else {
            set[value] = struct{}{}
        }
    }

    return false
}
