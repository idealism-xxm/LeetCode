// 链接：https://leetcode.com/problems/length-of-last-word/
// 题意：给定一个由大小写字母和空格组成的字符串，求最后一个单词的长度？

// 输入："Hello World"
// 输出：5

// 思路：模拟
//		从后往前计数到第一个空格即可
//      时间复杂度： O(n) ，空间复杂度： O(1)

func lengthOfLastWord(s string) int {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			// 空滤尾部空格
			if result != 0 {
				break
			}
		} else {
			result++
		}

	}
	return result
}
