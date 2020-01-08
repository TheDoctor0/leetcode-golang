package solutions

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    source := x
    var reversed int

    for x > 0 {
        reversed = (reversed * 10) + (x % 10)

        x = x / 10
    }

    return reversed == source
}
