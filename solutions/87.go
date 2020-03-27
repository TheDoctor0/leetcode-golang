package solutions

func isScramble(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }

    if len(s1) != len(s2){
        return false
    }

    letters := make([]int, 26)

    for i := 0; i < len(s1); i++ {
        letters[int(s1[i]) - int('a')]++
        letters[int(s2[i]) - int('a')]--
    }

    for i := 0; i < len(s1); i++ {
        if letters[int(s1[i]) - int('a')] != 0 {
            return false
        }
    }

    for i := 1; i < len(s1); i++ {
        if isScramble(s1[0: i], s2[0: i]) && isScramble(s1[i:], s2[i:]) {
            return true
        }

        if isScramble(s1[0: i], s2[len(s2) - i:]) && isScramble(s1[i:], s2[0: len(s2) - i]) {
            return true
        }
    }

    return false
}
