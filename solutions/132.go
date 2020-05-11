package solutions

func minCut(s string) int {
    length := len(s)
    cut := make([]int, length)
    checked := make([][]bool, length)

    for i := range checked {
        checked[i] = make([]bool, length)
    }

    for i := 0; i < length; i++ {
        cut[i] = i

        for j :=0; j <= i; j++ {
            if s[j] == s[i] && (i - j <= 1 || checked[j + 1][i - 1]) {
                checked[j][i] = true

                if j > 0 {
                    cut[i] = min(cut[i],cut[j - 1] + 1)
                } else {
                    cut[i] = 0
                }
            }
        }
    }

    return cut[length - 1]
}
