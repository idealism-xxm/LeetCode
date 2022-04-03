// 链接：https://leetcode.com/problems/next-permutation/
// 题意：给定一个整数数组的排列 nums ，将其修改成下一个排列。
//
//      下一个排列就是比当前排列大的最小排列，
//      如果不存在，就是最小的排列。
//      例如： [1,2,3] 的下一个排列是 [1,3,2]
//            [2,3,1] 的下一个排列是 [3,1,2]
//            [3,2,1] 的下一个排列是 [1,2,3]


// 数据限制：
//  1 <= nums.length <= 100
//  0 <= nums[i] <= 100


// 输入： nums = [1,2,3]
// 输出： [1,3,2]

// 输入： nums = [3,2,1]
// 输出： [1,2,3]

// 输入： nums = [1,1,5]
// 输出： [1,5,1]


// 思路： 一次迭代
//
//      我们可以发现如果一个排列是降序的，
//      那么它的下一个排列就是最小的排列，
//      而这个排列我们可以通过翻转当前排列获得。
//
//      从后往前找到第一个满足 nums[i] < nums[i+1] 的 i ，
//      此时 nums[i+1..] 是降序的，这部分已经是最大的排列了。
//
//      所以要得到 nums 的下一个排列，
//      就需要让 nums[i] 变成比现在的数大的最小的数，
//      那么找到 nums[i+1..] 中大于 nums[i] 的最小的数字 nums[j] ，
//      交换 nums[i] 和 nums[j] ，并且让 nums[i+1..] 变为升序。
//
//      因为 nums[j] > nums[i]，
//      且 nums[j-1] >= nums[j] > nums[j+1] ，
//      所以 nums[j-1] >= nums[j] > nums[i] >= nums[j+1] ，
//      那么交换后， nums[i+1..] 仍旧是降序的，
//      只需要将 nums[i+1..] 翻转即可变成升序。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数字
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


impl Solution {
    pub fn next_permutation(nums: &mut Vec<i32>) {
        // 找到第一个满足 nums[i] < nums[i+1] 的 i
        let mut i = nums.len() as i32 - 2;
        // 如果 i 还合法，并且 nums[i] >= nums[i+1] ，
        // 则继续寻找前一个 i
        while i >= 0 && nums[i as usize] >= nums[i as usize + 1] {
            i -= 1;
        }

        // 如果 i >= 0 ，
        // 则找到了第一个满足 nums[i] < nums[i+1] 的 i ，
        // 需要继续处理
        if i >= 0 {
            // 找到第一个满足 nums[j] > nums[i] 的 j
            let mut j = nums.len() - 1;
            // 如果 nums[j] <= nums[i] ，
            // 则继续寻找前一个 j
            while nums[j] <= nums[i as usize] {
                j -= 1;
            }
            // 此时 nums[j] 是大于 nums[i] 最小的数，
            // 交换 nums[i] 和 nums[j]
            nums.swap(i as usize, j);
        }

        // 翻转 nums[i+1..] ，将其变成升序
        nums[(i+1) as usize ..].reverse();
    }
}
