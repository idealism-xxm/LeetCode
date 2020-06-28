// 链接：https://leetcode-cn.com/problems/product-of-array-except-self/
// 题意：给定一个整数数组，求每个位置除了当前数字之外数字的乘积？

// 输入： [1,2,3,4]
// 输出： [24,12,8,6]

// 思路： 双指针
//
//      很容易就可以想到用两个数组先分别计算左起乘积 left 和右起乘积 right，
//      然后再计算当前位置的乘积 result[i] = left[i] * right[i]
//
//      空间复杂度为常数的还是没想到，看了题解才发现用双指针真得太巧妙了，
//      上面用到两个数组是因为我们将计算左（右）起乘积与计算结果分开了，
//      如果我们只计算右起乘积 right ，左起乘积跟着结果一起计算，
//          那么就可以优化为只需要一个额外数组了
//      可以发现 result[i] = left[i] * right[i] 的结果只与两个值有关，
//          那么我们转换一下思路，不要通过 i 去找对应的 left 和 right ，
//          而是通过 left[l] 和 right[r] 去找对应的 result ，
//          这样就可以在计算 left 和 right 的时候直接计算对 result 的贡献了
//      初始化 result[1~n] = 1 ，维护左右两个指针 l 和 r ，
//          以及对应的左起乘积 left 和 右起乘积 right ，
//          这样每次计算 left 和 right 前先将其算入对应的 result 中即可
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut result = vec![1; nums.len()];
        let mut left = 1;
        let mut right = 1;
        for l in 0..nums.len() {
            // 计算右指针
            let r = nums.len() - 1 - l;

            // 走到 l 时，左起乘积已算好，计入它对 result[l] 的贡献
            result[l] *= left;
            // 走到 r 时，右起乘积已算好，计入它对 result[r] 的贡献
            result[r] *= right;

            // 计算 l 位置的左起乘积
            left *= nums[l];
            // 计算 r 位置的右起乘积
            right *= nums[r];
        }

        result
    }
}
