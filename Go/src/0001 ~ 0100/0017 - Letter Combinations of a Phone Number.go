// 链接：https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// 题意：给定一个数字串，每一位数字范围在 [2, 9] 内，
//      在九宫格键盘下一次按下相应的数字键，
//      求所有可能打出来的英文字符串？


// 数据限制：
//  0 <= digits.length <= 4
//  digits[i] 是一个 ['2', '9'] 范围内的数位


// 输入： digits = "23"
// 输出： ["ad","ae","af","bd","be","bf","cd","ce","cf"]

// 输入： digits = ""
// 输出： []

// 输入： digits = "2"
// 输出： ["a","b","c"]


// 思路1： 递归
//
//      我们使用 dfs(digits, index, cur, ans) 遍历收集所有可能的字符串，其中：
//          1. digits: 输入的数字串
//          2. index: 当前遍历到的下标
//          3. cur: 当前已遍历的数字串的一个可能的字符串
//          4. ans: 当前收集到的所有可能的字符串的列表
//
//      在 dfs 中，我们按照如下逻辑处理即可：
//          1. index == len(digits) ，则表明已经遍历完数字串，
//              此时 cur 就是一个可能的字符串，将其加入到 ans 中。
//          2. index != len(digits) ，则表明还需要继续遍历数字串，
//              遍历 digits[index] 对应的字母列表串中的字符 ch ，
//              将 cur[index] 设置为 ch ，然后递归调用 dfs 。
//
//
//      时间复杂度：O(4 ^ n)
//          1. 需要遍历全部可能的字符串，最差情况下所有的数字键都是 4 个字母的，
//              共有 O(4 ^ n) 个可能的字符串
//      空间复杂度：O(4 ^ n)
//          1. 需要收集全部可能的字符串，最差情况下所有的数字键都是 4 个字母的，
//              共有 O(4 ^ n) 个可能的字符串


// 定义每个数位对应的字母列表
var DIGIT_TO_LETTERS = [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}


func letterCombinations(digits string) []string {
	// 如果没有按任何数字键，则返回空列表
    if len(digits) == 0 {
		return nil
	}

	// ans 用于收集所有可能的字符串
	ans := make([]string, 0)
	// cur 表示当前按键下能形成的某个字符列表，初始化为空列表
	cur := make([]rune, len(digits))
	// 递归收集所有可能的字符串
	dfs(digits, 0, cur, &ans)

	return ans
}

func dfs(digits string, index int, cur []rune, ans *[]string) {
	// 如果已按下全部数字键，则 cur 就是一个可能的字符串，收集后返回
	if index == len(digits) {
		*ans = append(*ans, string(cur))
		return
	}

	// 遍历 digits[index] 下对应的的所有字母
	for _, ch := range DIGIT_TO_LETTERS[(digits[index] - '0')] {
		// 将当前字母放入 cur 中
		cur[index] = ch
		// 递归收集下一个字母
		dfs(digits, index + 1, cur, ans)
	}
}


// 思路2： 迭代
//
//      由于每个数字都会且仅会产生一个字母，
//      那么最终所有可能的字符串长度都是 len(digits) ，
//      即字符串的第 i 个字母是由 digits[i] 决定的。
//
//      如果知道 digits[:i] 对应的一个可能的字符串 cur ，
//      那么在 cur 后面分别加上 digits[i] 对应的字母 ch ，
//      即可得到 digits[:i + 1] 对应的一些可能的字符串。
//
//      我们可以运用递推的思路使用迭代处理，
//      让 ans[i] 表示 digits[:i] 对应的所有可能的字符串。
//
//      遍历 ans[i] 中的每个字符串 cur ，
//      在其后面加上 digits[i] 对应的每个字母 ch 组成新的字符串，
//      所有这样的新字符串就组成了 ans[i + 1] 。
//
//      初始化令 ans[0] = [""] ，
//      那么最后 ans[len(digits)] 就是所有可能的字符串，
//      实际处理时可以优化为滚动数组，降低空间复杂度。
//
//
//      时间复杂度：O(4 ^ n)
//          1. 需要遍历全部可能的字符串，最差情况下所有的数字键都是 4 个字母的，
//              共有 O(4 ^ n) 个可能的字符串
//      空间复杂度：O(4 ^ n)
//          1. 需要收集全部可能的字符串，最差情况下所有的数字键都是 4 个字母的，
//              共有 O(4 ^ n) 个可能的字符串


// 定义每个数位对应的字母列表
var DIGIT_TO_LETTERS = [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}


func letterCombinations(digits string) []string {
	// 如果没有按任何数字键，则返回空列表
    if len(digits) == 0 {
		return nil
	}

	// ans 用于收集所有可能的字符串，
	// 初始放入空串，方便后续迭代
	ans := []string {""}
	// 遍历数字串的数位 digit
	for _, digit := range digits {
		// 获取数位对应的字母列表
		letters := DIGIT_TO_LETTERS[digit - '0']
		// 定义下一个可能的字符串列表，容量为 ans 的长度乘以 letters 的长度，
		// 因为 ans 中的每个字符串都会被加上 letters 中的每个字母
		nxt := make([]string, 0, len(ans) * len(letters))
		// 遍历 ans 中的每个字符串
		for _, cur := range ans {
			// 遍历 letters 中的每个字母
			for _, ch := range letters {
				// 拼接 digit 对应的所有字母，然后放入 nxt 中
				nxt = append(nxt, cur + string(ch))
			}
		}

		ans = nxt
	}

	return ans
}
