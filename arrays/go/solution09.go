package leetcode

func twoSum(nums []int, target int) []int {

	indexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, ok := indexMap[complement]; ok {
			return []int{j, i}
		}

		indexMap[num] = i
	}
	return nil
}
