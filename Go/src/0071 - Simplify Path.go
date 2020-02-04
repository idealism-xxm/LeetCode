// 链接：https://leetcode.com/problems/simplify-path/
// 题意：给定一个文件的绝对路径，将其简化为规范路径？

// 输入："/home/"
// 输出："/home"

// 输入："/../"
// 输出："/"

// 输入："/home//foo/"
// 输出："/home/foo"

// 输入："/a/./b/../../c/"
// 输出："/c"

// 输入："/a/../../b/../c//.//"
// 输出："/c"

// 输入："/a//b////c/d//././/.."
// 输出："/a/b/c"

// 思路：模拟
//		用 '/' 分成多个部分，然后模拟简化即可

import "strings"

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	length := len(parts)
	result := make([]string, length)
	result[0] = ""
	last := 0
	for i := 1; i < length; i++ {
		if parts[i] == "." || parts[i] == "" {
			continue
		}
		if parts[i] == ".." {
			// 根目录到父目录还是根目录
			if last > 0 {
				last--
			}
		} else {
			last++
			result[last] = parts[i]
		}
	}
	if last == 0 {
		return "/"
	}
	return strings.Join(result[:last + 1], "/")
}
