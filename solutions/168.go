package solutions

func convertToTitle(number int) string {
    title := ""

    for number > 0 {
        value := number % 26
        number = number / 26

        if value == 0 {
            value = 26
            number -= 1
        }

        title = string('A' + byte(value - 1)) + title
    }

    return title
}
