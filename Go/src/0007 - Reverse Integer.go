// 链接：https://leetcode.com/problems/reverse-integer/
// 题意：给定一个整数，保持符号不变，翻转每一位数字，即个位变为最高位，十位变为次高位，依此类推。同时要保证不含前导零。
// 输入：-120
// 输出：-21

// 思路：去掉符号并记录，初始化结果为 result = 0，每次取出最低位的数字 digit，则 result = result * 10 + digit
// 刚开始没仔细看提示，以为是保证答案在 32位 整数内，结果是需要特判（觉得这种影响数据结果的信息还是要放在题目描述中）
func reverse(x int) int {
	sign := 1 // 默认是整数
	// 如果是负数，则对 x 取绝对值，记录负号
	if x < 0 {
		x = -x
		sign = -1
	}

	result := 0 // 初始化结果为 0
	for ; x > 0;  {
		digit := x % 10 // 取出最低位的数字
		result = result * 10 + digit

		x = x / 10 // 准备处理下一位数
	}
	result = sign * result

	// 超出 32 位整型范围，需要返回 0
	if result < -(1 << 31) || result > (1 << 31) - 1 {
		return 0
	}
	return result
}