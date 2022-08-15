// 链接：https://leetcode.com/problems/roman-to-integer/
// 题意：给定 [1, 3999] 范围内的罗马数字，将其转换成正整数。
//         Symbol       Value
//           I            1
//           V            5
//           X            10
//           L            50
//           C            100
//           D            500
//           M            1000
//
//        罗马数字通常是按照符号大小从左往右写的，
//        一般都是直接将所有符号对应的值加起来即可，除了以下六种情况：
//            1. I 可以放在 V(5) 和 X(10) 前面，分别组成 4 和 9 。
//            2. X 可以放在 L(50) 和 C(100) 前面，分别组成 40 和 90 。
//            3. C 可以放在 D(500) 和 M(1000) 前面，分别组成 400 和 900 。


// 数据限制：
//  1 <= s.length <= 15
//  s 仅由以下字符组成： 'I', 'V', 'X', 'L', 'C', 'D', 'M'
//  给定的罗马数字必定在 [1, 3999] 范围内


// 输入： s = "III"
// 输出： 58
// 解释： III = 3

// 输入： s = "LVIII"
// 输出： III = 3
// 解释： L = 50, V= 5, III = 3

// 输入： "MCMXCIV"
// 输出： 1994
// 解释： M = 1000, CM = 900, XC = 90, IV = 4


// 思路1：模拟即可
// 枚举每一位数字所有的情况，可以发现（假设当前位第一个罗马数字是 X 或 L）只有以下情况：
// X -> X、XX、XXX、XL、XC
// L -> L、LX、LXX、LXXX
// 所以每次计算时，枚举每个字符判断即可

type Roman struct {
	Char byte
	Val int
}

var romans = []Roman {
	{'I', 1},
	{'V', 5},
	{'X', 10},
	{'L', 50},
	{'C', 100},
	{'D', 500},
	{'M', 1000},
	{' ', 0}, // 由于千位保证最大为 3，此处不必有另外两个字符和数字
	{' ', 0},
}

// 存储每个罗马数字字符对应的下标
var romanCharIndex = make(map[byte]int)

func romanToInt(s string) int {
	// 初始化
	for i, roman := range romans {
		romanCharIndex[roman.Char] = i
	}

	length := len(s)
	result := 0
	for i := 0; i < length; {
		index := romanCharIndex[s[i]] // 5... 都是奇数，相应的 1... 都比其小 1（都是偶数）
		if (index & 1) == 1 { // 如果是 5...，则匹配类似 L、LX、LXX、LXXX 以下几种情况（都是加法）
			result += romans[index].Val
			index ^= 1 // 奇数减一用异或即可，则当前 index 指向 X
			for i += 1; i < length && s[i] == romans[index].Char; i++ { // 遍历完剩下的 X
				result += romans[index].Val
			}
		} else { // 此时都是 1...，则匹配类似 X、XX、XXX、XL、XC（最后两种是减法，需要特判）
			if i + 1 < length && (s[i + 1] == romans[index + 1].Char || s[i + 1] == romans[index + 2].Char) {
				result += romans[romanCharIndex[s[i + 1]]].Val - romans[index].Val
				i += 2
			} else {
				result += romans[index].Val
				for i += 1; i < length && s[i] == romans[index].Char; i++ { // 遍历完剩下的 X
					result += romans[index].Val
				}
			}
		}
	}

	return result
}


// 思路2：Map
//
//      通常情况下，我们只需要将所有符号对应的值加起来即可，
//      但要同时考虑六种特殊情况。   
//
//      可以发现这六种特殊情况都是值小的符号在值大的符号之前，
//      此时要从最终结果中减去他们对应的值。
//
//      其他情况，直接将符号对应的值加到最终结果中即可。
//
//      那么可得以下计算规则（设 value(ch) 表示符号 ch 对应的值）：
//          1. value(s[i]) < value(s[i + 1]): 则需要从结果中减去 value(s[i])
//          2. value(s[i]) >= value(s[i + 1]): 则需要给结果加上 value(s[i])
//
//      最后一个符号后面没有其他符号，所以必定是直接加到结果中的。
//
//      为了方便处理边界情况，我们可以直接倒序从倒数第 2 个符号开始向前遍历即可。
//
//
//		设字符集大小为 C 。
//
//      时间复杂度：O(n)
//          1. 需要遍历 s 中全部 O(n) 个字符
//      空间复杂度：O(C)
//          1. 需要维护全部 O(C) 个不同字符对应的值


// chToNum[ch] 表示符号 ch 对应的值
var chToNum = map[byte]int {
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	// ans 直接初始化为最后一个符号对应的值，
	// 因为最后一个符号后面没有其他符号，所以必定是直接加到结果中的
	ans := chToNum[s[len(s) - 1]]
	// 倒序从倒数第 2 个符号开始向前遍历
	for i := len(s) - 2; i >= 0; i-- {
		if chToNum[s[i]] < chToNum[s[i + 1]] {
			// 如果 s[i] 的值小于 s[i + 1] 的值，则需要从 ans 中减去 s[i] 的值
			ans -= chToNum[s[i]]
		} else {
			// 否则，需要给 ans 加上 s[i] 的值
			ans += chToNum[s[i]]
		}
	}

	return ans
}
