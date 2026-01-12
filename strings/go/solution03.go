package leetcode

func firstUniqChar(s string) int {

	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		c := s[i]
		count[c-'a']++
	}

	for i := 0; i < len(s); i++ {
		c := s[i]

		if count[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
