package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	countMap := make(map[int]int)
	result := []int{}

	for _, num := range nums1 {
		countMap[num]++
	}

	for _, num := range nums2 {
		if countMap[num] > 0 {
			result = append(result, num)
			countMap[num]--
		}
	}
	return result
}
