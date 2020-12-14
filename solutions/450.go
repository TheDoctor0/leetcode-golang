package solutions

func deleteNode(root *TreeNode, key int) *TreeNode {
    dummy := &TreeNode{Left:root}
    previous := dummy
    pointer := root

    for pointer != nil && pointer.Val != key {
        previous = pointer

        if pointer.Val > key {
            pointer = pointer.Left
        } else if pointer.Val < key {
            pointer = pointer.Right
        }
    }

    if pointer == nil {
        return root
    }

    delete(previous, pointer)

    return dummy.Left
}

func delete(previous *TreeNode, pointer *TreeNode) {
    if pointer.Left == nil {
        if previous.Left == pointer {
            previous.Left = pointer.Right
        } else {
            previous.Right = pointer.Right
        }

        return
    }

    if pointer.Right == nil {
        if previous.Left == pointer {
            previous.Left = pointer.Left
        } else {
            previous.Right = pointer.Left
        }
        return
    }

    origin := pointer
    previous, pointer = pointer, pointer.Left

    for pointer.Right != nil {
        previous = pointer
        pointer = pointer.Right
    }

    origin.Val, pointer.Val = pointer.Val, origin.Val

    delete(previous, pointer)
}
