package solutions

import (
    "fmt"
)

func getHint(secret string, guess string) string {
    var countSecret, countGuess [10]int
    bulls, cows := 0,0

    for index := range secret {
        secretDigit := int(secret[index] - '0')
        guessDigit := int(guess[index] - '0')

        if guessDigit == secretDigit {
            bulls++
            continue
        }

        if countGuess[secretDigit] > 0 {
            cows++
            countGuess[secretDigit]--
        } else {
            countSecret[secretDigit]++
        }

        if countSecret[guessDigit] > 0 {
            cows++
            countSecret[guessDigit]--
        } else {
            countGuess[guessDigit]++
        }
    }

    return fmt.Sprintf("%dA%dB", bulls, cows)
}
