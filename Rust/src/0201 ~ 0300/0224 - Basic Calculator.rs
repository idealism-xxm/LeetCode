// 链接：https://leetcode.com/problems/basic-calculator/
// 题意：给定一个包含加减和括号的算式（整数非负），计算结果？

// 输入： "1 + 1"
// 输出： 2

// 输入： " 2-1 + 2 "
// 输出： 3

// 思路： 调度场算法
//
//      其实大学学过简化版的，就是用两个栈直接计算，不转换成后缀表达式
//
//      推荐中英结合看，中文有详细的步骤说明（说明了左结合和右结合运算符的区别）
//      https://zh.wikipedia.org/wiki/%E8%B0%83%E5%BA%A6%E5%9C%BA%E7%AE%97%E6%B3%95
//
//      不断读取输入获取下一个 token
//      1. token 是数字，加入到数字栈中
//      2. token 是左括号，加入到符号栈中
//      3. token 是右括号，不断弹出符号栈中一个符号，直至弹出的符号是左括号
//          循环体内中不断弹出数字栈顶两个数字和进行对应的计算，然后将结果放入数字栈中
//      4. token 是操作符，记为 cur_operator ，记录符号栈顶的操作符为 last_operator ，
//          (1) cur_operator 是左结合性，那么 cur_operator 的优先级
//              小于等于 last_operator 的优先级，则可对数字栈顶的两个数字
//              执行 last_operator 对应的计算
//          (2) cur_operator 是右结合性，那么 cur_operator 的优先级
//              小于 last_operator 的优先级，则可对数字栈顶的两个数字
//              执行 last_operator 对应的计算
//
//      时间复杂度： O(n)
//      空间复杂度： O(n)

use std::collections::HashMap;
use std::collections::VecDeque;

lazy_static!{
    // 初始化每个运算符的优先级
    static ref OPERATOR_PRIORITY: HashMap<char, i8> = vec![
        ('$', 0),
        ('+', 1),
        ('-', 1),
    ].into_iter().collect();
    // 初始化每个运算符的左右结合性
    static ref OPERATOR_ASSOCIATIVITY_LEFT: HashMap<char, bool> = vec![
        ('$', true),
        ('+', true),
        ('-', true),
    ].into_iter().collect();
}

impl Solution {
    pub fn calculate(s: String) -> i32 {
        // 添加一个优先级最低的符号，方便后续处理
        let s = s + "$";
        let mut num = 0;
        let mut is_num = false;
        let mut num_stack: VecDeque<i32> = VecDeque::new();
        let mut operator_stack: VecDeque<char> = VecDeque::new();
        for ch in s.chars() {
            if ch.is_numeric() {
                // 如果当前是数字，则将 is_num 置为 true
                is_num = true;
                // 并把当前位的数加在 num 上
                num = num * 10 + ch.to_digit(10).unwrap() as i32;
            } else {
                if is_num {
                    // 如果前面的 token 是数字，则需要放入 num_stack
                    num_stack.push_back(num);
                    // 当前字符开始已不是数字
                    is_num = false;
                    num = 0;
                }
                match ch {
                    // 左括号直接入栈
                    '(' => operator_stack.push_back('('),
                    ')' => {
                        // 不断弹出数字栈顶两个元素和符号栈顶一个元素进行计算
                        // 直至遇到左括号
                        while let Some(operator) = operator_stack.pop_back() {
                            // 遇到左括号跳出循转
                            if operator == '(' {
                                break;
                            }
                            // 执行计算
                            Solution::do_calculate(&mut num_stack, operator);
                        }
                    },
                    '+' | '-' | '$' => {
                        // 若符号栈中还有元素，则进行循环
                        while let Some(operator) = operator_stack.back() {
                            // 判断当前是否满足弹出计算的条件，不满足则跳出
                            if !Solution::should_calculate(*operator, ch) {
                                break;
                            }
                            // 对数字栈顶两个数字执行符号栈顶对应对运算
                            Solution::do_calculate(&mut num_stack, operator_stack.pop_back().unwrap())
                        }
                        // 当前符号入栈
                        operator_stack.push_back(ch);
                    },
                    ' ' => {},
                    _ => unreachable!(),
                }
            }
        }
        // 最后栈顶元素就是最终结果
        num_stack.pop_back().unwrap()
    }

    // 对 num_stack 栈顶两个数字执行 operator 对应对计算，然后将结果入栈
    fn do_calculate(num_stack: &mut VecDeque<i32>, operator: char) {
        // 弹出栈顶两个数字并计算结果
        let right = num_stack.pop_back().unwrap();
        let left = num_stack.pop_back().unwrap();
        let result = match operator {
            '+' => left + right,
            '-' => left - right,
            _ => unreachable!(),
        };
        // 结果数字入栈
        num_stack.push_back(result);
    }

    // 根据栈中的符号和当前符号判断是否需要计算
    fn should_calculate(last_operator: char, cur_operator: char) -> bool {
        // 获取运算符的优先级
        let last_priority = OPERATOR_PRIORITY.get(&last_operator);
        // 如果 last_operator 不是操作符号，则直接返回 false
        if last_priority.is_none() {
            return false;
        }
        let last_priority = last_priority.unwrap();
        let cur_priority = OPERATOR_PRIORITY.get(&cur_operator).unwrap();
        // 获取 cur_operator 的结合性
        let cur_associativity_left = OPERATOR_ASSOCIATIVITY_LEFT.get(&cur_operator).unwrap();

        if *cur_associativity_left {
            // 如果当前运算符是左结合，那么其优先级小于等于栈顶运算符的优先级，就要先执行计算
            cur_priority <= last_priority
        } else {
            // 如果当前运算符是右结合，那么其优先级小于栈顶运算符的优先级，就要先执行计算
            cur_priority < last_priority
        }
    }
}
