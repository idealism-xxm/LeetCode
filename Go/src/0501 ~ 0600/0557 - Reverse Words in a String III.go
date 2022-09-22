// 链接：https://leetcode.com/problems/reverse-words-in-a-string-iii/
// 题意：给定一个字符串 s ，将其中的每个单词翻转。


// 数据限制：
//  1 <= s.length <= 5 * 10 ^ 4
//  s 仅含有 ASCII 字符
//  s 不含前导空格和末尾空格
//  s 中至少有一个单词
//  所有的单词都通过一个空格分隔


// 输入： s = "Let's take LeetCode contest"
// 输出： "s'teL ekat edoCteeL tsetnoc"

// 输入： s = "God Ding"
// 输出： "doG gniD"


// 思路： 模拟
//
//      先将 s 按照 ' ' 分隔成多个单词 words 。
//
//      再将 words 中的每个单词翻转。
//
//      最后用 ' ' 将 words 中的全部单词连起来即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 s 中全部 O(n) 个字符
//      空间复杂度：O(n)
//          1. 需要维护结果字符串中全部 O(n) 个字符


func reverseWords(s string) string {
    // 先将 s 按照 ' ' 分隔成多个单词
    words := strings.Split(s, " ")
    // 再将 words 中的每个单词翻转
    for i := range words {
        words[i] = reverse(words[i])
    }
    // 最后用 ' ' 将 words 中的全部单词连起来
    return strings.Join(words, " ")
}

func reverse(word string) string {
    runes := []rune(word)
    // 使用双指针翻转 rune 数组
    for l, r := 0, len(runes) - 1; l < r; l, r = l + 1, r - 1 {
        runes[l], runes[r] = runes[r], runes[l]
    }
    return string(runes)
}
