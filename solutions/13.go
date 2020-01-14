package solutions

func romanToInt(s string) int {
    if len(s) == 0 {
        return 0
    }

    var roman = map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    previous := rune(s[len(s) - 1])
    integer := roman[previous]

    for i := len(s) - 2; i >= 0; i-- {
        char := rune(s[i])

        if roman[previous] <= roman[char] {
            integer += roman[char]
        } else {
            integer -= roman[char]
        }

        previous = char
    }

    return integer
}
