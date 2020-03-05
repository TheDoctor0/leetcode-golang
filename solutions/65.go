package solutions

func isEmpty(char rune) bool {
    return char == 9 || char == 32
}

func isSymbol(char rune) bool {
    return char == 45 || char == 43
}

func isDigit(char rune) bool {
    return char >= 48 && char <= 57
}

func isFloat(result []rune, exp bool) bool {
    start := 0
    end := len(result)

    if end <= start {
        return false
    }

    point := result[start] == 46

    if !isSymbol(result[start]) && !isDigit(result[start]) && !point {
        return false
    }

    if isSymbol(result[start]) {
        if end - start <= 1 {
            return false
        }

        start++
    }

    point = result[start] == 46

    if point {
        if exp {
            return false
        }

        if end - start <= 1 {
            return false
        }

        point = true
        start++
    }

    if !isDigit(result[start]) {
        return false
    }

    for start < end {
        if result[start] == 101 {
            if exp {
                return false
            }

            start++

            if start >= end {
                return false
            }

            return isFloat(result[start: end], true)
        }

        if result[start] == 46 {
            if exp {
                return false
            }
            
            if point {
                return false
            }

            point = true
            start++

            continue
        }
        
        if !isDigit(result[start]) {
            return false
        }

        start++
    }

    return true
}

func isNumber(s string) bool {
    result := []rune(s)
    start := 0
    end := len(result)

    for start < end {
        if isEmpty(result[start]) {
            start++

            continue
        }

        break
    }

    for end > 0 {
        if isEmpty(result[end - 1]) {
            end--

            continue
        }

        break
    }

    if end <= start {
        return false
    }

    return isFloat(result[start:end], false)
}
