// 链接：https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/
// 题意：给定一个非负整数数组 nums ，每次操作可以进行如下处理：
//          1. 选择一个小于等于 nums 中最小的非零数的正数 x
//          2. 从每个正数 nums[i] 中减去 x
//
//      求最少多少次操作可以使得所有数都为 0 ？


// 数据限制：
//  1 <= nums.length <= 100
//  0 <= nums[i] <= 100


// 输入： nums = [1,5,0,3,5]
// 输出： 3
// 解释： 第一次操作，选择 x = 1 。现在 nums = [0,4,0,2,4]
//       第一次操作，选择 x = 2 。现在 nums = [0,2,0,0,2]
//       第一次操作，选择 x = 2 。现在 nums = [0,0,0,0,0]

// 输入： nums = [0]
// 输出： 0
// 解释： nums 中的所有数都是 0


// 思路： 贪心
//
//      假设 nums 中最小的非零数为 y ，那么每次选择的数 x 必定等于 y ，
//      这样一次操作就能将所有的 y 变为 0 。
//
//      如果选择的数 x 小于 y ，则至少要两次操作才能将所有的 y 变为 0 ，
//      最终所需的操作次数会更多。
//
//      由于每次都需要从每个正数 nums[i] 中减去 x ，并不影响正数之间的差值，
//      所以所需的操作数就是所有不同的正数的个数。
//
//      那我们用一个集合维护所有不同的正数，最后返回这个集合的大小即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数字
//      空间复杂度：O(n)
//          1. 需要用集合维护所有不同的数，最差情况下有 O(n) 个


func minimumOperations(nums []int) int {
    // num_set 维护所有不同的正数
    numSet := make(map[int]bool)
    // 遍历 nums 中的每个数，将所有正数放入 num_set 中
    for _, num := range nums {
        if num != 0 {
            numSet[num] = true
        }
    }

    // 不同正数的个数就是所需的操作数
    return len(numSet)
}
