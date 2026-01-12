package leetcode

func singleNumer(nums []int) int {

	resultado := 0

	for _, num := range nums {
		resultado ^= num
	}

	return resultado
}
