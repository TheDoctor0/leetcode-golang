package solutions

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }

    nodeA, nodeB := headA, headB

    for nodeA != nil || nodeB != nil {
        if nodeA == nodeB {
            return nodeA
        }

        if nodeA == nil {
            nodeA = headB
            nodeB = nodeB.Next

            continue
        }

        if nodeB == nil {
            nodeB = headA
            nodeA = nodeA.Next

            continue
        }

        nodeA = nodeA.Next
        nodeB = nodeB.Next
    }

    return nil
}
