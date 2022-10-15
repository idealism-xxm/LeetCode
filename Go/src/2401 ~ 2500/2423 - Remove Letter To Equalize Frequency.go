// 链接：https://leetcode.com/problems/remove-letter-to-equalize-frequency/
// 题意：给定一个字符串 word ，判断恰好删除其中某一个字母后，
//      word 中剩余的字母出现次数是否都相同？


// 数据限制：
//  2 <= word.length <= 100
//  word 仅由英文小写字母组成


// 输入： word = "abcc"
// 输出： true
// 解释： 删除第 3 个字母，得到 "abc" ，所有字母都只出现一次

// 输入： word = "aazz"
// 输出： false
// 解释： 删除字母 'a' ，得到 "azz" ， 'a' 出现一次， 'z' 出现两次
//       删除字母 'z' ，得到 "aaz" ， 'a' 出现两次， 'z' 出现一次


// 思路： Map
//
//      先用 chToCnt 维护 word 中每种字母的出现次数。
//
//      然后枚举删除每一个字母 ch ，将其次数减 1 。
//
//      再判断此时 chToCnt 中出现次数不为 0 的那些字母的出现次数是否相同。
//
//
//      设字符集大小为 C 。
//
//      时间复杂度：O(nC)
//          1. 需要遍历 word 中全部 O(n) 个字母，
//              每次都需要遍历 chToCnt 中全部 O(C) 种不同的元素
//      空间复杂度：O(1)
//          1. 需要维护 chToCnt 中全部 O(C) 种不同的字母的出现次数


func equalFrequency(word string) bool {
    // 统计 word 中每种字母的出现次数
    chToCnt := make(map[rune]int)
    for _, ch := range word {
        chToCnt[ch] += 1
    }

    // 枚举删除每一个字母 ch
    for _, ch := range word {
        chToCnt[ch] -= 1
        // 如果删除 ch 后满足题意，则直接返回 true
        if isOk(chToCnt) {
            return true
        }
        chToCnt[ch] += 1
    }

    // 此时无论删除哪字母都不满足题意，返回 false
    return false
}

func isOk(chToCnt map[rune]int) bool {
    // target 维护目标出现次数， -1 表示暂时还未遇到
    target := -1
    // 遍历所有非 0 的出现次数
    for _, cnt := range chToCnt {
        if cnt == 0 {
            // 如果是 0 ，则已不在 word 中，直接处理下一个
            continue
        }

        if target == -1 {
            // 如果还没有遇到过出现次数，则 target 为 cnt
            target = cnt;
        } else if cnt != target {
            // 如果出现次数不同，则不满足题意，直接返回 false
            return false;
        }
    }

    // 此时所有的非 0 出现次数都相同
    return true
}
