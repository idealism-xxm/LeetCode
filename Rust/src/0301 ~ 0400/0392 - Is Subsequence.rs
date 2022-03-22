// 链接：https://leetcode.com/problems/is-subsequence/
// 题意：给定两个字符串 s 和 t ，判断 s 是否是 t 的子序列？
//
//      一个字符串的子序列是删除其中一些字符后得到的新字符串。


// 数据限制：
//  0 <= s.length <= 100
//  0 <= t.length <= 10 ^ 4
//  s 和 t 仅由英文小写字母组成


// 输入： s = "abc", t = "ahbgdc"
// 输出： true
// 解释： "ahbgdc" 删除 'h', 'g', 'd' 后得到字符串 "abc" ，
//       所以 s 是 t 的子序列

// 输入： s = "axc", t = "ahbgdc"
// 输出： false
// 解释： t 没有 s 中的字母 'x' ，
//       所以 s 不是 t 的子序列


// 思路： 双指针
//
//      定义双指针 l 和 r ， 
//      l 表示 s 下一个待匹配的字符，
//      t 表示 t 中下一个可匹配的字符。
//
//      然后不断循环匹配 s[l] 和 t[r] ，直至 s 或者 t 全部遍历完。
//
//      每次匹配时，如果 s[l] == t[r] ，
//      则当前 s 中的字符匹配成功，右移左指针 l 。
//
//      无论是否匹配成功，每次 t[r] 中的字符都已进行过匹配，
//      需要右移右指针 r 。
//		
//
//		时间复杂度： O(|s| + |t|)
//          1. 需要遍历 s 中的全部 O(|s|) 个字符
//          2. 需要遍历 t 中的全部 O(|t|) 个字符
//		空间复杂度： O(1)
//          1. 只需要定义两个额外指针变量


impl Solution {
    pub fn is_subsequence(s: String, t: String) -> bool {
        // 将字符串转成字节切片，方便后续使用
        let (s, t) = (s.as_bytes(), t.as_bytes());
        // 定义左指针 l ，表示 s 下一个待匹配的字符
        let mut l = 0;
        // 定义右指针 r ，表示 t 中下一个可匹配的字符
        let mut r = 0;

        // 如果 s 和 t 均未遍历完，则需要继续循环匹配
        while l < s.len() && r < t.len() {
            // 如果 s 中的字符与 t 中的字符相同，
            // 则匹配成功，移动左指针匹配下一个字符
            if s[l] == t[r] {
                l += 1;
            }
            // 无论是否匹配成功，都需要右移右指针
            r += 1;
        }

        // 最后如果 s 中的字符全部匹配成功，则说明 s 是 t 的子序列
        l == s.len()
    }
}