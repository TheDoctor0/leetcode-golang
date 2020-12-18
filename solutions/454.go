package solutions

func fourSumCount(A []int, B []int, C []int, D []int) int {
    dictionary := make(map[int]int)
    length := len(A)
    result := 0

    for i := 0; i < length; i++ {
        for j := 0; j < length; j++ {
            dictionary[A[i] + B[j]]++
        }
    }

    for i := 0; i < length; i++ {
        for j := 0; j < length; j++ {
            result += dictionary[-C[i]-D[j]]
        }
    }

    return result
}
