// 链接：https://leetcode.com/problems/excel-sheet-column-title/
// 题意：给定一个正整数，返回其在 Excel 下对应的标题？
//    	1 -> A
//    	2 -> B
//    	3 -> C
//    	...
//    	26 -> Z
//    	27 -> AA
//    	28 -> AB
//    	...

// 输入： 1
// 输出： "A"

// 输入： 28
// 输出： "AB"

// 输入： 701
// 输出： "ZY"

// 思路： 模拟
//
//		其实就是一种进制转换，从最低位模拟即可
//

import "bytes"

func convertToTitle(n int) string {
	var result bytes.Buffer
	for n > 0 {
		// n-- ：转换成从 0 开始的
		n--
		remain := n % 26
		n = n / 26
		result.WriteByte(byte('A' + remain))
	}
	return reverse(result.String())
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
