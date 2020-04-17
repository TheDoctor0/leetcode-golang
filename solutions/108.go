package solutions

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    middle := len(nums) / 2
    result := new(TreeNode)
    result.Val = nums[middle]

    if middle != 0 {
        result.Left = sortedArrayToBST(nums[:middle])
        result.Right = sortedArrayToBST(nums[middle + 1:])
    }

    return result
}
