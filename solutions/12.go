package solutions

func intToRoman(num int) string {
    roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    values := []int{1000,  900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

    var result string

    for key, value := range values {
        var multiply = num / value

        for i := 0; i < multiply; i++ {
            num -= value

            result += roman[key]
        }
    }

    return result
}