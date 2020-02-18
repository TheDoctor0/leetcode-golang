package solutions

func groupAnagrams(words []string) [][]string {
    cache := make(map[[26]byte]int)
    result := make([][]string, 0)

    for word := range words {
        list := [26]byte{}

        for letter := range words[word] {
            list[words[word][letter] - 'a']++
        }

        if i, ok := cache[list]; ok {
            result[i] = append(result[i], words[word])
        } else {
            result = append(result, []string{words[word]})

            cache[list] = len(result) - 1
        }
    }

    return result
}
