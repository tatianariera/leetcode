package leetcode

func strStr(haystack string, needle string) int {

	n := len(haystack)
	m := len(needle)

	if m == 0 {
		return 0
	}

	for i := 0; i <= n-m; i++ {
		var j int

		for j = 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}

		if j == m {
			return i
		}
	}
	return -1
}
