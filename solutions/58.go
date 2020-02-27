package solutions

func lengthOfLastWord(s string) int {
    space, result := 0, 0

    for i, char := range s {
        if char == ' ' {
            space = i + 1
        } else {
            result = i + 1 - space
        }
    }

    return result
}
