package solutions

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }

    result := make([]string, 0)

    combinations(&result, digits, "", 0)

    return result
}

func combinations(result *[]string, digits string, letters string, i int) {
    if i == len(digits) {
        *result = append(*result, letters)

        return
    }

    var hash = map[uint8]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    for _, c := range hash[digits[i]] {
        combinations(result, digits, letters + string(c), i + 1)
    }
}
