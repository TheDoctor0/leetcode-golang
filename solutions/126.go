package solutions

type graphNode struct {
    parentNode *graphNode
    word       string
}
type void struct{}

var member void
var possibleStates = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func findLadders(startWord string, endWord string, words []string) (ladders [][]string) {
    dictionary := make(map[string]void)

    for wordIndex := range words {
        dictionary[words[wordIndex]] = member
    }

    if _, ok := dictionary[endWord]; !ok || startWord == endWord {
        return make([][]string, 0)
    }

    return createLadders(bfsLadders(startWord, endWord, dictionary))
}

func bfsLadders(startWord string, endWord string, dictionary map[string]void) (map[string][]graphNode, map[string][]graphNode) {
    sourceGraphNode, targetGraphNode := graphNode{word: startWord}, graphNode{word: endWord}
    sourceMap, targetMap := make(map[string][]graphNode), make(map[string][]graphNode)
    found, forward, backwards := false, true, true

    sourceMap = addToMap(sourceMap, sourceGraphNode)
    targetMap = addToMap(targetMap, targetGraphNode)

    for !found && (len(sourceMap) > 0 || len(targetMap) > 0) {
        forward, backwards = getDirection(sourceMap, targetMap)

        if forward {
            sourceMap = traverseNextLevel(sourceMap, endWord, dictionary)
        } else if backwards {
            targetMap = traverseNextLevel(targetMap, startWord, dictionary)
        } else {
            continue
        }

        found, sourceMap, targetMap = hasConverged(sourceMap, targetMap)
    }

    return sourceMap, targetMap
}

func traverseNextLevel(sourceMap map[string][]graphNode, endWord string, dictionary map[string]void) map[string][]graphNode {
    targetMap := make(map[string][]graphNode, 0)

    for _, nodes := range sourceMap {
        for nodeIndex := range nodes {
            currentNode := nodes[nodeIndex]

            delete(dictionary, currentNode.word)

            nextStatesLength := len(currentNode.word) * len(possibleStates) - 1

            if nextStatesLength > len(dictionary) {
                states := nextStatesFromDictionary(currentNode, dictionary)

                for stateIndex := range states {
                    targetMap = addToMap(targetMap, states[stateIndex])
                }
            } else {
                for index := 0; index < len(currentNode.word); index++ {
                    states := nextStates(currentNode, index, endWord, dictionary)

                    for _, newNodes := range states {
                        for nodeIndex := range newNodes {
                            targetMap = addToMap(targetMap, newNodes[nodeIndex])
                        }
                    }
                }
            }
        }
    }

    return targetMap
}

func getDirection(sourceMap map[string][]graphNode, targetMap map[string][]graphNode) (bool, bool) {
    forwardDirection := (len(sourceMap) <= len(targetMap)) && len(sourceMap) != 0
    backwardsDirection := (len(sourceMap) > len(targetMap)) || len(sourceMap) == 0

    return forwardDirection, backwardsDirection
}

func nextStates(currentNode graphNode, index int, endWord string, dictionary map[string]void) (states map[string][]graphNode) {
    currentWord := currentNode.word
    mapNewStates := make(map[string][]graphNode)

    for stateIndex := range possibleStates {
        prefix, suffix := currentWord[:index], ""

        if index < len(currentWord) - 1 {
            suffix = currentWord[index + 1:]
        }

        newState := prefix + possibleStates[stateIndex] + suffix

        if newState == endWord {
            result := make(map[string][]graphNode, 0)
            result = addToMap(result, graphNode{parentNode: &currentNode, word: newState})

            return result
        }

        if _, ok := dictionary[newState]; ok {
            mapNewStates = addToMap(mapNewStates, graphNode{parentNode: &currentNode, word: newState})
        }
    }

    return mapNewStates
}

func hasConverged(sourceMap map[string][]graphNode, targetMap map[string][]graphNode) (bool, map[string][]graphNode, map[string][]graphNode) {
    converged := false

    for word := range sourceMap {
        if _, ok := targetMap[word]; ok {
            converged = true

            break
        }
    }

    if converged {
        for word := range sourceMap {
            if _, ok := targetMap[word]; !ok {
                delete(sourceMap, word)
            }
        }

        for word := range targetMap {
            if _, ok := sourceMap[word]; !ok {
                delete(targetMap, word)
            }
        }
    }

    return converged, sourceMap, targetMap
}

func addToMap(mapNode map[string][]graphNode, node graphNode) map[string][]graphNode {
    mapNode[node.word] = append(mapNode[node.word], node)

    return mapNode
}

func createLadders(sourceStates map[string][]graphNode, targetStates map[string][]graphNode) (ladders [][]string) {
    resultLadders := make([][]string, 0)

    for word, sourceNodes := range sourceStates {
        for sourceNodeIndex := range sourceNodes {
            sourceLadder := createLadder(&sourceNodes[sourceNodeIndex], true)
            targetNodes := targetStates[word]

            for targetNodeIndex := range targetNodes {
                targetLadder := createLadder(&targetNodes[targetNodeIndex], false)
                targetLadder = targetLadder[1:]

                resultLadders = append(resultLadders, append(sourceLadder, targetLadder...))
            }
        }
    }

    return resultLadders
}

func createLadder(node *graphNode, reverse bool) (ladder []string) {
    ladder = make([]string, 0)
    ladder = append(ladder, node.word)

    for {
        if node.parentNode == nil {
            break
        } else {
            node = node.parentNode
            ladder = append(ladder, node.word)
        }
    }

    if reverse {
        reversedLadder := make([]string, len(ladder))
        length := len(ladder) - 1

        for index := 0; index <= length; index++ {
            reversedLadder[length - index] = ladder[index]
        }

        return reversedLadder
    }

    return
}

func nextStatesFromDictionary(parent graphNode, dictionary map[string]void) []graphNode {
    nextStates := make([]graphNode, 0)

    for word := range dictionary {
        if checkDifference(word, parent.word) == 1 {
            nextStates = append(nextStates, graphNode{parentNode: &parent, word: word})
        }
    }

    return nextStates
}

func checkDifference(word1 string, word2 string) int {
    difference := 0

    for index := 0; index < len(word1); index++ {
        if word1[index] != word2[index] {
            difference++
        }
    }

    return difference
}
