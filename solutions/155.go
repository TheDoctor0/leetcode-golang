package solutions

type MinStack struct {
    Stack []int
    Min []int
}

func Constructor() MinStack {
    return MinStack{Stack: []int{}, Min: []int{}}
}

func (stack *MinStack) Push(x int)  {
    stack.Stack = append(stack.Stack, x)

    if len(stack.Min) == 0 {
        stack.Min = append(stack.Min, 0)

        return
    }

    if stack.Stack[stack.Min[len(stack.Min) - 1]] < x {
        stack.Min = append(stack.Min, stack.Min[len(stack.Min) - 1])
    } else {
        stack.Min = append(stack.Min, len(stack.Stack) - 1)
    }
}

func (stack *MinStack) Pop()  {
    stack.Stack = stack.Stack[: len(stack.Stack) - 1]
    stack.Min = stack.Min[: len(stack.Min) - 1]
}

func (stack *MinStack) Top() int {
    return stack.Stack[len(stack.Stack) - 1]
}

func (stack *MinStack) GetMin() int {
    return stack.Stack[stack.Min[len(stack.Min) - 1]]
}
