// 链接：https://leetcode.com/problems/excel-sheet-column-number/
// 题意：给定一个 Excel 的标题，返回其对应的列号？
//    	A -> 1
//    	B -> 2
//    	C -> 3
//    	...
//    	Z -> 26
//    	AA -> 27
//    	AB -> 28
//    	...

// 输入： "A"
// 输出： 1

// 输入： "AB"
// 输出： 28

// 输入： "ZY"
// 输出： 701

// 思路： 模拟
//
//		0168 的反向操作，仍然是类似进制转换，模拟即可
//

func titleToNumber(s string) int {
	result := 0
	for _, ch := range s {
		// 由于 ch - 'A' 是 base 0 的，所以要 + 1 转换成 base 1 的
		result = result * 26 + int(ch - 'A' + 1)
	}
	return result
}
