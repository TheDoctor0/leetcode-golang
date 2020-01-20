package solutions

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    resultNode := ListNode{0, nil}
    currentNode := &resultNode
    carry := 0

    for l1 != nil || l2 != nil {
        var first, second int

        if l1 != nil {
            first = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            second = l2.Val
            l2 = l2.Next
        }

        sum := first + second + carry
        carry = sum / 10
        currentNode.Next = &ListNode{sum % 10, nil}
        currentNode = currentNode.Next
    }

    if carry != 0 {
        currentNode.Next = &ListNode{1, nil}
    }

    return resultNode.Next
}
