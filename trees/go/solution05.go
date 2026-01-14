package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	node := make(nums[mid] * TreeNode)

	node.Left

}
