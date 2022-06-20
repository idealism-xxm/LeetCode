// 链接：https://leetcode.com/problems/rearrange-characters-to-make-target-string/
// 题意：给定两个字符串 s 和 target ，可以重新排列 s 中的字母形成一个新的字符串，
//       求 target 在新字符串中，最多能出现几次？


// 数据限制：
//  1 <= s.length <= 100
//  1 <= target.length <= 10
//  s 和 target 仅含有英文小写字母


// 输入： s = "ilovecodingonleetcode", target = "code"
// 输出： 2
// 解释： "code" 第一次出现可以使用 s 中下标为 4, 5, 6, 7 的字母，
//        "code" 第二次出现可以使用 s 中下标为 17, 18, 19, 20 的字母。

// 输入： s = "abcba", target = "abc"
// 输出： 1
// 解释： "abc" 第一次出现可以使用 s 中下标为 0, 1, 2 的字母。
//        注意，我们不能重复使用字母 'c'  。

// 输入： s = "abbaccaddaeea", target = "aaaaa"
// 输出： 1
// 解释： "aaaaa" 第一次出现可以使用 s 中下标为 0, 3, 6, 9, 12 的字母。


// 思路： Map
//
//      分别统计 s 和 target 中所有字符出现的次数到 total 和 required 中。
//
//      然后枚举 required 中的每种字符 ch 及出现次数 cnt ，
//      计算 s 中的字符 ch 最多能用于 target 的数量 total[ch] / cnt 。
//
//      所有这些数量的最小值，就是就是 target 最多能出现的次数。
//
//
//      设字符集大小为 C 。
//
//      时间复杂度：O(n + m)
//          1. 需要遍历 s 中全部 O(n) 个字符
//          2. 需要遍历 target 全部 O(m) 个字符
//      空间复杂度：O(C)
//          1. 需要维护 s 中全部 O(C) 种字符的出现次数
//          2. 需要维护 target 中全部 (C) 种字符出现的次数


func rearrangeCharacters(s string, target string) int {
    // 统计 s 和 target 中所有字符出现的次数
    total := make(map[rune]int)
    for _, ch := range s {
        total[ch] += 1
    }
    required := make(map[rune]int)
    for _, ch := range target {
        required[ch] += 1
    }
    // ans 维护 target 最多能出现的次数，
    // 初始化为 len(s) ，因为 target 最多只能出现 len(s) 次
    ans := len(s)
    // 遍历 target 中的每种字符及其出现次数
    for ch, cnt := range required {
        // 计算 s 中的字符 ch 最多能用于 target 的数量，并更新 ans 的最小值
        ans = min(ans, total[ch] / cnt)
    }
    return ans     
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
