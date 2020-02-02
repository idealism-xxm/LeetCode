// 链接：https://leetcode.com/problems/plus-one/
// 题意：给定一个非空整型数组，表示一个非负无前导零待整数
//		对这个整数加一，返回结果数组？

// 输入：[1,2,3]
// 输出：[1,2,4]

// 输入：[4,3,2,1]
// 输出：[4,3,2,2]

// 思路：模拟即可
//		从个位开始计算，考虑进位即可，注意最高位进位待情况

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		result := digits[i] + carry
		carry = result / 10
		digits[i] = result % 10
	}
	if carry == 0 {
		return digits
	}
	return append([]int{carry}, digits...)
}
