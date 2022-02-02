// 链接：https://leetcode.com/problems/string-to-integer-atoi/
// 题意：给定一个字符串，返回可以表示的整数（题目要求更像是 JavaScript 中的 parseInt 函数）
//		从第一个非空白字符开始，到第一个非数字字符止（前闭后开，且包含：+ -）
//		返回这一段字符表示的整数（超出 32 位整型，则返回相应的最值）。

// 数据限制：
//  0 <= s.length <= 200
//  s 仅由英文字母（大小写）、数字 (0-9) 、 ' ' 、 '+' 、 '-' 和 '.' 组成

// 输入：s = "42"
// 输出：42
// 解释：用于转换成数字的字符串为 s[0..2] = "42" ，
//      对应的整数是 42

// 输入：s = "   -42"
// 输出：-42
// 解释：用于转换成数字的字符串为 s[3..6] = "-42" ，
//      对应的整数是 -42

// 输入：s = "4193 with words"
// 输出：4193
// 解释：用于转换成数字的字符串为 s[0..4] = "4193" ，
//      对应的整数是 4193

// 输入：s = "-91283472332"
// 输出：-2147483648
// 解释：用于转换成数字的字符串为 s[0..12] = "-91283472332" ，
//      对应的整数是 -91283472332 ，
//      但这个整数不在 32 位整型范围内，
//      所以要返回最小值 -2147483648


// 思路： 模拟
//
//      1. 去除字符串空白前缀
//      2. 根据第一个字符判断符号 sign 和初始值 num
//      3. 遍历剩余的数字字符，直至遇到非数字字符或者字符串结束，
//          对每个字符 ch 执行 num = num * 10 + int(ch) ，
//          注意判断乘法和加法溢出时，要根据 sign 返回最大值/最小值
//      4. 最后返回带符号的结果 sign * num
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

impl Solution {
    pub fn my_atoi(s: String) -> i32 {
        // 去除空白前缀，并转成迭代器
        let mut iter = s.trim_start().chars();
        // 根据第一个字符判断符号和初始值
        let (sign, mut num) = match iter.next() {
            // 如果是正号，则最后乘以 1 ，当前整数初始值为 0
            Some('+') => (1, 0),
            // 如果是负号，则最后乘以 -1 ，当前整数初始值为 0
            Some('-') => (-1, 0),
            // 如果是数字，则最后乘以 1 ，当前整数初始值为 int(ch)
            Some(ch) if ch.is_digit(10) => (1, ch.to_digit(10).unwrap() as i32),
            // 如果没有字符，或者第一个字符不是数字和正负号，则返回 0
            _ => return 0,
        };
        // 如果是正数，则溢出应该返回 i32::MAX
        // 如果是负数，则溢出应该返回 i32::MIN
        let overflow_num = if sign == 1 { i32::MAX } else { i32::MIN };
        // 遍历后续的数字字符，计算结果
        for ch in iter.take_while(|ch| ch.is_digit(10)) {
            // 对 num 乘 10 ，如果溢出，则返回 overflowing_num
            let (res, overflow) = num.overflowing_mul(10);
            if overflow {
                return overflow_num;
            }
            // 对 res + int(ch) ，如果溢出，则返回 overflowing_num
            let (res, overflow) = res.overflowing_add(ch.to_digit(10).unwrap() as i32);
            if overflow {
                return overflow_num;
            }

            num = res;
        }

        // 返回带符号的结果
        return sign * num;
    }
}
