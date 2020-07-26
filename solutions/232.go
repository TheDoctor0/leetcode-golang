package solutions

type MyQueue struct {
    array []int
}

func Constructor() MyQueue {
    return MyQueue{}
}

func (this *MyQueue) Push(x int) {
    this.array = append(this.array, x)
}

func (this *MyQueue) Pop() int {
    if this.Empty() {
        return 0
    }

    top := this.Peek()
    this.array = this.array[1: len(this.array)]

    return top
}

func (this *MyQueue) Peek() int {
    if this.Empty() {
        return 0
    }

    return this.array[0]
}

func (this *MyQueue) Empty() bool {
    return len(this.array) == 0
}
