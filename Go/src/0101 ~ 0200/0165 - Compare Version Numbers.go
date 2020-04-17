// 链接：https://leetcode.com/problems/compare-version-numbers/
// 题意：给定两个版本号（只含有数字和 .），比较这两个版本号的大小 ？

// 输入： version1 = "0.1", version2 = "1.1"
// 输出： -1

// 输入： version1 = "1.0.1", version2 = "1"
// 输出： 1

// 输入： version1 = "7.5.2.4", version2 = "7.5.3"
// 输出： -1

// 输入： version1 = "1.01", version2 = "1.001"
// 输出： 0
// 解释： 忽略前导零， "01" 和 "001" 都代表数字 0

// 输入： version1 = "1.0", version2 = "1.0.0"
// 输出： 0
// 解释： 第一个版本号的第三部分没有，则默认为 0

// 思路： 模拟
//
//		1. 按照 . 进行划分
//		2. 短的补 "0" ，让长度相等
//		3. 每一部分转成数字对比即可
//
//		当然也可以不进行划分，同时扫描两个字符串，遇到 . 或者字符串到末尾时停止，
//		计算当前对应的数字对比即可
//

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	// 1. 按照 . 进行划分
	parts1 := strings.Split(version1, ".")
	parts2 := strings.Split(version2, ".")

	// 2. 短的补 "0" ，让长度相等
	length1, length2 := len(parts1), len(parts2)
	if length1 < length2 {
		for i := length1; i < length2; i++ {
			parts1 = append(parts1, "0")
		}
	} else if length1 > length2 {
		for i := length2; i < length1; i++ {
			parts2 = append(parts2, "0")
		}
	}

	// 3. 每一部分转成数字对比即可
	length := max(length1, length2)
	for i := 0; i < length; i++ {
		num1, _ := strconv.Atoi(parts1[i])
		num2, _ := strconv.Atoi(parts2[i])
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}
	// 所有数字都相同，则版本号相同
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
