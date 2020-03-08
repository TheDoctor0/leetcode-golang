package solutions

import (
    "bytes"
    "strings"
)

func fullJustify(words []string, maxWidth int) []string {
    var current, result []string
    currentLength := 0

    for index, word := range words {
        if currentLength + len(current) + len(word) > maxWidth {
            if len(current) == 1 {
                currentLine := current[0] + strings.Repeat(" ", maxWidth - len(current[0]))

                result = append(result, currentLine)
            } else {
                difference := maxWidth - currentLength
                spaces := difference / (len(current) - 1)
                more := difference % (len(current) - 1)
                
                currentLine := bytes.Buffer{}

                for currentIndex, currentWord := range current {
                    currentLine.WriteString(currentWord)

                    if currentIndex != len(current) - 1 {
                        blanks := 0

                        if more > 0 {
                            blanks = 1
                            more--
                        }

                        currentLine.WriteString(strings.Repeat(" ", spaces + blanks))
                    }
                }

                result = append(result, currentLine.String())
            }

            current, currentLength = []string{}, 0
        }

        currentLength += len(word)
        current = append(current, word)

        if index == len(words) - 1 {
            lastLine := strings.Join(current, " ")
            lastLine = lastLine + strings.Repeat(" ", maxWidth - len(lastLine))

            result = append(result, lastLine)
        }
    }

    return result
}
