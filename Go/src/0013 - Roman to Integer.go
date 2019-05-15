// 链接：https://leetcode.com/problems/roman-to-integer/
// 题意：给定 [1, 3999] 范围内的罗马数字，将其转换成正整数
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// 输入："MCMXCIV"
// 输出：1994

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

// 思路2：模拟，挖掘更深的关系（通过后，看见有更好更简洁的解法）
//		眼光不要局限于每一位数字，继续观察 思路1 中发现的规律
// X -> X、XX、XXX、XL、XC 	（前三个加法，后两个减法）
// L -> L、LX、LXX、LXXX	（全部是加法）
// 若 当前字符是减去时，必定是 当前字符的数字比下一个字符的小
// 若 当前字符时加上时，必定是 当前字符的数字比下一个字符的小

var romanCharNum = map[byte]int {
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	length := len(s)
	result := romanCharNum[s[length - 1]] // 最后一个罗马字符一定是加上去的
	for i := length - 2; i >= 0; i-- {
		if romanCharNum[s[i]] < romanCharNum[s[i + 1]] {
			result -= romanCharNum[s[i]]
		} else {
			result += romanCharNum[s[i]]
		}
	}

	return result
}