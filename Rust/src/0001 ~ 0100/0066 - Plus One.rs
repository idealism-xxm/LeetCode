// 链接：https://leetcode.com/problems/plus-one/
// 题意：给定一个非空整型数组，表示一个非负无前导零的整数
//      对这个整数加一，返回结果数组？


// 数据限制：
//  1 <= digits.length <= 100
//  0 <= digits[i] <= 9
//  digits 不含前导零


// 输入： digits = [1,2,3]
// 输出： [1,2,4]
// 解释： digits 表示的整数为 123 ，
//       对其加上 1 就是 123 + 1 = 124

// 输入： [4,3,2,1]
// 输出： [4,3,2,2]
// 解释： digits 表示的整数为 4321 ，
//       对其加上 1 就是 4321 + 1 = 4322

// 输入： [9]
// 输出： [1,0]
// 解释： digits 表示的整数为 9 ，
//       对其加上 1 就是 9 + 1 = 10


// 思路： 模拟
//
//      从个位开始计算，用 carry 维护每一位的进位即可。
//
//      注意最高位产生进位时（即最后 carry 不为 0 时），
//      需要在结果数组前面加一个 carry 。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历计算全部 O(n) 个数位
//      空间复杂度：O(n)
//          1. 最高位需要进位时，需要生产一个 O(n) 的数组


impl Solution {
    pub fn plus_one(mut digits: Vec<i32>) -> Vec<i32> {
        // carry 表示当前位的进位，初始化为 1 ，表示对个位加 1
        let mut carry = 1;
        for i in (0..digits.len()).rev() {
            // 计算第 i 位的中间计算结果
            digits[i] += carry;
            // 计算对第 i - 1 位产生的进位
            carry = digits[i] / 10;
            // 计算第 i 位的实际结果，保证在范围在 [0, 9] 内
            digits[i] %= 10;
        }

        // 如果最高位有进位，则需要在结果数组前面加一个 carry
        if carry != 0 {
            digits.insert(0, carry)
        }

        digits
    }
}
