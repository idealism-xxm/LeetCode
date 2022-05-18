# 链接：https://leetcode.com/problems/maximum-consecutive-floors-without-special-floors/
# 题意：给定两个数 bottom 和 top ，表示可选区间 [bottom, top] ，
#      以及一个整数数组 special ，其中 bottom <= special[i] <= top ，
#      special[i] 表示这个数不能选，求能选择的最大连续区间长度？


# 数据限制：
#  1 <= special.length <= 10 ^ 5
#  1 <= bottom <= special[i] <= top <= 10 ^ 9
#  special 中的所有数都是唯一的


# 输入： bottom = 2, top = 9, special = [4,6]
# 输出： 3
# 解释： 总共有三个可选区间： [2,3], [5,5], [7,9] ，
#       最大连续区间是 [7,9] ，长度为 3

# 输入： bottom = 6, top = 8, special = [7,6,8]
# 输出： 0
# 解释： 所有数都不可选，最大连续区间长度为 0


# 思路： 排序
#
#      那么我们可以对 special 按升序排序，由于 special 中是不可选的数字，
#      那么 [special[i - 1] + 1, special[i] - 1] 就是一个可选的连续区间，
#      长度为 special[i] - special[i - 1] - 1 ，
#      这些连续区间长度的最大值就是答案。
#
#      但这样没有考虑到可选区间的边界，所以为了方便处理，
#      可以先加入两个哨兵数字 bottom - 1 和 top + 1 ，
#      然后再按照上面的方法处理即可。
#
#
#      时间复杂度：O(nlogn)
#          1. 需要对数组进行排序，时间复杂度为 O(nlogn)
#          2. 需要遍历数组全部 O(n) 个数字
#      空间复杂度：O(n)
#          1. 需要加入哨兵数字 bottom - 1 和 top + 1 ，
#              在不修改入参的情况下需要 O(n) 的空间
#          2. （当然不使用哨兵数字时，可以优化为 O(1) ，但不方便处理）


class Solution:
    def maxConsecutive(self, bottom: int, top: int, special: List[int]) -> int:
        # 加入哨兵数字
        special.append(bottom - 1)
        special.append(top + 1)
        # 按升序排序
        special.sort()
        # ans 表示能选择的最大连续区间长度
        ans: int = 0
        # 遍历所有不能选择的数字作为区间端点
        for i in range(1, len(special)):
            # 更新能选择的连续区间长度的最大值
            ans = max(ans, special[i] - special[i - 1] - 1)

        return ans
