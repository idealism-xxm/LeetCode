// 链接：https://leetcode.com/problems/number-of-subarrays-with-gcd-equal-to-k/
// 题意：给定一个整数数组 nums 和一个整数 k ，求有多少个子数组的最大公约数为 k ？


// 数据限制：
//  1 <= nums.length <= 1000
//  1 <= nums[i], k <= 10 ^ 9


// 输入： nums = [9,3,1,2,6,3], k = 3
// 输出： 4
// 解释： 以下 4 个子数组的最大公约数为 3 ：
//          - [9,(3),1,2,6,3]
//          - [9,3,1,2,6,(3)]
//          - [(9,3),1,2,6,3]
//          - [9,3,1,2,(6,3)]

// 输入： nums = [4], k = 7
// 输出： 0
// 解释： 任何子数组的最大公约数都不是 7


// 思路1： 模拟
//
//      我们可以直接枚举全部 O(n ^ 2) 个子数组，
//      统计最大公约数为 k 的子数组即可。
//
//      对于子数组 nums[i..=j] 来说，如果已知其最大公约数为 g ，
//      那么子数组 nums[i..=j+1] 的最大公约数为 gcd(g, nums[j+1]) 。
//
//      这样我们在枚举子数组的时候，就能直接递推出其最大公约数，
//      无需遍历子数组全部数字。
//
//
//      设 nums 中数的最大值为 M 。
//
//      时间复杂度：O(n ^ 2 * logM)
//          1. 需要枚举全部 O(n ^ 2) 个子数组，
//              每次递推只需要求一次最大公约数，辗转相除法的时间复杂度为 O(logM)
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


func subarrayGCD(nums []int, k int) int {
    // ans 维护满足题意的子数组的个数
    ans := 0
    for i := range nums {
        // g 维护子数组的最大公约数
        g := nums[i]
        for j := i; j < len(nums); j++ {
            // 递推出子数组 nums[i..=j] 的最大公约数
            g = gcd(g, nums[j])
            // 如果子数组 g 等于 k ，则计入答案中
            if g == k {
                ans += 1
            }
        }
    }

    return ans
}

// 辗转相除法计算最大公约数
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b;
    }

    return a
}
