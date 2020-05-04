// 链接：https://leetcode.com/problems/largest-number/
// 题意：给定 DNA 单链序列，求连续 10 个核苷酸序列出现多次的是哪些？

// 输入： s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
// 输出： ["AAAAACCCCC", "CCCCCAAAAA"]

// 思路1： map + 计数
//
//		将所有连续 10 个核苷酸序列计数，然后收集出现次数大于 1 次的即可
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)

func findRepeatedDnaSequences(s string) []string {
	// 1. 统计每个序列出现的次数
	sequenceCount := make(map[string]int)
	for end := 10; end <= len(s); end++ {
		sequenceCount[s[end - 10: end]]++
	}

	// 2. 收集出现次数大于 1 次的为答案
	var result []string
	for sequence := range sequenceCount {
		if sequenceCount[sequence] > 1 {
			result = append(result, sequence)
		}
	}
	return result
}

// 思路2： Rabin-Karp
//
//		前面一种方法每次都需要取子串，常数为 10 ，
//		看了题解发现 Rabin-Karp 不用取子串，
//		并且能直接从上一次的 hash 计算出本次 hash ，常数变小了，
//		时间复杂度常数不会在子串长度变长的情况下增加
//
//		我们使用 4 进制进行 hash ，将字符串转换成对应的 4 进制数字，
//		再转回 10 进制进行计数处理，然后将出现次数大于 1 的转回核苷酸序列即可
//		（由于不支持大数， Go 中最长能支持到的子串长度位 32 ）
//		当知道前一个 hash 时，右移窗口，需要去除最前面的字符：减去其对应的数字 * 对应的权重
//		然后再乘以 4 ，最后再加上新出现的字符对应的数字即可
//
//		由于是 4 进制，所以很容易就能转换成位运算操作，每一位数字占 2 位二进制位：
//		转换 hash 时，最前面两位置 0 ，后 18 位保留，然后左移动 2 位，
//		将新字符对应对数字放在最后两位即可
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return nil
	}

	// 每个字符对应的 4 进制数字
	charToNum := map[byte]int {
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}
	// 每个 4 进制数字对应的字符
	chars := "ACGT"

	// 1. 统计每个序列对应的 hash 出现的次数
	// 1.1 计算前 10 位的 hash 并计数
	hashCount := make(map[int]int)
	hash := 0
	for i := 0; i < 10; i++ {
		hash = (hash << 2) + charToNum[s[i]]
	}
	hashCount[hash]++
	// 1.2 移动窗口，计算后续的 hash 并计数
	// 最前面的字符的权重
	weight := 1 << (9 << 1)
	for end := 11; end <= len(s); end++ {
		hash = ((hash - charToNum[s[end - 11]] * weight) << 2) + charToNum[s[end - 1]]
		hashCount[hash]++
	}

	// 2. 收集出现次数大于 1 次的为答案
	var result []string
	for h := range hashCount {
		if hashCount[h] > 1 {
			// 还原 hash 对应的子串
			sequence := ""
			for i := 0; i < 10; i++ {
				remain := h & 3
				sequence = chars[remain:remain + 1] + sequence
				h >>= 2
			}
			result = append(result, sequence)
		}
	}
	return result
}
