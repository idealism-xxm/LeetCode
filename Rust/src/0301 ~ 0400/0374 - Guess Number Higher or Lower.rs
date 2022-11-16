// 链接：https://leetcode.com/problems/guess-number-higher-or-lower/
// 题意：有一个猜数字游戏，每轮游戏会选择 1 到 n 内随机一个数 pick ，
//      通过调用给定到接口找到并返回这个数字。
//
//      给定的接口 int guess(int num) 返回值共三种情况：
//          1. -1: 选择的数字 pick 比猜测的数字 num 小，即 pick < num
//          2. 1: 选择的数字 pick 比猜测的数字 num 大，即 pick > num
//          3. 0: 选择的数字 pick 与猜测的数字 num 相等，即 pick == num


// 数据限制：
//  1 <= n <= 2 ^ 31 - 1
//  1 <= pick <= n


// 输入： n = 10, pick = 6
// 输出： 6

// 输入： n = 1, pick = 1
// 输出： 1

// 输入： n = 2, pick = 1
// 输出： 1


// 思路： 二分
//
//      题目给定的接口 guess 能判断我们猜测的数字与选中数字的大小情况，
//      那么我们可以使用二分的方法，每次都能至少排除一半可能的数字。
//
//      初始化二分区间 [l, r] 为 [1, n] ，然后不断二分判断处理。
//      设每次区间中点为 mid = (l + r) / 2 ，判断结果 result = guess(mid) ，
//      则有：
//          1. result == -1: mid 更大，则 pick 必定在 [1, mid - 1] 内，
//          2. result == 1: mid 更小，则 pick 必定在 [mid + 1, r] 内
//          3. result == 0: mid 就是 pick ，直接返回 mid 即可
//
//      题目保证选择的数在 [1, n] 内，所以必定会在二分区间为空之前返回。
//
//
//      时间复杂度：O(logn)
//          1. 需要对全部 O(n) 个数进行二分，时间复杂度为 O(logn)
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量即可


/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * unsafe fn guess(num: i32) -> i32 {}
 */

impl Solution {
    unsafe fn guessNumber(n: i32) -> i32 {
        // 二分区间为 [1, n]
        let (mut l, mut r) = (1, n);
        // 当二分区间不为空时，继续处理
        while l <= r {
            // 区间中点作为猜测的数，调用 guess 进行判断。
            // 注意这里可能存在上溢，所以采用减法的形式计算中点
            let mid = l + (r - l) / 2;
            match guess(mid) {
                // 如果 mid 更大，则 pick 必定在 [1, mid - 1] 内
                -1 => r = mid - 1,
                // 如果 mid 更小，则 pick 必定在 [mid + 1, r] 内
                1 => l = mid + 1,
                // 如果恰好就是 pick ，则直接返回
                0 => return mid,
                // 其他情况不存在，不可能走到这
                _ => unreachable!(),
            }
        }

        // 只有在二分区间为空的情况下会走到这，而题目保证选择的数在 [1, n] 内，
        // 所以必定会在二分区间为空之前返回，不可能走到这
        unreachable!()
    }
}
