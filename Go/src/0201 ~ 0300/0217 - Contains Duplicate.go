// 链接：https://leetcode.com/problems/contains-duplicate/
// 题意：给定一个数组 nums ，判断是否存在重复的数？


// 数据限制：
//  1 <= nums.length <= 10 ^ 5
//  -(10 ^ 9) <= nums[i] <= 10 ^ 9


// 输入： [1,2,3,1]
// 输出： true
// 解释： 含有重复的数字 1

// 输入： [1,2,3,4]
// 输出： false
// 解释： 不含重复的数字

// 输入： [1,1,1,3,3,4,3,2,4,2]
// 输出： true
// 解释： 含有重复的数字 1, 2, 3, 4


// 思路： Set/Map
//
//		将 nums 转成集合，判断其长度是否等于 nums 的长度即可。
//
//      如果长度相等，则不含重复的数字，返回 false ；
//      否则含有重复的数字，返回 true 。
//
//      时间复杂度： O(n)
//          1. 需要遍历 nums 中全部 O(n) 个数字
//      空间复杂度： O(n)
//          1. 需要用一个集合维护全部不同的数字，最差情况下有 O(n) 个


func containsDuplicate(nums []int) bool {
    // 将 nums 转成集合
    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }

    // 集合的长度不等于 nums 的长度时，才含有重复数字
    return len(numSet) != len(nums)
}
