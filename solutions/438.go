package solutions

func findAnagrams(s string, p string) []int {
    hash := 0

    for i := 0; i < len(p); i++  {
        hash += 1 << (p[i] - 'a')
    }

    var result []int
    i, j, currentHash := 0, 0, 0

    for j < len(s) {
        currentHash += 1 << (s[j] - 'a')

        if j - i + 1 == len(p) {
            if currentHash == hash {
                result = append(result, i)
            }

            currentHash -= 1 << (s[i] - 'a')
            i++
        }

        j++
    }

    return result
}
