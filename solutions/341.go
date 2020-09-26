package solutions

type NestedIterator struct {
    list []*NestedInteger
    i       int
    current *NestedIterator
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{list: nestedList}
}

func (this *NestedIterator) Next() int {
    if this.list[this.i].IsInteger() {
        result := this.list[this.i].GetInteger()
        this.i++

        return result
    }

    return this.current.Next()
}

func (this *NestedIterator) HasNext() bool {
    if this.i >= len(this.list) {
        return false
    }

    if this.list[this.i].IsInteger() {
        return true
    }

    if this.current == nil {
        this.current = Constructor(this.list[this.i].GetList())
    }

    if this.current.HasNext() {
        return true
    }

    this.i++
    this.current = nil

    return this.HasNext()
}
