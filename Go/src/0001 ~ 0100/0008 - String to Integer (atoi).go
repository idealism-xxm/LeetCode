// 链接：https://leetcode.com/problems/string-to-integer-atoi/
// 题意：给定一个字符串，返回可以表示的整数（题目要求更像是 JavaScript 中的 parseInt 函数）
//		从第一个非空白字符开始，到第一个非数字字符止（前闭后开，且包含：+ -）
//		返回这一段字符表示的整数（超出 32 位整型，则返回相应的最值）。

// 数据限制：
//  0 <= s.length <= 200
//  s 仅由英文字母（大小写）、数字 (0-9) 、 ' ' 、 '+' 、 '-' 和 '.' 组成

// 输入：s = "42"
// 输出：42
// 解释：用于转换成数字的字符串为 s[0..2] = "42" ，
//      对应的整数是 42

// 输入：s = "   -42"
// 输出：-42
// 解释：用于转换成数字的字符串为 s[3..6] = "-42" ，
//      对应的整数是 -42

// 输入：s = "4193 with words"
// 输出：4193
// 解释：用于转换成数字的字符串为 s[0..4] = "4193" ，
//      对应的整数是 4193

// 输入：s = "-91283472332"
// 输出：-2147483648
// 解释：用于转换成数字的字符串为 s[0..12] = "-91283472332" ，
//      对应的整数是 -91283472332 ，
//      但这个整数不在 32 位整型范围内，
//      所以要返回最小值 -2147483648


// 思路： 模拟
//
//      1. 去除字符串空白前缀
//      2. 根据第一个字符判断符号 sign 和初始值 num
//      3. 遍历剩余的数字字符，直至遇到非数字字符或者字符串结束，
//          对每个字符 ch 执行 num = num * 10 + int(ch) ，
//          注意判断乘法和加法溢出时，要根据 sign 返回最大值/最小值
//      4. 最后返回带符号的结果 sign * num
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

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