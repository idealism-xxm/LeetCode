// 链接：https://leetcode.com/problems/reverse-prefix-of-word/
// 题意：给定一个字符串 word 和一个字符 ch ，翻转第一个 ch 及其之前的子串。


// 数据限制：
//   1 <= word.length <= 250
//   word 由英文小写字母组成
//   ch 是英文小写字母


// 输入： word = "abcdefd", ch = "d"
// 输出： "dcbaefd"
// 解释： 翻转子串 "abcd" 后得到 "dcba" ，新字符串为 "dcbaefd"

// 输入： word = "xyxzxe", ch = "z"
// 输出： "zxyxxe"
// 解释： 翻转子串 "xyxz" 后得到 "zxyx" ，新字符串为 "zxyxxe"

// 输入： word = "abcd", ch = "z"
// 输出： "abcd"
// 解释： 没有字符 "z" ，不需要翻转


// 思路： 模拟
//
//       按照题意找到第一个 ch 后，翻转前缀串即可。
//
//
//       时间复杂度： O(n)
//           1. 需要遍历 word 中全部 O(n) 个字母
//       空间复杂度： O(n)
//           1. 需要生成长度为 O(n) 的结果串


func reversePrefix(word string, ch byte) string {
    for i := range word {
        // 如果存在 ch ，则翻转前缀串
        if word[i] == ch {
            return reverse(word[:i+1]) + word[i + 1:]
        }
    }
    // 此时不存在 ch ，直接返回原字符串
    return word
}

func reverse(str string) string {
    runes := []rune(str)
    for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
