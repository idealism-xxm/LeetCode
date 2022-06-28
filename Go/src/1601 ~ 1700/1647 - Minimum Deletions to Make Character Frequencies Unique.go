// 链接：https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
// 题意：给定一个字符串 s ，求最少删除多少个字符后，字符串 s 中所有不同的字符出现的次数各不相同。


// 数据限制：
//  1 <= s.length <= 10 ^ 5
//  s 仅含有英文小写字母


// 输入： s = "aab"
// 输出： 0
// 解释： 'a' 和 'b' 的出现次数已经各不相同。

// 输入： s = "aaabbbcc"
// 输出： 2
// 解释： 删除两个 'b' 后， 'a', 'b', 'c' 的出现次数就各不相同。

// 输入： s = "ceabaacb"
// 输出： 2
// 解释： 删除全部两个 'c' 后， 'a', 'b', 'e' 的出现次数各不相同。
//       注意：我们只关心最后仍在 s 中的字符。


// 思路： 排序 + 贪心
//
//      为了保证删除字符数最少，我们需要贪心地保证出现次数最多的字符不被删（假设出现次数为 max_cnt ），
//      那么我们就能确定在剩余字符中，出现次数最多的那个字符最多只能有 max_cnt - 1 个字符。
//
//      如此反复就可以按照初始字符的出现次数，从大到小确定每一种字符应该删除的字符数。
//
//      所以可以维护字符最大允许出现的次数 maxCntAllowed ，
//      初始化为 len(s) ，表示初始出现最多的字符最多只能有 len(s) 个字符。
//
//      统计每种字符的出现次数，对出现次数按照降序排序，然后遍历每一个出现次数 cnt ，
//      如果 cnt > maxCntAllowed ，则说明要删除 maxCntAllowed - cnt 个该字符。
//      即令 ans += maxCntAllowed - cnt; cnt = maxCntAllowed 。
//
//      然后需要更新下一个字符最大允许出现的次数 = 当前字符出现次数 - 1 ，
//      同时要保证一个字符最多只能被删光，即 maxCntAllowed 最小为 0 。
//
//
//      设字符集大小为 C 。
//
//      时间复杂度：O(n + ClogC)
//          1. 需要遍历 s 中全部 O(n) 个字符
//          2. 需要对全部 O(C) 种字符进行排序，时间复杂度为 O(ClogC)
//      空间复杂度：O(C)
//          1. 需要维护全部 O(C) 种字符的次数


import "sort"


func minDeletions(s string) int {
    // 统计每一种字符出现的次数
    cnts := make([]int, 26)
    for _, ch := range s {
        cnts[ch - 'a'] += 1
    }
    // 我们只关心字符的出现次数，不关心是哪种字符，
    // 所以直接对 cnts 按出现次数降序排序即可
    sort.Slice(cnts, func(i, j int) bool {return cnts[i] > cnts[j]})

    // ans 维护需要删除的字符数量
    ans := 0
    // maxCntAllowed 表示当前字符最大允许出现的次数
    maxCntAllowed := len(s)
    // 遍历所有出现过的字符（ cnt > 0 ）的出现次数
    for _, cnt := range cnts {
        if cnt == 0 {
            break
        }

        // 如果 cnt 超过了 maxCntAllowed ，
        // 则需要将该字符删除至出现 maxCntAllowed 次
        if cnt > maxCntAllowed {
            ans += cnt - maxCntAllowed
            cnt = maxCntAllowed
        }

        // 下一个字符最大允许出现的次数 = 当前字符出现次数 - 1 ，
        // 同时要保证一个字符最多只能被删光
        maxCntAllowed = max(0, cnt - 1)
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
