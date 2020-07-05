package solutions

type WordDictionary struct {
    children [26]*WordDictionary
    isEnd bool
}

func Constructor() WordDictionary {
    return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string)  {
    current := this

    for i := 0; i < len(word); i++ {
        index := word[i] - 'a'

        if current.children[index] == nil {
            current.children[index] = &WordDictionary{}
        }

        current = current.children[index]
    }

    current.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    if len(word) == 0 {
        return this.isEnd
    }

    if word[0] == '.' {
        for j := 0; j < 26; j++ {
            if this.children[j] != nil && this.children[j].Search(word[1:]) {
                return true
            }
        }
    } else {
        index := word[0] - 'a'

        if child := this.children[index]; child != nil {
            return child.Search(word[1:])
        }
    }

    return false
}
