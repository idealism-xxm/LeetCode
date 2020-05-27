// 链接：https://leetcode.com/problems/reverse-words-in-a-string-ii/
// 题意：给定一个字节数组表示的字符串，将其按单词原地翻转，要求空间复杂度为 O(1)？

// 输入： ["t","h","e"," ","s","k","y"," ","i","s"," ","b","l","u","e"]
// 输出： ["b","l","u","e"," ","i","s"," ","s","k","y"," ","t","h","e"]

// 思路： 遍历
//
//		先翻转整个数组，然后再翻转每个单词即可
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func reverseWords(s []byte)  {
	length := len(s)
	// 先翻转整个数组
	reverseRange(s, 0, length - 1)

	// 然后找到每一个空格前的担起进行翻转即可
	l := 0
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			reverseRange(s, l, i - 1)
			// 下一个单词起点是 i + 1
			l = i + 1
		}
	}
	// 最后一个单词也需要翻转
	reverseRange(s, l, length - 1)
}

func reverseRange(s []byte, l, r int) {
	for ; l < r; l, r = l + 1, r - 1 {
		s[l], s[r] = s[r], s[l]
	}
}
