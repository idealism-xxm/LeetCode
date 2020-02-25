// 链接：https://leetcode.com/problems/text-justification/
// 题意：给定单词数组，将其进行文本对齐，
//		最后一行左对齐，其余行两端对齐，
//		空格不能均分的情况下，优先向左补充。

// 输入：
//  words = ["This", "is", "an", "example", "of", "text", "justification."]
//  maxWidth = 16
// 输出：
// [
//   "This    is    an",
//   "example  of text",
//   "justification.  "
// ]

// 输入：
// words = ["What","must","be","acknowledgment","shall","be"]
// maxWidth = 16
// 输出：
// [
//   "What   must   be",
//   "acknowledgment  ",
//   "shall be        "
// ]

// 输入：
// words = ["Science","is","what","we","understand","well","enough","to","explain",
//         "to","a","computer.","Art","is","everything","else","we","do"]
// maxWidth = 20
// 输出：
// [
//   "Science  is  what we",
//   "understand      well",
//   "enough to explain to",
//   "a  computer.  Art is",
//   "everything  else  we",
//   "do                  "
// ]

// 思路：模拟即可
//		先贪心把尽可能多的单词放入（注意两个单词间多加一个空格）
//		然后计算每个空白处的空格数量，填充即可

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	length := len(words)
	var result []string
	for i := 0; i < length; {
		// 计算本行可填入的单词数
		end, width := i+1, len(words[i])
		for ; end < length && width+len(words[end])+1 <= maxWidth; end++ {
			width += len(words[end]) + 1
		}

		row := words[i]
		if end == length {
			// 最后一行左对齐
			for j := i + 1; j < end; j++ {
				row += " " + words[j]
			}
		} else {
			// 两端对齐
			num := end - i - 1 // 空白数
			if num != 0 {
				// 计算每个空白需要补充多少空格
				remain := maxWidth - width
				average, extra := remain/num, remain%num
				for j := i + 1; j < end; j, extra = j+1, extra-1 {
					count := 1 + average
					if extra > 0 {
						count++
					}
					row += strings.Repeat(" ", count) + words[j]
				}
			}
		}
		// 最后一行和只有一个单词的行需要补齐
		row += strings.Repeat(" ", maxWidth-len(row))
		result = append(result, row)

		i = end
	}
	return result
}
