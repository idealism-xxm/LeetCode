// 链接：https://leetcode.com/problems/next-permutation/
// 题意：给定一个整数数组，求下一个排列（下一个排列指字典序比当前排列大的最小的排列）

// 输入：1,2,3
// 输出：1,3,2

// 输入：3,2,1
// 输出：1,2,3

// 思路1：模拟即可
//      从后往前找，每次记住最大的数 maxNum
//      找到第一个小于当前最大的数的下标 index
//      在从 index + 1 往后找比 nums[index] 大的最小数的下标 targetIndex
//      交换 nums[index] 和 nums[targetIndex]，然后对 nums[index + 1] 按升序排序
//      时间复杂度：O(nlogn)

import "sort"

func nextPermutation(nums []int)  {
    length := len(nums)
    maxNum := nums[length - 1]
    index := length - 2
    for ; index >= 0; index--  {
        if nums[index] < maxNum {
            break // 找到比最大数小的第一个数的下标
        }
        maxNum = nums[index] // 此时nums[index] >= maxNum，直接更新即可
    }
    if index != -1 { // 如果不是最大的排列，则找比 nums[index] 大的最小数并交换
        targetIndex := index + 1
        for i := targetIndex + 1; i < length; i++ {
            if nums[i] > nums[index] && nums[i] < nums[targetIndex] {
                targetIndex = i // 找比 nums[index] 大的最小数
            }
        }
        nums[index], nums[targetIndex] = nums[targetIndex], nums[index] // 交换
    }

    sort.Ints(nums[index + 1:]) // num[index + 1:] 按升序排序
}