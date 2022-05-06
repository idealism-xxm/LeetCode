// 链接：https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/
// 题意：给定一个字符串 s 和一个整数 k ，每次可以删除 k 个连续且相同的字符，
//      然后将被删除子串的左右两侧连在一起。
//
//      不断进行这个操作，直至不能再删除任何字符，返回最终的字符串。
//
//      题目保证最终的字符串唯一。


// 数据限制：
//  1 <= s.length <= 10 ^ 5
//  2 <= k <= 10 ^ 4
//  s 仅含有英文小写字母


// 输入： s = "abcd", k = 2
// 输出： "abcd"
// 解释： 不能执行任何删除操作

// 输入： s = "deeedbbcccbdaa", k = 3
// 输出： "aa"
// 解释： 先删除 "eee" 和 "ccc", 得到 "ddbbbdaa"
//       再删除 "bbb", 得到 "dddaa"
//       最后删除 "ddd", 得到 "aa"

// 输入： s = "pbbcggttciiippooaais", k = 2
// 输出： "ps"


// 思路： 栈
//
//      我们维护两个栈 stack 和 cnt ，
//      其中 stack 存放字符，cnt 存放字符连续出现次数。
//
//      为了方便处理，我们初始化将不存在的字符 '#' 压入 stack 中，
//      并将对应的出现次数 0 压入 cnt 中。
//
//      遍历字符串 s 的每个字符 ch ：
//          1. ch != stack.top(): 则说明 ch 是新字符，
//              则 ch 的连续出现次数 ch_cnt 是 1
//          2. ch == stack.top(): 则说明 ch 是连续出现的字符，
//              则 ch 的连续出现次数 ch_cnt 是 cnt.top() + 1
//
//          然后将 ch 压入 stack 中，并将 ch_cnt 压入 cnt 中。
//
//          若此时 cnt.top() == k ，则栈顶的 k 个字符都是 ch ，
//          可执行删除操作，弹出 stack 和 cnt 的栈顶的 k 个元素。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历全部 O(n) 个字符
//      空间复杂度：O(n)
//          1. 需要维护两个栈 stack 和 cnt ，
//              最差情况下需要保存全部 O(n) 个字符及其出现次数


impl Solution {
    pub fn remove_duplicates(s: String, k: i32) -> String {
        let k = k as usize;
        // stack 存放字符， cnt 存放字符连续出现次数
        let mut stack = vec!['#'; s.len() + 1];
        let mut cnt = vec![0; s.len() + 1];
        // 为了方便处理，我们初始化将不存在的字符 '#' 压入 stack 中，
        // 并将对应的出现次数 0 压入 cnt 中。
        let mut top = 0;
        // 遍历字符串 s 的每个字符 ch
        for ch in s.chars() {
            // 获取 ch 的连续出现次数
            let ch_cnt = if ch == stack[top] { 
                // ch 是连续出现的字符，
                // 则 ch 的连续出现次数是 cnt.top() + 1
                cnt[top] + 1 
            } else {
                // ch 是新字符，则 ch 的连续出现次数是 1
                1 
            };
            // 将 ch 压入 stack 中，并将 ch_cnt 压入 cnt 中
            top += 1;
            stack[top] = ch;
            cnt[top] = ch_cnt;
            // 若此时 cnt.top() == k ，则可执行删除操作，
            // 弹出 stack 和 cnt 的栈顶的 k 个元素。
            if ch_cnt == k {
                top -= k;
            }
        }

        // 将 stack 中的字符转换为字符串
        stack.iter().skip(1).take(top).collect()
    }
}
