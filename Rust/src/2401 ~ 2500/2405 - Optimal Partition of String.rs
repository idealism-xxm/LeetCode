// 链接：https://leetcode.com/problems/optimal-partition-of-string/
// 题意：给定一个字符串 s ，将其划分成多个子串，
//      使得每个子串中的每种字符最多只出现一次。
//
//      求最少能划分成多少个子串？


// 数据限制：
//  1 <= s.length <= 10 ^ 5
//  s 仅含有英文小写字母


// 输入： s = "abacaba"
// 输出： 4
// 解释： 有两种最优的划分方式： ("a","ba","cab","a") 和 ("ab","a","ca","ba")

// 输入： s = "ssssss"
// 输出： 6
// 解释： 唯一的划分方式为 ("s","s","s","s","s","s")


// 思路： 贪心 + Set
//
//      为了使划分的子串尽可能少，那我们在划分的时候，
//      必定要贪心地将当前字符 ch 塞入前一个子串中，
//      除非前一个子串已含有字符 ch 。
//
//      可以用一个名为 used 的集合统计每个划分的子串中的字符集合。
//
//      然后遍历 s 中的每一个字符 ch ，进行如下处理：
//          1. 如果 ch 在 used 中，则 ch 只能在一个新划分的子串中，
//              则计入当前子串，清空已使用的字符
//          2. 此时 ch 必定不在 used 中，直接放入 used 中即可
//
//      最后还要对划分的结果再 +1 ，因为最后划分的子串没有计入。
//
//
//      设字符集大小为 C 。
//
//      时间复杂度：O(n)
//          1. 需要遍历 s 中全部 O(n) 个字符
//      空间复杂度：O(C)
//          1. 需要维护每个子串中全部不同的字符，最差情况下有 O(C) 个


use std::collections::HashSet;


impl Solution {
    pub fn partition_string(s: String) -> i32 {
        // ans 维护当前已经划分出的子串数
        let mut ans = 0;
        // used 表示当前正在划分的子串中的字符集合
        let mut used = HashSet::new();
        for ch in s.as_bytes().iter() {
            // 如果 ch 已在 used 中，则 ch 只能在一个新划分的子串中，
            // 则计入当前子串，清空已使用的字符
            if used.contains(ch) {
                ans += 1;
                used.clear();
            }

            used.insert(ch);
        }

        // 最后一个正在划分的子串没有计入，需要手动 +1
        ans + 1
    }
}
