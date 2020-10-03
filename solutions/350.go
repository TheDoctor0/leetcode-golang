package solutions

func intersect(nums1 []int, nums2 []int) []int {
    hashMap := make(map[int]int, len(nums1))
    var result []int

    for _, number := range nums1 {
        hashMap[number]++
    }

    for _, number := range nums2 {
        if value, ok := hashMap[number]; ok && value > 0 {
            result = append(result, number)
            hashMap[number]--
        }
    }

    return result
}
