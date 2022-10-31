// 链接：https://leetcode.com/problems/minimum-window-substring/
// 题意：给定两个字母串 s 和 t ，
//      求 s 中最短子串，使得其中每个字母的出现次数都大于等于 t 中对应字母的出现次数？
//
//      进阶：使用时间复杂度为 O(m + n) 的算法。


// 数据限制：
//  m == s.length
//  n == t.length
//  1 <= m, n <= 10 ^ 5
//  s 和 t 由英文字母组成


// 输入： s = "ADOBECODEBANC", t = "ABC"
// 输出： "BANC"

// 输入： s = "a", t = "a"
// 输出： "a"

// 输入： s = "a", t = "aa"
// 输出： ""
// 解释： t 中有两个 'a' ，而 s 中只有一个 'a'


// 思路：滑动窗口
//
//      如果一道题目需要在所有满足某种状态的连续子串/连续子数组中，
//      找到满足题意的一个，那么可以考虑滑动窗口。
//
//      本题需要在含有 t 中全部字母的连续子串中，找到最短的那个。
//
//      那么我们使用滑动窗口 [l, r) 表示含有 t 中全部字母的连续子串，
//      初始化为左边界 l = 0 ，右边界 r = -1 ，表示初始窗口为空。
//
//      同时我们使用名为 count 的 map 来维护滑动窗口中每个字母还需出现的次数，
//      初始化为 t 中每个字母的出现次数。
//
//      并用 remain 维护出现次数不足的不同字母数，
//      当 remain 为 0 时，滑动窗口内的连续子串满足题意。
//
//      然后不断右移左边界 l ，准备将其从到滑动窗口中移除。
//
//      题目是需要我们找到最短的含有 t 中全部字母的连续子串，
//      所以我们贪心地尽可能不扩大滑动窗口，除非其不含 t 中全部字母。
//
//      此时，我们需要不断右移右边界 r ，直至 remain 为 0 。
//      此时的滑动窗口就是以 l 为左边界的最小滑动窗口，大小为 r - l 。
//
//      然后再将 s[l] 真正从滑动窗口中移除，并更新 count 和 remain 。
//
//      我们统计所有滑动窗口的大小的最大值到 ans 中，
//      则 ans 就是含有 t 中全部字母的连续子串的长度。
//
//
//      设字符集大小为 C 。
//
//      时间复杂度： O(m + n)
//          1. 需要遍历 s 中全部 O(m) 个字母
//          2. 需要遍历 t 中全部 O(n) 个字母
//      空间复杂度： O(C)
//          1. 需要维护 count 中全部 O(C) 个不同字母的出现次数


use std::collections::HashMap;
use std::ops::{ AddAssign, SubAssign };


impl Solution {
    pub fn min_window(s: String, t: String) -> String {
        let s = s.as_bytes();
        let m = s.len();
        // count[ch] 表示滑动窗口 [l, r) 中字母 ch 还需出现的次数
        let mut count = HashMap::new();
        // 初始化为 t 中每个字母的出现次数
        for &ch in t.as_bytes() {
            count.entry(ch).or_insert(0).add_assign(1);
        }
        // remain 表示滑动窗口 [l, r) 中还需出现的不同字母数
        let mut remain = count.len();

        let (mut ans_l, mut ans_r, mut ans_len) = (0, 0, m + 1);
        // 初始化为空滑动窗口 
        let mut r = 0;
        // 右移左边界 l ，准备将其从到滑动窗口中移除
        for l in 0..m {
            // 不断右移右边界 r ，直至 remain 为 0
            while remain != 0 && r < m {
                // 滑动窗口内 s[r] 需要出现的次数 -1
                count.entry(s[r]).or_insert(0).sub_assign(1);
                // 如果等于 0 ，则表示 s[r] 在窗口中出现的次数恰好满足题意，剩余不同字母数 -1
                if count[&s[r]] == 0 {
                    remain -= 1;
                }

                r += 1
            }
            // 如果所有字母的出现次数都满足题意，且其长度更短，则更新答案
            if remain == 0 && r - l < ans_len {
                ans_l = l;
                ans_r = r;
                ans_len = r - l;
            }

            // 最后将 s[l] 从滑动窗口中移除
            count.entry(s[l]).or_insert(0).add_assign(1);
            // 如果等于 1 ，则表示 s[l] 移除后，
            // s[l] 在窗口中出现的次数恰好不满足题意，剩余不同字母数 -1
            if count[&s[l]] == 1 {  
                remain += 1;
            }
        }

        if ans_len == m + 1 {
            // 如果从未更新过，则无答案
            "".to_owned()
        } else {
            // 更新过，则获取对应的子串
            String::from_utf8(s[ans_l..ans_r].to_vec()).unwrap()
        }
    }
}
