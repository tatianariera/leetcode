package leetcode

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {

	i := 0
	n := len(s)

	for i < n && s[i] == ' ' {
		i++
	}

	sign := 1

	if i < n {
		switch s[i] {
		case '-':
			sign = -1
			i++
		case '+':
			i++
		}
	}

	var result int64 = 0

	for i < n && unicode.IsDigit(rune(s[i])) {
		digit := s[i] - '0'
		result = result*10 + int64(digit)

		if sign == 1 && result > math.MaxInt32 {
			return math.MaxInt32
		}

		if sign == -1 && result < math.MinInt32 {
			return math.MinInt32
		}

		i++
	}

	return int(sign * int(result))
}
