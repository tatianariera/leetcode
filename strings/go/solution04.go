package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		idx := t[i] - 'a'
		count[idx]--

		if count[idx] < 0 {
			return false
		}
	}

	return true
}
