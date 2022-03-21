// 链接：https://leetcode.com/problems/partition-labels/
// 题意：给定一个字符串 s ，将其划分为尽可能多的片段，
//      使得同一字母最多出现在一个片段中。
//
//      返回一个表示每个字符串片段的长度的列表。


// 数据限制：
//  1 <= s.length <= 500
//  s 仅由英文小写字母组成


// 输入： s = "ababcbacadefegdehijhklij"
// 输出： [9,7,8]
// 解释： 划分结果为 "ababcbaca", "defegde", "hijhklij" ，
//       每个字母最多出现在一个片段中。
//       注意，划分为 "ababcbacadefegde", "hijhklij" 是错误的，
//       因为划分的片段数较少。


// 输入： s = "eccbbbbdec"
// 输出： [10]
// 解释： 划分结果为 "eccbbbbdec" ，所有字母必须都在一个字符串中。


// 思路： 贪心
//
//      我们可以用 last_index 维护每个字符最后一次出现的位置。
//
//      用 curPartStart 和 cur_part_end 表示当前片段的起始和结束位置，
//      初始化均为 0 。
//
//      重新遍历 s 中的第 i 个字符 ch ，
//      因为需要让 ch 出现在同一个片段中，
//      所以当前片段至少要持续到 last_index[ch] 。
//
//      如果 cur_part_end < last_index[ch] ，
//      则表明当前片段应该更新为到 last_index[ch] 结束，
//      即 cur_part_end = last_index[ch] 。
//
//      然后再判断 cur_part_end 是否就是当前遍历的字符下标，
//      如果是，那么这个片段中出现的所有字符都仅在当前片段中，
//      贪心切下当前片段即可。
//
//      收集当前片段的长度 cur_part_end - curPartStart + 1 ，
//      并更新下一个片段的起始下标 curPartStart = i + 1 。
//
//
//      设 K 为字符集大小
//      
//      时间复杂度：O(n)
//          1. 需要遍历 s 中全部 O(n) 个字符
//      空间复杂度：O(K)
//          1. 需要维护 last_index ，存储全部 O(K) 个不同的字符
//          2. 需要维护 ans ，最多会有 O(K) 个片段


func partitionLabels(s string) []int {
    // lastIndex[ch] 表示 ch 在 s 中的最后一个出现的位置
    lastIndex := make(map[rune]int)
    // 带下标遍历 s 中的字符
    for i, ch := range s {
        // 更新每个字符最后一次出现的位置
        lastIndex[ch] = i
    }

    // ans 收集所有片段的长度
    ans := make([]int, 0)
    // curPartStart 表示当前片段的起始下标
    curPartStart := 0
    // curPartEnd 表示当前片段的结束下标
    curPartEnd := 0
    // 带下标遍历 s 中的字符
    for i, ch := range s {
        // 如果 ch 最后出现的位置大于当前片段的结束下标
        if lastIndex[ch] > curPartEnd {
            // 那么 curPartEnd 需要更新为 ch 最后出现的位置
            curPartEnd = lastIndex[ch]
        }
        // 如果此时 i 是当前片段的结束下标，
        // 那么这个片段中出现的所有字符都仅在当前片段中，
        // 贪心切下当前片段即可
        if curPartEnd == i {
            // 收集这个片段的长度即可
            ans = append(ans, curPartEnd - curPartStart + 1)
            // 下一个片段的起始下标为 i + 1
            curPartStart = i + 1
        }
    }

    // 现在 ans 中就是所有片段的长度
    return ans
}
