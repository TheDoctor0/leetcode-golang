package solutions

func minWindow(s string, t string) string {
    if s == "" || t == "" {
        return ""
    }

    target, count := [128]int{}, [128]int{}
    required, current, min := 0, 0, len(s) + 1
    result := ""

    for i := range t {
        if target[t[i]] == 0 {
            required++
        }

        target[t[i]]++
    }

    for left, right := 0, 0; right < len(s); right++ {
        char := s[right]
        count[char]++

        if count[char] == target[char] {
            current++
        }

        for left <= right && current == required {
            char = s[left]

            if right - left + 1 < min {
                min = right - left + 1
                result = s[left : right+1]
            }

            count[char]--

            if count[char] < target[char] {
                current--
            }

            left++
        }
    }

    return result
}
