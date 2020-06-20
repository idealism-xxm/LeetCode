// 链接：https://leetcode.com/problems/summary-ranges/
// 题意：给定一个升序无重复的整数数组，返回对应的每一段连续范围？

// 输入： [0,1,2,4,5,7]
// 输出： ["0->2","4->5","7"]
// 解释： 0, 1, 2 是一段连续范围， 4, 5 是一段连续范围

// 输入： [0,2,3,4,6,8,9]
// 输出： ["0","2->4","6","8->9"]
// 解释： 2, 3, 4 是一段连续范围， 8, 9 是一段连续范围

// 思路： 双指针
//
//      维护两个指针 l 和 r ，表示 [l, r] 内的数是连续的，
//      每次不断移动 r 直到数组末尾或者 r + 1 不连续，
//      收集此时的范围 [l, r] ，然后更新 l = r + 1; r += 1;
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

impl Solution {
    pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {
        let mut result: Vec<String> = Vec::new();
        // 左右指针，指向当前连续范围的起止下标
        let mut l = 0;
        let mut r = 0;
        while r < nums.len() {
            // 不断让 r 右移，直到达到末尾或者下一个数不连续
            while r + 1 < nums.len() && nums[r + 1] == nums[r] + 1 {
                r += 1
            }

            // 如果只有一个数，直接放入
            if l == r {
                result.push(format!("{}", nums[l]))
            } else {
                result.push(format!("{}->{}", nums[l], nums[r]))
            }
            // 处理下一段范围
            l = r + 1;
            r += 1;
        }

        result
    }
}
