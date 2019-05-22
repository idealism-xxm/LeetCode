// 链接：https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// 题意：给定一个有序数组，求去重后的数组长度（并需要修改入参，使前面一部分为去重后的数组）

// 输入：[1,1,2]
// 输出：2

// 输入：[0,0,1,1,1,2,2,3,3,4]
// 输出：5

// 思路：快慢指针
//      从第二个数字开始，只有当前数字不等于前一个数字时，才计数并赋值

func removeDuplicates(nums []int) int {
    length := len(nums)
    if length <= 1 {
        return length
    }

    result := 1
    for l := 1; l < length; l++ {
        if nums[l] != nums[l - 1] {
            nums[result] = nums[l] 
            result++
        }
    }

    return result
}