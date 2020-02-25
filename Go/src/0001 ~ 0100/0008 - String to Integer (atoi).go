// 链接：https://leetcode.com/problems/string-to-integer-atoi/
// 题意：给定一个字符串，返回可以表示的整数（题目要求更像是 javascript 中的 parseInt 函数）
//		从第一个非空白字符开始，到第一个非数字字符止（前闭后开，且包含：+ -）
//		则其输出表示的整数（超出32位整型，则返回相应的最值）

// 输入：   -42
// 输出：-42

// 输入：4193 with words
// 输出：4193

// 输入：words and 987
// 输出：0

// 输入：-91283472332
// 输出：-2147483648

// 思路：与上一题类似，就是特判比较多
// 先截取符合条件的子串，去掉符号并记录
// 初始化结果为 result = 0，按顺序取每一位的数字 digit，则 result = result * 10 + digit

import "math"

func myAtoi(str string) int {
	length := len(str)

	sign := 1 // 默认位正数
	start, end := -1, -1 // 获取满足题意的子串下标

	i := 0
	for ; i < length && str[i] == ' '; i += 1 {} // 过滤空白字符
	// 如果第一个非空白字符是 正负号
	if i < length && str[i] == '+' {
		i += 1
	} else if i < length && str[i] == '-' {
		sign = -1
		i += 1
	}
	for ; i < length && str[i] == '0'; i += 1 {} // 过滤前导 0

	start = i // 数字串起始下标为 i
	for ; i < length && '0' <= str[i] && str[i] <= '9'; i += 1 {} // 过滤所有的数字串
	end = i // 数字串的起始下标是第一个非数字字符（前闭后开）

	// 长度超过 32位整型最长长度，直接返回最值
	subLength := end - start
	if subLength > 10 {
		if sign == -1 {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}
	// 计算数字值（依旧有可能超出 32位整型范围）
	result := 0 // 此处也兼容 第一个不是非数字字符
	for i = start; i < end; i += 1 {
		result = result * 10 + int(str[i] - '0')
	}
	result *= sign

	// 超出 32 位整型范围，需要返回 最值
	if result < math.MinInt32 {
		return math.MinInt32
	}
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	return result
}