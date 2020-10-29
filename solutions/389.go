package solutions

func findTheDifference(s string, t string) byte {
    var difference byte

    for i := range s {
        difference ^= s[i]
        difference ^= t[i]
    }

    difference ^= t[len(t) - 1]

    return difference
}
