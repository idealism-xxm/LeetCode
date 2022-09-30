// 链接：https://leetcode.com/problems/length-of-the-longest-alphabetical-continuous-substring/
// 题意：给定一个小写字母字符串 s ，求其中最长的子串字母序连续字符串？
//
//      字母序连续字符串指字母表中连续字母组成的字符串，
//      即 "abcdefghijklmnopqrstuvwxyz" 任意子串。


// 数据限制：
//  1 <= s.length <= 10 ^ 5
//  s 仅由英文小写字母组成


// 输入： s = "abacaba"
// 输出： 2
// 解释： 共有 4 个不同的字母序连续字符串： "a", "b", "c", "ab" ，
//       其中最长的是 "ab" ，长度为 2

// 输入： s = "abcde"
// 输出： 5
// 解释： "abcde" 本身就是最长的字母序连续字符串，长度为 5


// 思路： 贪心
//
//      对于字母序连续字符串来说，第一个字母确定后，后续的字母都会确定。
//
//      所以我们可以维护当前字母序连续字符串的长度 cnt ，以及所有 cnt 的最大值 ans ，
//      然后遍历 s 中的第 i 个字母，进行如下处理：
//          1. s[i - 1] + 1 == s[i]: 贪心地将其纳入字母序连续字符串中，
//              则其长度 cnt += 1
//          2. s[i - 1] + 1 != s[i]：则只能以当前 s[i] 为新的字母序连续字符串。
//              此时先更新 ans 的最大值为 max(ans, cnt) ，然后另长度 cnt = 1
//
//      最后，要计入最后一个字母序连续字符串的长度，即 ans = max(ans, cnt) 。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 s 中全部 O(n) 个字符一次
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量即可


impl Solution {
    pub fn longest_continuous_substring(s: String) -> i32 {
        let mut s = s.as_bytes();
        // ans 维护最长的字母序连续字符串的长度
        let mut ans = 0;
        // cnt 表示当前字母序连续字符串的长度，
        // 初始为字母序连续字符串仅由第一个字母组成，长度为 1
        let mut cnt = 1;
        // 遍历 s 中的每个字母
        for i in 1..s.len() {
            if s[i - 1] + 1 == s[i] {
                // 如果是连续的字母，则字母序连续字符串的长度 +1
                cnt += 1;
            } else {
                // 如果不是连续的字母，则先更新 ans 的最大值
                ans = ans.max(cnt);
                // 再以当前字母为新的字母序连续字符串，长度为 1
                cnt = 1;
            }
        }

        // 注意要计入最后一个字母序连续字符串的长度
        ans.max(cnt)
    }
}
