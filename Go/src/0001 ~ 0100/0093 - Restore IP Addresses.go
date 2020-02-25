// 链接：https://leetcode.com/problems/restore-ip-addresses/
// 题意：给定一个只有数字的 ip 地址字符串，返回所有可能的 ip 地址？

// 输入："25525511135"
// 输出：["255.255.11.135", "255.255.111.35"]

// 思路：递归
//		每一层至枚举当前部分可用的数字，注意 不能含有前导零，且数字必须小于等于 255

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return nil
	}
	return dfs(4, s, "")
}

// 通过 s 生成 ip 地址的后 n 部分
func dfs(n int, s string, current string) []string {
	// 剩余部分的字符串为最后一部分
	if n == 1 {
		// 含有前导零，则不是合法的分割
		if s != "0" && s[0] == '0' {
			return nil
		}
		num, _ := strconv.Atoi(s)
		// 当前部分的数字必须大于 255 则不合法
		if num > 255 {
			return nil
		}
		return []string{current + s}
	}

	// 第一位是 0 ，则当前部分只能为 0
	if s[0] == '0' {
		return dfs(n - 1, s[1:], current + "0.")
	}

	var result []string
	// 当前部分的长度最长为 min(3, len(s) - n + 1) ，保证后续部分有足够可用字符
	for i := min(3, len(s) - n + 1); i > 0; i-- {
		num, _ := strconv.Atoi(s[:i])
		// 当前部分的数字必须小于等于 255 才合法
		if num <= 255 {
			result = append(result, dfs(n - 1, s[i:], current + s[:i] + ".")...)
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}