// 链接：https://leetcode.com/problems/longest-valid-parentheses/
// 题意：给定一个只含有 '(' 和 ')' 的字符串，求最长的有效括号子串的长度？


// 数据限制：
//  0 <= s.length <= 3 * 10 ^ 4
//  s[i] 是 '(' 或 ')'


// 输入： s = "(()"
// 输出： 2
// 解释： 最长有效括号子串为 "()"

// 输入： s = ")()())"
// 输出： 4
// 解释： 最长有效括号子串为 "()()"

// 输入： s = ""
// 输出： 0


// 思路2： 栈
//
//      与括号匹配相关的题目，很多都需要使用栈来处理。
//
//      定义 ans 维护最长合法括号子串的长度，初始化为 0 。
//
//      定义一个栈 stack ，用于存储当前未匹配的 '(' 和 ')' 的下标，
//      为了方便后续处理，初始放入 -1 ，表示有一个未匹配的 ')' 。
//
//      遍历 strs 中的第 i 个括号 ch ，进行如下处理：
//          1. ch == '(': 则必定未匹配，将 i 直接入栈
//          2. ch == ')': 则获取栈顶下标 top
//              (1) top 指向的括号是 '(': 则匹配成功，将栈顶下标出栈，
//                  此时 s[stack.top()+1..=i] 是一个合法括号子串，
//                  更新 ans 的最大值为 max(ans, i - stack.top())
//              (2) top 指向的括号不是 '(': 则匹配失败，将 i 入栈
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 s 中全部 O(n) 个括号
//      空间复杂度：O(n)
//          1. 需要维护一个栈 s ，最差情况下需要存储 s 中全部 O(n) 个未匹配的括号


impl Solution {
    pub fn longest_valid_parentheses(s: String) -> i32 {
        let s = s.as_bytes();
        // ans 表示当前最长合法括号子串的长度，初始化为 0
        let mut ans = 0;
        // stack 存储当前未匹配的 '(' 和 ')' 的下标，
        // 为了方便处理，初始放入 -1 ，表示有一个未匹配的 ')'
        let mut stack = vec![-1];
        // 带下标遍历 strs 的每个括号
        for (i, &ch) in s.iter().enumerate() {
            if ch == b'(' {
                // 如果当前是 '(' ，则必定未匹配，将其下标直接入栈
                stack.push(i as i32);
            } else {
                // 如果当前是 ')' ，则获取栈顶下标
                let top = *stack.last().unwrap();
                if top != -1 && s[top as usize] == b'(' {
                    // 如果栈顶下标对应的括号是 '(' ，则匹配成功，将栈顶下标出栈
                    stack.pop();
                    // 此时 s[stack.top()+1..=i] 是一个合法括号子串
                    ans = ans.max(i as i32 - *stack.last().unwrap());
                } else {
                    // 如果栈顶下标对应的括号不是 '(' ，则必定未匹配，
                    // 将当前 ')' 的下标入栈
                    stack.push(i as i32);
                }
            }
        }

        ans
    }
}
