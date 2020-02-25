// 链接：https://leetcode.com/problems/first-missing-positive/
// 题意：给定一个乱序的整数数组，找到最小的没有出现的正整数，要求时间复杂度为 O(n) 且 仅开辟常数空间

// 输入：[1,2,0]
// 输出：3

// 输入：[3,4,-1,1]
// 输出：2

// 思路：模拟，刚开始没想出来，不过基本想到了如何处理，所以看了提示1就知道怎么做了
//      由于题目要求时间复杂度为 O(n)，所以不能排序，只能遍历
//      由于只能开辟常数空间，所以不能用 map
//      但其实传入的数组本身就可以当map
//      1. 若 i == nums[i] - 1，则表明该数在正确的位置，不需要移动
//      2. 若 nums[i] <= 0 || nums[i] > len(nums)，则不进行处理，必定不会影响结果
//      3. 若 i != nums[i] - 1 && 1 <= nums[i] <= len(nums)，则移动到对应的位置，并且不断循环移动该位置的数字
//      最后再遍历一次，第一个 i != nums[i] - 1 的位置就是答案，若不存在这样的位置，则答案是 len(nums) + 1
//      时间复杂度：O(n)

func firstMissingPositive(nums []int) int {
    length := len(nums)
    if length == 0 {
        return 1
    }

    for i := 0; i < length; i++ {
        curNum := nums[i] // 保存当前位置的数字，准备开始移动
        for 1 <= curNum && curNum <= length && nums[curNum - 1] != curNum { // 当前数字可以被移动时，不断循环移动当前位置的数字
            nextNum := nums[curNum - 1] // 保存指定位置的数，准备继续循环
            nums[curNum - 1] = curNum // 当前数字移动到指定位置
            curNum = nextNum // 赋值进行循环
        }
    }

    for i := 0; i < length; i++  {
        if i != nums[i] - 1 { // 第一个 i != nums[i] - 1 的位置就是答案
            return i + 1
        }
    }
    return length + 1 // 不存在这样的位置，则答案是 length + 1
}