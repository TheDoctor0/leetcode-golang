package solutions

func isMatch(s string, p string) bool {
    stringIndex, patternIndex, stringReverseIndex := 0, 0, 0
    asteriskIndex := -1

    for stringIndex < len(s) {
        match := patternIndex < len(p) && (s[stringIndex] == p[patternIndex] || p[patternIndex] == '?')

        if match {
            patternIndex += 1
            stringIndex += 1
        } else if patternIndex < len(p) && p[patternIndex] == '*' {
            asteriskIndex = patternIndex
            stringReverseIndex = stringIndex

            patternIndex += 1
        } else if asteriskIndex != -1 {
            patternIndex = asteriskIndex + 1

            stringIndex = stringReverseIndex + 1
            stringReverseIndex = stringIndex
        } else {
            return false
        }
    }

    for patternIndex < len(p) && p[patternIndex] == '*' {
        patternIndex += 1
    }

    return patternIndex == len(p)
}
