# 链接：https://leetcode.com/problems/remove-duplicates-from-sorted-array/
# 题意：给定一个非降序排序的整数数组 nums ，原地移除重复的数字，
#      使每个数字只出现一次，并保持相对顺序不变。
#      将这些不重复的 k 个数字放在 nums 的前 k 个位置，并返回 k 。
#
#      进阶：直接修改 nums ，使用空间复杂度为 O(1) 的算法。


# 数据限制：
#  1 <= nums.length <= 3 * 10 ^ 4
#  -100 <= nums[i] <= 100
#  nums 是非降序排序的


# 输入： nums = [1,1,2]
# 输出： 2, nums = [1,2,_]
# 解释： 总有 2 个不重复的数字，所以数字前 2 个位置的数分别为 1,2 。
#       后续位置的数字无论是什么都不影响答案的正确性。

# 输入： nums = [0,0,1,1,1,2,2,3,3,4]
# 输出： 5, nums = [0,1,2,3,4,_,_,_,_,_]
# 解释： 总有 5 个不重复的数字，所以数字前 5 个位置的数分别为 0,1,2,3,4 。
#       后续位置的数字无论是什么都不影响答案的正确性。


# 思路： 双指针
#
#      我们维护两个指针 l 和 r 。
#
#      l 表示不重复的数字个数，也是下一个可以放入数字的下标。
#      初始化为 1 ，表示第 1 个数必定不是重复的，方便后续处理。
#
#      r 表示下一个待考虑的数字的下标，初始化为 1 ，方便后续处理。
#
#      不断右移 r ，如果 nums[r] != nums[r - 1] ，
#      则 nums[r] 是不重复的，将其放到 l 处，并右移 l 。
#
#
#      时间复杂度： O(n)
#          1. 需要遍历 nums 中全部 O(n) 个数字
#      空间复杂度： O(1)
#          1. 只需要使用常数个额外变量即可


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        # l 表示不重复的数字个数，也是下一个可以放入数字的下标。
        # 初始化为 1 ，表示第 1 个数必定不是重复的
        l: int = 1
        # 遍历剩余所有的数字
        for r in range(1, len(nums)):
            # 如果当前数字不等于前一个数字，
            # 则 nums[r] 是不重复的，放入 l 处
            if nums[r] != nums[r - 1]:
                nums[l] = nums[r]
                l += 1

        return l
