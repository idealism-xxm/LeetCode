// 链接：https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
// 题意：给定一个由 (, ) 和英文小写字母组成的字符串，
//      移除最少的括号（ ( 或 ) ），使得结果字符串是合法的括号字符串，
//      并返回这个括号字符串。
//
//      合法的「括号字符串」满足以下任意一个条件：
//          1. 空字符串或只包含小写字母的字符串
//          2. 可以被写作 AB (A 连接 B) 的字符串，其中 A 和 B 都是合法的「括号字符串」
//          3. 可以被写作 (A) 的字符串，其中 A 是一个合法的「括号字符串」


// 数据限制：
//  1 <= s.length <= 10 ^ 5
//  s[i] 是 '(' , ')' 或者英文小写字母


// 输入： s = "lee(t(c)o)de)"
// 输出： "lee(t(c)o)de"
// 解释： "lee(t(co)de)", "lee(t(c)ode)" 也是合法的括号字符串

// 输入： s = "a)b(c)d"
// 输出： "ab(c)d"

// 输入： s = "))(("
// 输出： ""
// 解释： 空字符串也是合法的括号字符串


// 思路： 栈
//
//      LeetCode 20 - 有效的括号 这题的进阶版。
//
//      像这种括号匹配的题目基本都需要用栈来处理。
//
//		因为所有右括号都是与最近的左括号匹配的，
//      所以可以用栈来记录所有未匹配的左括号。
//
//      同时，我们维护一个数组 available ，
//      其中 available[i] 表示 s[i] 是否合法。
//
//      然后我们遍历字符串 s 的第 i 个字符 ch ：
//          1. ch == '(': 当前左括号可能会与后续的右括号匹配，
//              先把当前下标 i 压入栈中
//          2. ch == ')': 如果栈不为空，则栈顶左括号和当前右括号匹配，
//              弹出该左括号的下标 j ，标记 s[j] 和 s[i] 为合法的括号，
//          3. ch == 'a' - 'z': 英文小写字母必定合法
//
//      遍历完成后， available 中为 true 的下标就是 s 中合法字符的下标，
//      将这些字符收集成字符串返回即可
//
//
//      时间复杂度：O(n)
//          1. 需要遍历全部 O(n) 个字符
//      空间复杂度：O(n)
//          1. 需要用 available 来记录 s 中 O(n) 个字符是否合法
//          2. 需要存储未匹配的左括号，最差情况下有 O(n) 个左括号未匹配
//          3. 生成的结果字符串最差情况下有 O(n) 个字符


impl Solution {
    pub fn min_remove_to_make_valid(s: String) -> String {
        // available[i] 表示 s[i] 是否合法，初始化均不合法
        let mut available = vec![false; s.len()];
        // stack 存储当前可供匹配的左括号
        let mut stack = vec![];
        // 带下标遍历 s 中的每一个字符
        for (i, ch) in s.chars().enumerate() {
            // 根据字符种类进行不同处理
            match ch {
                // 如果是左括号，则直接把当前下标压入栈中，
                // 因为当前左括号可能会与后续的右括号匹配
                '(' => stack.push(i),
                // 如果是右括号
                ')' => {
                    // 如果栈中还有左括号，则当前括号对合法，
                    // 弹出该左括号的下标 j
                    if let Some(j) = stack.pop() {
                        // 标记左括号合法
                        available[j] = true;
                        // 标记右括号合法
                        available[i] = true;
                    }
                }
                // 如果是其他字符，则必定合法，标记即可
                _ => available[i] = true,
            }
        }

        s
            // 遍历 s 中的每一个字符
            .chars()
            // 带上下标
            .enumerate()
            // 只过滤合法的字符
            .filter(|(i, _)| available[*i])
            // 只取字符
            .map(|(_, ch)| ch)
            // 收集成字符串
            .collect()
    }
}