package solutions

func titleToNumber(title string) int {
    var result int

    for _, character := range title {
        result = result * 26 + int(character - 'A' + 1)
    }

    return result
}
