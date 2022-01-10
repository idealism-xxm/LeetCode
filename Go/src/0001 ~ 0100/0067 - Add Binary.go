// 链接：https://leetcode.com/problems/add-binary/
// 题意：给定两个 01 串，求他们的二进制下的和？

// 输入：a = "11", b = "1"
// 输出："100"

// 输入：a = "1010", b = "1011"
// 输出："10101"

// 思路：模拟即可
//		从个位开始计算，考虑进位即可，注意最高位进位的情况

import "strconv"

func addBinary(a string, b string) string {
	carry := 0
	result := ""
	for ai, bi := len(a) - 1, len(b) - 1; ai >= 0 || bi >= 0; ai, bi = ai - 1, bi - 1 {
		current := carry
		if ai >= 0 && a[ai] == '1' {
			current++
		}
		if bi >= 0 && b[bi] == '1' {
			current++
		}
		carry = current >> 1
		result = strconv.Itoa(current & 1) + result
	}

	if carry == 1 {
		result = "1" +  result
	}
	return result
}
