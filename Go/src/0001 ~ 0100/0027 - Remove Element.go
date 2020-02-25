// 链接：https://leetcode.com/problems/remove-element/
// 题意：给定一个数组 和 一个指定数，求移除数组中所有指定数后的数组长度（并需要修改入参，使前面一部分为移除后的数组）

// 输入：nums = [0,1,2,2,3,0,4,2], val = 2
// 输出：5

// 思路：快慢指针
//      思想和上一题一样，只不过比较的数字变成指定的
//      从第一个数字开始，只有当前数字不等于指定时，才计数并赋值

func removeElement(nums []int, val int) int {
    length := len(nums)
    if length <= 0 {
        return length
    }

    result := 0
    for l := 0; l < length; l++ {
        if nums[l] != val {
            nums[result] = nums[l]
            result++
        }
    }

    return result
}