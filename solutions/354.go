package solutions

import (
    "sort"
)

func maxEnvelopes(envelopes [][]int) int {
    envelopesCopy := envelopes

    sort.Slice(envelopesCopy, func (i int, j int) bool {
        if envelopesCopy[i][0] < envelopesCopy[j][0] {
            return true
        } else if envelopesCopy[i][0] > envelopesCopy[j][0] {
            return false
        } else {
            return envelopesCopy[i][1] > envelopesCopy[j][1]
        }
    })

    dp := make([]int, len(envelopesCopy))
    counter := 0

    for i := 0; i < len(envelopesCopy); i++  {
        position := sort.SearchInts(dp[:counter], envelopesCopy[i][1])
        dp[position] = envelopesCopy[i][1]

        if position == counter {
            counter++
        }
    }

    return counter
}
