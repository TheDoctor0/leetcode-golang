package solutions

type PeekingIterator struct {
    numbers  []int
    iterator *Iterator
}

func Constructor(iterator *Iterator) *PeekingIterator {
    return &PeekingIterator{make([]int, 0), iterator}
}

func (this *PeekingIterator) hasNext() bool {
    if len(this.numbers) > 0 || this.iterator.hasNext() {
        return true
    }

    return false
}

func (this *PeekingIterator) next() int {
    if len(this.numbers) > 0 {
        top := this.numbers[0]
        this.numbers = this.numbers[1:]

        return top
    }

    return this.iterator.next()
}

func (this *PeekingIterator) peek() int {
    if len(this.numbers) == 0 {
        this.numbers = append(this.numbers, this.iterator.next())
    }

    return this.numbers[0]
}
