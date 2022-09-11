// 链接：https://leetcode.com/problems/check-distances-between-same-letters/
// 题意：给定仅含英文小写字母组成的字符串 s ，且其中的每种字母恰好只出现两次。
//      同时给定一个长度为 26 的距离数组 distance ，下标 0-25 代表小写字母 'a'-'z' 。
//
//      判断 s 中每一种字母 ch 之间的距离，是否均为 distance[ch - 'a'] ？
//      如果一种字母不在 s 中，则不需要满足 distance 中的限制。


// 数据限制：
//  2 <= s.length <= 52
//  s 仅由英文小写字母组成
//  s 中的每一种字母恰好只出现两次
//  distance.length == 26
//  0 <= distance[i] <= 50


// 输入： s = "abaccb", distance = [1,3,0,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
// 输出： true
// 解释： - 'a' 出现在下标 0 和 2 处，所以满足 distance[0] = 1
//       - 'b' 出现在下标 1 和 5 处，所以满足 distance[1] = 3
//       - 'c' 出现在下标 3 和 4 处，所以满足 distance[2] = 0
//
//       注意： distance[3] = 5 ，但 'd' 没有在 s 中出现，它可以被忽略。
//       综上： s 是一个满足题意的字符串，返回 true 。

// 输入： s = "aa", distance = [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
// 输出： false
// 解释： - 'a' 出现在下标 0 和 2 处，所以满足 distance[0] = 1
//
//       综上： s 是一个满足题意的字符串，返回 true 。


// 思路： Map
//
//      我们可以用一个名为 firstIndex 的 map 维护 s 中每种字符第一次出现的下标。
//
//      枚举 s 中的第 i 个字母 ch ：
//          1. ch 不在 firstIndex 中: 则 ch 是第一次出现，
//              记录其下标即可， firstIndex[ch] = i
//          2. ch 在 firstIndex 中: 则 ch 是第二次出现，
//              判断其是否满足题意。
//              如果 i - firstIndex[ch] - 1 != distance[ch - 'a'] ，
//              则不满足题意，直接返回 false 即可。
//
//      最后还没返回的话， s 就是满足题意的字符串，直接返回 true
//
//
//      设字符集大小为 C 。
//
//      时间复杂度：O(n)
//          1. 需要遍历 s 中全部 O(n) 个字符一次
//      空间复杂度：O(C)
//          1. 需要维护全部 O(C) 个不同字母第一次出现的下标


func checkDistances(s string, distance []int) bool {
    // firstIndex 维护每种字母第一次出现的下标
    firstIndex := make(map[rune]int, len(s) >> 1)
    // 遍历 s 中的第 i 个字母 ch
    for i, ch := range s {
        if j, exists := firstIndex[ch]; exists {
            // ch 是第二次出现，如果其满足题意，则直接返回 false
            if i - j - 1 != distance[int(ch - 'a')] {
                return false
            }
        } else {
            // ch 是第一次出现，记录其下标即可
            firstIndex[ch] = i
        }
    }

    // 此时， s 就是满足题意的字符串，直接返回 true
    return true
}
