package solutions

func reverseString(s []byte)  {
    length := len(s)

    for i := 0; i < length / 2; i++ {
        s[i], s[length - 1 - i] = s[length - 1 - i], s[i]
    }
}
