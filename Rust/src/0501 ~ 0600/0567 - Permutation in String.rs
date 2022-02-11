// 链接：https://leetcode.com/problems/permutation-in-string/
// 题意：给定两个字符串 s1 和 s2 ，判断 s1 的排列之一是否为 s2 的 子串 ？


// 数据限制：
//  1 <= s1.length, s2.length <= 10 ^ 4
//  s1 和 s2 仅包含小写字母


// 输入：s1 = "ab", s2 = "eidbaooo"
// 输出：true
// 解释：s1 的排列之一 "ba" 是 s2 的子串

// 输入：s1 = "ab", s2 = "eidboaoo"
// 输出：false
// 解释：s1 的任何排列均不是 s2 的子串


// 思路：滑动窗口
//
//      先统计 s1 中每个字符的出现次数到 ch_to_cnt 中，
//      然后使用滑动窗口遍历字符串 s2 统计满足题意的起始下标即可。
//
//      设我们的滑动窗口的左右指针下标为 l 和 r ，初始都为 0 。
//      子串 s2[l..=r] 内每个字符出现的次数都小于等于 s1 中对应字符出现的次数。
//
//      那么每次加入字符 s2[r] 时，都要消耗 s2[r] 的次数，
//      即 ch_to_cnt[s2[r]] -= 1 。
//      如果此时 ch_to_cnt[s2[r]] == -1 ，则说明子串中 s2[r] 出现的次数过多，
//      需要不断右移左指针 l ，直至把一个与 s2[r] 相同的字符移除子串。
//
//      这样处理后， s[l..=r] 仍满足每个字符出现次数都不会超，
//      那么只要此时子串的长度 r - l + 1 是 s1.len() ，
//      就说明 s2[l..=r] 是 s1 的一个排列，直接返回 true 。
//
//      遍历完成还没找到满足题意的子串，则返回 false 。
//
//      【进阶】 LeetCode 438 是本题的升级版，需要找到所有这种子串的起始下标。
//
//      时间复杂度：O(|s1| + |s2|)
//      空间复杂度：O(|s1|)

impl Solution {
    pub fn check_inclusion(s1: String, s2: String) -> bool {
        // 统计 s1 中每个字符出现的次数
        let mut ch_cnt = vec![0; 26];
        for &c in s1.as_bytes() {
            ch_cnt[(c as u8 - b'a') as usize] += 1;
        }

        // 先将 s2 转成字符切片，方便后续遍历和引用
        let s2 = s2.as_bytes();
        // [l, r] 这个滑动窗口内的字符完全合法
        let mut l = 0;
        // 带下标遍历 s2 中的每个字符
        for (r, ch) in s2.iter().map(|ch| (ch - b'a') as usize).enumerate() {
            // 将 s2[r] 放入到滑动窗口内，即消耗 ch 字符的可使用次数 1 次
            ch_cnt[ch] -= 1;
            // 如果 ch 可使用次数变为 -1 ，
            // 则说明滑动窗口 [l, r] 内的字符 ch 超出可使用次数了，
            // 需要不断右移 l ，直至移除一个 ch 字符
            if ch_cnt[ch] == -1 {
                loop {
                    // 移除 s2[l] ，则其对应的字符可使用次数 +1
                    let ch_l = (s2[l] - b'a') as usize;
                    ch_cnt[ch_l] += 1;
                    l += 1;
                    // 如果刚刚移除了与 ch 相同的字符，
                    // 则现在 [l, r] 内的字符出现次数再次合法，
                    // 不再右移 l ，跳出循环
                    if ch_l == ch {
                        break;
                    }
                }
            }

            // 如果次数滑动窗口的长度 r - l + 1 恰好是 s1.len() ，
            // 则说明 s[l..=r] 就是 s1 的一个排列，
            // 直接返回 true
            if r - l + 1 == s1.len() {
                return true;
            }
        }

        // 遍历完成还没找到满足题意的子串，则返回 false
        false
    }
}
