package leetcode

import "sort"

func RemoveDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	k := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
