// 链接：https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// 题意：给定一个升序排序的数组 numbers 和一个整数 target ，
//      题目保证有且仅有一对索引 (l, r) ，
//      满足 numbers[l] + numbers[r] = target ，
//      找到这对索引，并返回 [l + 1, r + 1] 。
//
//      要求：空间复杂度为 O(1) 。


// 数据限制：
//  2 <= numbers.length <= 3 * 10 ^ 4
//  -1000 <= numbers[i] <= 1000
//  numbers 是升序排序的
//  -1000 <= target <= 1000
//  保证有且仅有一个合法解


// 输入： numbers = [2,7,11,15], target = 9
// 输出： [1,2]
// 解释： 2 与 7 的和为 9 ，
//       对应的索引为 (0, 1) ，因此返回 [1,2] 。

// 输入： numbers = [2,3,4], target = 6
// 输出： [1,3]
// 解释： 2 + 4 = 6 ，
//       对应的索引为 (0, 2) ，因此返回 [1,3] 。

// 输入： numbers = [-1,0], target = -1
// 输出： [1,2]
// 解释： -1 + 0 = -1 ，
//       对应的索引为 (0, 1) ，因此返回 [1,2] 。


// 思路： 双指针
//
//      题目保证有且仅有一对合法的索引，且数组是升序排序的，
//      所以我们可以采用双指针的方法从两端往中间扫。
//
//      我们定义双指针 l = 0; r = numbers.length - 1;
//      不断循环处理，直至返回结果。
//
//      每次循环时，记 sum = numbers[l] + numbers[r] ，
//      则有以下三种情况：
//          1. sum == target: (l, r) 是合法解，直接返回 [l + 1, r + 1]
//          2. sum < target: 则说明当前和较小，需要让和变大，
//              那么只能将 l 向右移动一位
//          3. sum > target: 则说明当前和较大，需要让和变小，
//              那么只能将 r 向左移动一位
//
//      由于题目保证必定有合法解，所以循环后面的代码不会运行到。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 numbers 中全部 O(n) 个数字
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量


impl Solution {
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        // 使用双指针找到合法解
        let (mut l, mut r) = (0, numbers.len() - 1);
        while l < r {
            // 计算当前左右指针指向的两个数的和
            let sum = numbers[l] + numbers[r];
            if sum == target {
                // 如果和为 target ，则 [l, r] 就是所求解
                return vec![l as i32 + 1, r as i32 + 1];
            }
            
            if sum < target {
                // 如果和小于 target ，则需要让和变大，只能右移 l
                l += 1;
            } else {
                // 如果和大于 target ，则需要让和变小，只能左移 r
                r -= 1;
            }
        }
        
        // 由于必定存在一个合法解，所以不会走到这
        unreachable!()
    }
}
