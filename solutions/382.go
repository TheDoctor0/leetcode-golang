package solutions

import (
    "math/rand"
)

type Solution struct {
    node *ListNode
}

func Constructor(head *ListNode) Solution {
    return Solution{node: head}
}

func (this *Solution) GetRandom() int {
    result := this.node.Val
    pointer := this.node
    counter := 1

    for pointer.Next != nil {
        if rand.Intn(counter + 1) == 0 {
            result = pointer.Next.Val
        }

        counter++
        pointer = pointer.Next
    }

    return result
}
