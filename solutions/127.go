package solutions

type WordNode struct {
    word string
    state int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    length := len(beginWord)
    dictionary := buildDictionary(wordList, length)
    queue := []WordNode{{word: beginWord, state: 1}}
    set := make(map[string]struct{})
    set[beginWord] = struct{}{}

    for len(queue) != 0 {
        currentNode := queue[0]
        queue = queue[1:]

        for i := 0; i < length; i++ {
            intermediateWord := findIntermediateWord(currentNode, i)

            for j := 0; j < len(dictionary[intermediateWord]); j++ {
                if dictionary[intermediateWord][j] == endWord {
                    return currentNode.state + 1
                }

                if _, ok := set[dictionary[intermediateWord][j]]; !ok {
                    set[dictionary[intermediateWord][j]] = struct{}{}

                    queue = append(queue, WordNode{word: dictionary[intermediateWord][j], state: currentNode.state + 1})
                }
            }
        }
    }

    return 0
}

func findIntermediateWord(currentNode WordNode, i int) string {
    return currentNode.word[: i] + "*" + currentNode.word[i + 1:]
}

func buildDictionary(wordList []string, length int) map[string][]string {
    dictionary := make(map[string][]string)

    for i := 0; i < len(wordList); i++ {
        for j := 0; j < length; j++ {
            dictionary[wordList[i][: j] + "*" + wordList[i][j + 1:]] = append(dictionary[wordList[i][: j] + "*" + wordList[i][j + 1:]], wordList[i])
        }
    }

    return dictionary
}
