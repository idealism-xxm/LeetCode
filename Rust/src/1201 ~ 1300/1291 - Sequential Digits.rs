// 链接：https://leetcode.com/problems/sequential-digits/
// 题意：顺次数 num 是 num 中每一个数位都比前一个数位大 1 的数字，
//      返回在 [low, high] 中的所有顺次数（从小到大排序）。

// 数据限制：
//  10 <= low <= high <= 10 ^ 9

// 输入： low = 100, high = 300
// 输出： [123,234]

// 输入： low = 1000, high = 13000
// 输出： [1234,2345,3456,4567,5678,6789,12345]


// 思路：BFS
//
//      我们先列出不同长度的顺次数：
//          一位的顺次数为 1,   2,   ..., 7,  8, 9
//          两位的顺次数为 12,  23,  ..., 78, 89
//          三位的顺次数为 123, 456, ..., 789
//          ...
//
//      可以发现如果我们已知一个顺次数 x ，如果 x % 10 < 9 ，
//      那么长度比 x 多一位的对应顺次数为 x * 10 + (x % 10 + 1) 
//
//      这样我们可以使用 BFS 遍历出所有的顺次数，初始化队列 q 中有全部的一位顺次数 [1, 9] ，
//      然后在循环中，从 q 拿出当前队首的顺次数 digits ，进行如下处理：
//          1. digits > high: 说明队列中最小的数已经比 high 大了，直接跳出循环
//          2. low < digits < high: 说明当前的数满足题意，放入到答案数组 ans 中
//          3. digits % 10 < 9: digits 有对应的下一个顺次数，
//                      将 digits * 10 + (digits % 10 + 1) 放入到队列 q 中
//
//      最后直接返回 ans 即可，因为队列是从小到大遍历的，已经保证了有序
//
//
//      总共只有 44 个数是顺次数，只会用到常数的时间和空间，
//      所以时间复杂度是 O(1) ，空间复杂度是 O(1)
//
//      时间复杂度： O(1)
//      空间复杂度： O(1)

use std::{vec, collections::VecDeque};

impl Solution {
    pub fn sequential_digits(low: i32, high: i32) -> Vec<i32> {
        let mut q: VecDeque<i32> = (1..=9).collect();
        let mut ans: Vec<i32> = vec![];
        while !q.is_empty() {
            let digits = q.pop_front().unwrap();
            if digits > high {
                break
            }
            if low <= digits {
                ans.push(digits);
            }

            let last_digit = digits % 10;
            if last_digit < 9 {
                q.push_back(digits * 10 + last_digit + 1);
            }
        }
        ans
    }
}
