// 链接：https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// 题意：给定一个数字串，每一位数字范围在 [2, 9] 内，在九宫格键盘下一次按下相应的数字键，求所有可能打出来的英文字符串

// 输入："23"
// 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]

// 思路：暴力递归即可

var digitLetters = []string {
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" { // 如果数字串使空串，则不会产生任何结果
		return []string {}
	}
	
	return combineLetters(digits, "")
}

func combineLetters(digits, cur string) []string {
	if digits == "" { // 如果串为空，则本次遍历的字符已经全部完成，直接返回
		return []string {cur}
	}

	var result []string // 本次递归内产生的结果
	num := digits[0] - '0' // 真实数字
	remainDigits := digits[1:] // 剩下的数字串

	for i := len(digitLetters[num]) - 1; i >= 0; i-- { // 遍历真实数字所有可能的字符，递归调用
		// 由于 先返回的使字典序最大的，所以要将最新的结果放在前面，保证字典序升序
		result = append(combineLetters(remainDigits, cur + string(digitLetters[num][i])), result...)
	}

	return result
}