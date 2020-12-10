package solutions

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var firstList, secondList []int

    result := &ListNode{}

    for l1 != nil {
        firstList = append([]int{l1.Val}, firstList...)

        l1 = l1.Next
    }

    for l2 != nil{
        secondList = append([]int{l2.Val}, secondList...)

        l2 = l2.Next
    }

    i, carry := 0, 0

    for i < len(firstList) || i < len(secondList){
        total := carry

        if i < len(firstList) {
            total += firstList[i]
        }

        if i < len(secondList) {
            total += secondList[i]
        }

        i++

        node := &ListNode{
            Val : total % 10,
            Next : result.Next,
        }
        carry = total / 10
        result.Next = node
    }

    if carry != 0 {
        node := &ListNode{
            Val : carry,
            Next : result.Next,
        }
        result.Next = node
    }

    return result.Next
}
