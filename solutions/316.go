package solutions

func removeDuplicateLetters(s string) string {
    lastInstanceOf, seen := [26]int{}, [26]bool{}

    for i := range s {
        lastInstanceOf[s[i] - 'a'] = i
    }

    var result []rune

    for i, letter := range s {
        for len(result) > 0 {
            lastRune := result[len(result) - 1]

            if letter <= lastRune && lastInstanceOf[lastRune -'a'] >= i && !seen[letter - 'a'] {
                seen[lastRune - 'a'] = false
                result = result[:len(result) - 1]

                continue
            }

            break
        }

        if !seen[letter - 'a'] {
            result = append(result, letter)
            seen[letter - 'a'] = true
        }
    }

    return string(result)
}
