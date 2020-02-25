// 链接：https://leetcode.com/problems/valid-number/
// 题意：给定一个字符串，判断是否为一个合法的数字？

// "0" => true
// " 0.1 " => true
// "abc" => false
// "1 a" => false
// "2e10" => true
// " -90e3   " => true
// " 1e" => false
// "e3" => false
// " 6e-1" => true
// " 99e2.5 " => false
// "53.5e93" => true
// " --6 " => false
// "-+3" => false
// "95a54e53" => false

// 思路：模拟
//		先用 e 分割几部分，然后分情况判断即可

import "strings"

func isNumber(s string) bool {
	// 去除首尾空白字符
	s = strings.TrimSpace(s)
	parts := strings.Split(s, "e")
	if len(parts) == 1 {
		// 如果不含 e ，则只用判断当前是否为数字即可
		isNum, _ := judge(parts[0])
		return isNum
	}
	if len(parts) == 2 {
		// 如果仅含有一个 e ，则底数必须是数字，指数必须是整数
		isFirstNum, _ := judge(parts[0])
		isSecondNum, isSecondInteger := judge(parts[1])
		return isFirstNum && isSecondNum && isSecondInteger
	}
	return false
}

func judge(s string) (isNumber, isInteger bool) {
	if len(s) == 0 {
		return false, false
	}

	// 去除正负号
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	if len(s) == 0 {
		return false, false
	}

	isInteger = true
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			// 如果是小数点，则不允许有多个
			if !isInteger {
				return false, false
			}
			isInteger = false
		} else if !('0' <= s[i] && s[i] <= '9') {
			// 如果既不是小数点，也不是数字，则不是数字
			return false, false
		}
	}
	// 小数必须含有数字
	return isInteger || len(s) > 1, isInteger
}
