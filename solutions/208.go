package solutions

type Trie struct {
    children [26]*Trie
    isEndOfWord bool
    root *Trie
}

func Constructor() Trie {
    return Trie{children: [26]*Trie{}, isEndOfWord: false, root: &Trie{}}
}

func (this *Trie) Insert(word string)  {
    if len(word) == 0 {
        return
    }

    current := this.root

    for i := 0; i < len(word); i++ {
        index := word[i] - 'a'

        if current.children[index] == nil {
            current.children[index] = &Trie{}
        }

        current = current.children[index]
    }

    current.isEndOfWord = true
}

func (this *Trie) Search(word string) bool {
    current := this.root

    for i := 0; i< len(word); i++ {
        index := word[i] - 'a'

        if current.children[index] == nil {
            return false
        }

        current = current.children[index]
    }

    if current.isEndOfWord == true {
        return true
    }

    return false
}

func (this *Trie) StartsWith(prefix string) bool {
    current := this.root

    if len(prefix) == 1 {
        if current.children[prefix[0] - 'a'] != nil {
            return true
        }
    }

    depth := 0

    for i := 0; i < len(prefix); i++ {
        index := prefix[i] - 'a'
        depth++

        if current.children[index] == nil {
            return false
        }

        if depth == len(prefix) {
            return true
        }

        current = current.children[index]
    }

    for i := 0; i < len(current.children); i++ {
        if current.children[i] != nil {
            return true
        }
    }

    return false
}
