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


impl Solution {
    pub fn reverse_prefix(word: String, ch: char) -> String {
        // 找到第一个 ch ，并将 word 分为两部分
        match word.split_once(ch) {
            Some((prefix, suffix)) => {
                // 如果存在 ch ，则翻转前缀串
                ch.to_string() + 
                &prefix.chars().rev().collect::<String>() + 
                suffix
            },
            // 此时不存在 ch ，直接返回原字符串
            None => word,
        }
    }
}
