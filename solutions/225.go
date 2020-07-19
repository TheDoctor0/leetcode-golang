package solutions

type MyStack struct {
    queue []int
}

func Constructor() MyStack {
    return MyStack{}
}

func (this *MyStack) Push(x int) {
    this.queue = append(this.queue, x)

    for i := 1; i < len(this.queue); i++ {
        top := this.queue[0]

        this.queue = this.queue[1:]
        this.queue = append(this.queue, top)
    }
}

func (this *MyStack) Pop() int {
    top := this.queue[0]

    this.queue = this.queue[1:]

    return top
}

func (this *MyStack) Top() int {
    return this.queue[0]
}

func (this *MyStack) Empty() bool {
    return len(this.queue) == 0
}
