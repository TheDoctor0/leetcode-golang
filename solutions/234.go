package solutions

func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }

    middle, end := head, head

    for end.Next != nil {
        end = end.Next

        if end.Next == nil {
            break
        }

        end, middle = end.Next, middle.Next
    }

    var previous *ListNode

    current := middle

    for current != nil {
        current.Next, previous, current = previous, current, current.Next
    }

    middle.Next = nil

    for head != nil {
        if end.Val != head.Val {
            return false
        }

        end, head = end.Next, head.Next
    }

    return true
}
