// 链接：https://leetcode.com/problems/palindrome-number/
// 题意：给定一个整数，判断其字符串表示是否是回文串

// 输入：121
// 输出：true

// 输入：-121
// 输出：false

// 思路：负数不是回文；非负数先获取其反转整数（个位 -> 最高位，十位 -> 次高位，依此类推），若两者相等则是回文，否则不是
func isPalindrome(x int) bool {
	// 负数肯定不是回文
	if x < 0 {
		return false
	}

	reverse := 0 // 非负整数 x 的反转整数
	tmp := x
	for ; tmp > 0;  {
		reverse = reverse * 10 + tmp % 10
		tmp /= 10
	}
	
	return reverse == x // 反转后等于原数，则是回文
}