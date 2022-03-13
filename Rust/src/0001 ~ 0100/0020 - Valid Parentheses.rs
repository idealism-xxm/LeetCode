// 链接：https://leetcode.com/problems/valid-parentheses/
// 题意：给定一个由 (, ), {, }, [ 和 ] 组成的括号字符串，
//      判断括号字符串是否合法？


// 数据限制：
//  1 <= s.length <= 10 ^ 4
//  s 仅有 "()[]{}" 中的括号组成


// 输入： s = "()"
// 输出： true

// 输入： s = "()[]{}"
// 输出： true

// 输入： s = "(]"
// 输出： false


// 思路： 栈
//
//		因为所有右括号都是与最近的左括号匹配的，
//      所以可以用栈来记录所有未匹配的左括号。
//
//      当遇到左括号时，将其压入栈中。
//      当遇到右括号时，判断栈顶元素是否与之匹配的左括号，
//      如果不匹配，则直接返回 false，否则将栈顶元素弹出。
//
//      最后判断栈是否为空，
//          1. 栈为空，即栈中没有左括号，
//              则所有左右括号都匹配成功，返回 true
//          2. 栈不为空，即栈中还有左括号，
//              则还有左括号没有匹配成功，返回 false
//
//
//      时间复杂度：O(n)
//          1. 需要遍历全部 O(n) 个字符
//      空间复杂度：O(n)
//          1. 需要存储未匹配的左括号，
//              最差情况下有 O(n) 个左括号未匹配

use std::collections::{VecDeque, HashMap};

impl Solution {
    pub fn is_valid(s: String) -> bool {
        // 初始化右括号匹配的左括号
        let closed_to_open: HashMap<u8, u8> = HashMap::from([
            (b')', b'('),
            (b'}', b'{'),
            (b']', b'['),
        ]);

        // 定义一个栈，只存储左括号，
        // 栈顶左括号表示下一个右括号需要匹配的
        let mut stack = VecDeque::new();
        // 遍历 s 中的括号
        for &ch in s.as_bytes() {
            match ch {
                // 如果是左括号，则直接入栈
                b'(' | b'{' | b'[' => stack.push_back(ch),
                // 如果是右括号，则此时栈顶左括号出栈，
                // 如果其不匹配当前右括号，则直接返回 false
                b')' | b'}' | b']' => {
                    if stack.pop_back().as_ref() != closed_to_open.get(&ch) {
                        return false;
                    }
                }
                // 题目数据保证不存在这种情况
                _ => unreachable!(),
            }
        }

        // 此时所有右括号都匹配成功，
        //  1. 栈中没有左括号，则所有左右括号都匹配成功，返回 true
        //  2. 栈中还有左括号，则还有左括号没有匹配成功，返回 false
        stack.is_empty()
    }
}
