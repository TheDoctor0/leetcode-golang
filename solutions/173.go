package solutions

func Constructor(root *TreeNode) BSTIterator {
    result :=  BSTIterator{stack: []*TreeNode{}}

    for root != nil {
        result.stack = append(result.stack, root)
        root = root.Left
    }

    return result
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    top := this.stack[len(this.stack) - 1]
    result := top.Val
    this.stack = this.stack[: len(this.stack) - 1]

    if top.Right != nil {
        top = top.Right

        for top != nil {
            this.stack = append(this.stack, top)
            top = top.Left
        }
    }

    return result
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.stack) != 0
}
