// 链接：https://leetcode.com/problems/add-digits/
// 题意：给定一个非负整数 num ，反复将各个位上的数字相加，
//      直到结果为一位数，然后返回这个结果。


// 数据限制：
//  0 <= num <= 2 ^ 31 - 1

// 输入：num = 38
// 输出："e"
// 解释：各位相加的过程为：
//      38 --> 3 + 8 --> 11
//      11 --> 1 + 1 --> 2
//      2 是一位数，直接返回 2

// 输入：num = 0
// 输出：0
// 解释：0 是一位数，直接返回 0


// 思路：模拟
//
//      直接根据题意模拟即可，使用两个循环处理。
//
//      外层循环判断 num 是不是一位数，
//      如果不是一位数就继续求每一位的和，
//      其中使用 res 维护 num 每一位的和。
//
//      内层循环判断 num 是不是 0 ，
//      如果不是 0 ，则 num 中还有数位需要统计，
//      在内层循环中先对 res 加上 num 个位的数 num % 10 ，
//      然后将 num 以十进制右移一位。
//
//      内层循环结束后， res 就是 num 每一位的和，
//      此时将 res 的赋值给 num 就完成了外层循环的操作。
//
//
//      进阶：不使用循环/递归在 O(1) 的时间复杂度内求出最终结果。
//
//      进阶的解法需要用到数学上的知识点，感兴趣的可以查看讨论区，
//      这里就不展开讲了，仅提及一下具体方法：
//          1. num == 0: 返回 0
//          2. num % 9 == 0: 返回 9
//          3. num % 9 != 0: 返回 num % 9
//
//
//      时间复杂度：O(lg(num))
//      空间复杂度：O(1)


impl Solution {
    pub fn add_digits(mut num: i32) -> i32 {
        // 当 num 不是一位数时，需要继续求每一位的和
        while num > 9 {
            // res 维护每一位的和，初始化为 0
            let mut res = 0;
            // 当 num 不是 0 时，还有数位需要统计
            while num > 0 {
                // res 加上 num 个位的值
                res += num % 10;
                // num 以十进制右移一位
                num /= 10;
            }
            // num 现在就是所有位的和
            num = res;
        }

        num
    }
}
