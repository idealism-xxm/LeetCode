// 链接：https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// 题意：给定一个按升序排序的数组 nums ，求 target 在 nums 中开始和结束的下标？
//      如果 target 不在 nums 中，则返回 [-1, -1] 。
//
//      要求：使用时间复杂度为 O(logn) 的方法。


// 数据限制：
//  0 <= nums.length <= 10 ^ 5
//  -(10 ^ 9) <= nums[i] <= 10 ^ 9
//  nums 是非递减数组
//  -(10 ^ 9) <= target <= 10 ^ 9


// 输入： nums = [5,7,7,8,8,10], target = 8
// 输出： [3,4]

// 输入： nums = [5,7,7,8,8,10], target = 6
// 输出： [-1,-1]

// 输入： nums = [], target = 0
// 输出： [-1,-1]


// 思路： 二分
//
//      本题其实就是要求实现 C++ 中的 lower_bound 和 upper_bound 即可。
//
//      两个函数的差别只有比较时不同：
//          1. lower_bound: nums[mid] <  target，则目标值在 右边区间
//          2. upper_bound: nums[mid] <= target，则目标值在 右边区间
//
//      在 Rust 中可以直接使用 Vec 的 partition_point 这个函数，
//      用上面两个不同的判断逻辑就能实现 lower_bound 和 upper_bound 。
//
//      我们可以先调用 lower_bound 求出 start ：
//          1. nums[start] != target ，则 nums 不在 nums 中，直接返回 [-1, -1]
//          2. nums[start] == target, 则再调用 upper_bound 求出 end ，
//              然后返回 [start, end - 1] 。
//
//              当然可以无需实现和使用 upper_bound ，
//              第二次调用 lower_bound(target + 1) 求出 end 即可。
//
//      lower_bound 函数的实现就是 LeetCode 0035 这题，可以直接复用。
//
//
//      时间复杂度：O(logn)
//          1. 需要二分区间 [0, n - 1] 两次，时间复杂度为 O(logn)
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


impl Solution {
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        // 先找到 nums 中第一个大于等于 target 的元素的下标
        let mut start = Solution::lower_bound(&nums, target);
        // 如果对应的数不是 target ，则 target 不在 nums 中，直接返回 [-1, -1]
        if start as usize == nums.len() || nums[start as usize] != target {
            return vec![-1, -1];
        }
        // 再找到 nums 中第一个大于等于 target + 1 的元素的下标，
        // 那么 end - 1 就是 target 在 nums 中的最后一个位置
        let end = Solution::lower_bound(&nums, target + 1);

        vec![start, end - 1]
    }

    // lower_bound 求 nums 中第一个大于等于 target 的元素的下标
    fn lower_bound(nums: &Vec<i32>, target: i32) -> i32 {
        // 二分区间的左边界，初始化为 0
        let mut l = 0;
        // 二分区间的右边界，初始化为 nums.len() - 1
        let mut r = nums.len() as i32 - 1;
        // 当区间不为空时，继续二分
        // （注意这里取等号是因为我们的区间是左闭右闭区间，
        // 且收缩 r 时不取到 mid）
        while l <= r {
            // 计算区间中点下标
            let mid = (l + r) >> 1;
            if nums[mid as usize] < target {
                // 如果区间中点的值小于 target ，
                // 则第一个大于等于 target 的元素
                // 必定在右边区间 [mid + 1, r] 中
                l = mid + 1;
            } else {
                // 如果区间中点的值大于等于 target ，
                // 则第一个大于等于 target 的元素
                // 必定在左边区间 [l, mid - 1] 中
                // （如果此时 mid 指向的元素，
                //   就是 nums 中第一个大于等于 target 的元素，
                //   那么在最后 l == r 时，会更新为 l = l + 1 ，
                //   最终选择到此时的 mid ）
                r = mid - 1;
            }
        }

        // 此时 l 指向 nums 中
        // 第一个大于等于 target 的元素的下标
        l
    }
}
