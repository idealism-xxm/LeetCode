// 链接：https://leetcode.com/problems/valid-palindrome/
// 题意：给定一个字符串，求去掉所有非字母数字并转换成小写后是否为回文？

// 输入： "A man, a plan, a canal: Panama"
// 输出： true

// 输入： "race a car"
// 输出： false

// 思路： 双指针
//		为了只遍历一次，不去除无效字符，不提前转换大小写
//		可以利用普通判断回文串的方法，在加上本题的一些特判即可：
//		前后指针分别找到第一个字母数字，然后都转成小写，如果相同则继续处理，如果不同则直接返回 false
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

import "unicode"

func isPalindrome(s string) bool {
	// 前后指针分别指向开始和结尾
	for l, r := 0, len(s) - 1; l < r; l, r = l + 1, r - 1 {
		// 前指针找到第一个字母数字
		for ; l < r && !isAlphanumeric(s[l]); l++ {}
		// 后指针找到第一个字母数字
		for ; l < r && !isAlphanumeric(s[r]); r-- {}
		// 如果还未相遇，且转换成小写字母后不相同，则直接返回 false
		if l < r && toLower(s[l]) != toLower(s[r]) {
			return false
		}
	}
	return true
}

func isAlphanumeric(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || unicode.IsDigit(rune(ch))
}

func toLower(ch byte) byte {
	if 'A' <= ch && ch <= 'Z' {
		return ch - 'A' + 'a'
	}
	return ch
}