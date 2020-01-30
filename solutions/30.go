package solutions

func findSubstring(s string, words []string) []int {
    stringLength, wordsCount := len(s), len(words)
    var result []int

    if stringLength == 0 || wordsCount ==0 {
        return result
    }

    wordLength := len(words[0])
    remainingCount := make(map[string]int, wordsCount)

    for i := 0; i < wordLength; i++ {
        count := resetCounts(words, remainingCount)
        index := i

        for index + wordsCount * wordLength <= stringLength {
            word := s[index + count * wordLength : index + (count + 1) * wordLength]
            remaining, ok := remainingCount[word]

            if !ok {
                index += wordLength * (count + 1)

                if count != 0 {
                    count = resetCounts(words, remainingCount)
                }

                continue
            }

            if remaining == 0 {
                index = moveIndex(index, wordLength, s, remainingCount)
                count--

                continue
            }

            remainingCount[word]--

            if count == wordsCount - 1 {
                result = append(result, index)
                index = moveIndex(index, wordLength, s, remainingCount)
            } else {
                count++
            }
        }
    }

    return result
}

func resetCounts(words []string, remainNum map[string]int) int{
    for _, word := range words {
        remainNum[word] = 0
    }

    for _,word := range words{
        remainNum[word]++
    }

    return 0
}

func moveIndex(index int, wordLength int, s string, remainNum map[string]int) int{
    remainNum[s[index : index + wordLength]]++
    index += wordLength

    return index
}