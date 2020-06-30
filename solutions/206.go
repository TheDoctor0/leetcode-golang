package solutions

func reverseList(head *ListNode) *ListNode {
    var result *ListNode

    for head != nil {
        temp := head.Next
        head.Next = result
        result = head
        head = temp
    }

    return result
}
