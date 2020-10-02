package solutions

func intersection(nums1 []int, nums2 []int) []int {
    hashMap := make(map[int]bool)
    var result []int

    for i := range nums1 {
        hashMap[nums1[i]] = false
    }

    for i := range nums2 {
        if value, ok := hashMap[nums2[i]]; ok {
            if value {
                continue
            }

            result = append(result, nums2[i])
            hashMap[nums2[i]] = true
        }
    }

    return result
}
