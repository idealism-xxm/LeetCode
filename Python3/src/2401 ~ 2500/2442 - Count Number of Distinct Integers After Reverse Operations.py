# 链接：https://leetcode.com/problems/count-number-of-distinct-integers-after-reverse-operations/
# 题意：给定一个正整数数组 nums ，将每个数按十进制位翻转后再放入 nums 中，
#      求最终数组中有多少个不同的数？


# 数据限制：
#  1 <= nums.length <= 10 ^ 5
#  1 <= nums[i] <= 10 ^ 6


# 输入： nums = [1,13,10,12,31]
# 输出： 6
# 解释： nums 最终为 [1,13,10,12,31,1,31,1,21,13] ，
#       其中不同的数共 6 个： 1, 10, 12, 13, 21, 31

# 输入： nums = [2,2,2]
# 输出： nums 最终为 [2,2,2,2,2,2] ，
#       其中不同的数共 1 个： 2


# 思路： Set/Map
#
#      用一个集合维护最终数组中的全部数字。
#
#      然后遍历 nums 中每个数字 num ，先将 num 放入集合中，
#      再将 num 按十进制位翻转后放入集合中。
#
#      最后返回集合的长度即可。
#
#
#      时间复杂度：O(nlogn)
#          1. 需要遍历 nums 中全部 O(n) 个数，每次都要遍历全部 O(logn) 个十进制位
#      空间复杂度：O(n)
#          1. 需要维护最终数组中全部不同的数，最差情况下有 O(n) 个


class Solution:
    def countDistinctIntegers(self, nums: List[int]) -> int:
        # 用集合 num_set 维护 nums 中全部数字
        num_set: Set[int] = set()
        for num in nums:
            # num 必定在最终数组中
            num_set.add(num)
            # 将 num 按十进制位翻转后放入集合中
            res: int = 0
            while num > 0:
                res = res * 10 + num % 10
                num //= 10
            num_set.add(res)

        return len(num_set)
