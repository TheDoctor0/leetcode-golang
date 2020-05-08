package solutions

func sumNumbers(root *TreeNode) int {
    return recursiveSumNumbers(root, 0)
}

func recursiveSumNumbers(node *TreeNode, sum int) int {
    if node == nil {
        return 0
    }

    sum = sum * 10 + node.Val

    if node.Left == nil && node.Right == nil {
        return sum
    }

    left := 0

    if node.Left != nil {
        left = recursiveSumNumbers(node.Left, sum)
    }

    right := 0

    if node.Right != nil {
        right = recursiveSumNumbers(node.Right, sum)
    }

    return left + right
}
