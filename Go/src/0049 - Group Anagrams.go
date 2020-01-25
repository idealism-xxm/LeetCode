// 链接：https://leetcode.com/problems/group-anagrams/
// 题意：给定一个小写字符串数组，将字符串含有字母及个数相同的字符串分组，
//		返回所有分组列表？

// 输入：["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出：
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]

// 思路1：排序模拟即可
//		将每个字符串排序，将结果相同的字符串放入同一个列表即可
//		时间复杂度： O(n * k * log(k)) ，其中 k 是最长的字符串长度

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	strToStrArr := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		strToStrArr[sortedStr] = append(strToStrArr[sortedStr], str)
	}
	var result [][]string
	for _, strArr := range strToStrArr {
		result = append(result, strArr)
	}
	return result
}

func sortString(str string) string {
	strs := strings.Split(str, "")
	sort.Strings(strs)
	return strings.Join(strs, "")
}

// 思路2：计数模拟
//		看了官方题解后，发现还可以用计数的方式统计每个字母出现的次数
//		再组装成字符串当作键，由于使用了计数代替排序，复杂度降低了
//		时间复杂度： O(n * k) ，其中 k 是最长的字符串长度

import (
	"bytes"
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	strToStrArr := make(map[string][]string)
	for _, str := range strs {
		count := make([]int, 128)
		for i := range str {
			count[str[i]]++
		}
		var key bytes.Buffer
		for ch := 'a'; ch <= 'z'; ch++ {
			key.WriteString(strconv.Itoa(count[ch]))
			key.WriteByte('#')
		}
		keyStr := key.String()
		strToStrArr[keyStr] = append(strToStrArr[keyStr], str)
	}
	var result [][]string
	for _, strArr := range strToStrArr {
		result = append(result, strArr)
	}
	return result
}
