// 链接：https://leetcode.com/problems/reverse-words-in-a-string/
// 题意：给定一个字符串，按照单词进行翻转，并去除多余空格 ？

// 输入： "the sky is blue"
// 输出： "blue is sky the"

// 输入： "  hello world!  "
// 输出： "world! hello"

// 输入： "a good   example"
// 输出： "example good a"

// 思路： 模拟
//
//		先进行 trim ，然后以空格分开（用正则就不需要去除空串），
//		翻转后再用空格连起来即可
//

import (
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	// 去除两边空格
	s = strings.TrimSpace(s)
	// 以空格分开
	spaceReg, _ := regexp.Compile(" +")
	parts := spaceReg.Split(s, -1)
	// 进行翻转
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.Join(parts, " ")
}
