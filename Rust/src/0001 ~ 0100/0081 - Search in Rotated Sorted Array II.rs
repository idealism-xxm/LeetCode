// 链接：https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
// 题意：给定一个升序的有重复数字的整型数组 nums ，
//      将后面一部分（不清楚有多少数）放到前面，
//		判断指定的数 target 是否在数组内？


// 数据限制：
//  1 <= nums.length <= 5000
//  -(10 ^ 4) <= nums[i] <= 10 ^ 4
//  nums 确保在某个点旋转
//  -(10 ^ 4) <= target <= 10 ^ 4


// 输入： nums = [2,5,6,0,0,1,2], target = 0
// 输出： true
// 解释： 0 在数组 nums 中

// 输入： nums = [2,5,6,0,0,1,2], target = 3
// 输出： false
// 解释： 3 不在数组 nums 中


// 思路： 二分
//
//      大部分情况下和 LeetCode - 0033 是一样的，
//      由于多了重复数字，会存在无法判断的情况，
//      所以最差情况下的时间复杂度为 O(n) 。
//
//      在每次确定起点（最小值）的区间后，判断 target 所在的区间即可：
//          1. nums[l] == nums[mid] == nums[r]: 
//              无法区分起点（最小值）的区间，也无法知道 target 在哪个区间，
//              只能排除 nums[l] 和 nums[r] 均不为 target 这种情况，
//              令 l = l + 1, r = r - 1
//
//          2. nums[l] <= nums[mid]: 左边区间是连续的，
//              所以可以判断 target 是否在 [num[l], num[mid]] 内
//
//              (1) num[l] <= target <= nums[mid]:
//                  target 在左边区间数的范围内，令 r = mid - 1
//              (2) 其他情况，认为目标数在右边区间，令 l = mid + 1
//
//          3. nums[mid] <= nums[r]: 右边区间是连续的，
//              所以可以判断 target 是否在 [num[mid], num[r]] 内
//              （注意不可能存在 nums[l] > nums[mid] > nums[r] 这种情况）
//
//              (1) num[mid] <= target <= nums[r]:
//                  target 在右边区间数的范围内，令 l = mid + 1
//              (2) 其他情况，认为目标数在左边区间，令 r = mid - 1
//
//
//		时间复杂度： 最好 O(logn) | 最差 O(n)
//          1. 最好情况下，所有分支每次都会排除一半数字，
//              时间复杂度为 O(logn)
//          2. 最差情况下，所有数字都一样，并且 target 不在 nums 中，
//              时间复杂度为 O(n)
//		空间复杂度： O(1)
//          1. 只需要使用常数个额外变量


impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> bool {
        // 二分区间左边界
        let mut l = 0;
        // 二分区间右边界（注意必须使用 i32 类型，因为 r 最终可能小于 0 ）
        let mut r = nums.len() as i32 - 1;
        while l <= r {
            // 计算区间中点下标
            let mut mid = (l + r) >> 1;
            // 如果区间中点下标对应的数字等于 target ，则直接返回 true
            if nums[mid as usize] == target {
                return true;
            }
            if nums[l as usize] == nums[mid as usize] && nums[mid as usize] == nums[r as usize] {
                // 如果三个点的值都相等，无法区分起点（最小值）的区间，
                // 也无法知道 target 在哪个区间，
                // 只能排除 nums[l] 和 nums[r] 均不为 target 这种情况
                l += 1;
                r -= 1;
            } else if nums[l as usize] <= nums[mid as usize] {
                // 左边区间是连续的，
                // 所以可以判断 target 是否在 [num[l], num[mid]] 内
                if nums[l as usize] <= target && target <= nums[mid as usize] {
                    // target 在左边区间数的范围内，下次查找的区间为 [l, mid - 1]
                    r = mid - 1;
                } else {
                    // 其他情况，认为目标数在右边区间，下次查找的区间为 [mid + 1, r]
                    l = mid + 1;
                }
            } else {
                // 此时必定是 nums[mid] <= nums[r] ，
                // 因为不可能存在 nums[l] > nums[mid] > nums[r] 这种情况。
                //
                // 则右边区间是连续的，
                // 所以可以判断 target 是否在 [num[mid], num[r]] 内
                if nums[mid as usize] <= target && target <= nums[r as usize] {
                    // target 在右边区间数的范围内，下次查找的区间为 [mid + 1, r]
                    l = mid + 1;
                } else {
                    // 其他情况，认为目标数在左边区间，下次查找的区间为 [l, mid - 1]
                    r = mid - 1;
                }
            }
        }

        // 此时表明没有找到 target ，返回 false
        false
    }
}
