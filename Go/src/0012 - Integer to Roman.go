// 链接：https://leetcode.com/problems/integer-to-roman/
// 题意：给定 [1, 3999] 范围内的正整数，将其转换成罗马数字
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// 输入：1994
// 输出："MCMXCIV"

// 思路：模拟即可
// 		可以发现每一位的数字只会用到3个罗马数字，分别表示 1,5,10
// 		所以可以构建通用的位处理逻辑，循环处理即可，增加数字位数时，添加相应的罗马数字即可

import (
	"bytes"
)

func intToRoman(num int) string {
	romanChar := []byte {'I', 'V', 'X', 'L', 'C', 'D', 'M', ' ', ' '} // 由于千位保证最大为 3，此处不必有另外两个字符

	result := ""
	start := 0
	for ; num > 0; num /= 10 {
		result = buildDigitRoman(romanChar[start], romanChar[start + 1], romanChar[start + 2], num % 10) + result
		start += 2
	}

	return result
}

// 构建当前位的罗马数字
func buildDigitRoman(one, five, ten byte, digit int) string {
	var result bytes.Buffer
	switch digit {
	case 1:
		writeBytes(&result, one)
	case 2:
		writeBytes(&result, one, one)
	case 3:
		writeBytes(&result, one, one, one)
	case 4:
		writeBytes(&result, one, five)
	case 5:
		writeBytes(&result, five)
	case 6:
		writeBytes(&result, five, one)
	case 7:
		writeBytes(&result, five, one, one)
	case 8:
		writeBytes(&result, five, one, one, one)
	case 9:
		writeBytes(&result, one, ten)
	}

	return result.String()
}

// buffer 中写入多个 字节
func writeBytes(buffer *bytes.Buffer, bytes ...byte) {
	for _, cur := range bytes {
		buffer.WriteByte(cur)
	}
}