package solutions

func recoverTree(root *TreeNode)  {
    var previous, first, second *TreeNode = nil, nil, nil
    swapped := false

    traverseRecoveredTree(root, &previous, &first, &second, &swapped)

    if !swapped {
        (*first).Val, (*second).Val = (*second).Val, (*first).Val
    }
}

func traverseRecoveredTree(root *TreeNode, previous, first, second **TreeNode, swapped *bool) {
    if root == nil {
        return
    }

    traverseRecoveredTree(root.Left, previous, first, second, swapped)

    if *swapped {
        return
    }

    if *previous != nil && *first == nil && (*previous).Val > root.Val {
        *first = *previous
        *second = root
    } else if *first != nil && (*previous).Val > root.Val {
        (*first).Val, root.Val = root.Val, (*first).Val

        *swapped = true

        return
    }

    *previous = root

    traverseRecoveredTree(root.Right, previous, first, second, swapped)
}
