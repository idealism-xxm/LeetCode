// 链接：https://leetcode.com/problems/reverse-string/
// 题意：给定一个字符数组，将这个字符数组反转。


// 数据限制：
//  1 <= s.length <= 10 ^ 5
//  s[i] 是可打印的 ascii 字符


// 输入： s = ["h","e","l","l","o"]
// 输出： ["o","l","l","e","h"]

// 输入： s = ["H","a","n","n","a","h"]
// 输出： ["h","a","n","n","a","H"]


// 思路： 双指针
//
//      定义左指针 l 和右指针 r，初始化为 0 和 s.length - 1 。
//
//      当 l < r 时，说明还需要继续交换，
//      我们可以交换 s[l] 和 s[r] ，
//      然后将 l 向右移动一位，将 r 向左移动一位。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 s 中全部 O(n) 个字符
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


impl Solution {
    pub fn reverse_string(s: &mut Vec<char>) {
        // 定义左指针 l ，初始化为 0
        let mut l = 0;
        // 定义右指针 r ，初始化为 s.length - 1
        let mut r = s.len() - 1;
        // 当 l < r 时，需要继续交换
        while l < r {
            // 交换 s[l] 和 s[r]
            s.swap(l, r);
            // l 向右移动一位
            l += 1;
            // r 向左移动一位
            r -= 1;
        }
    }
}
