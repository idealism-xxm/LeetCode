# 链接：https://leetcode.com/problems/most-frequent-even-element/
# 题意：给定一个整数数组 nums ，返回出现次数最多的偶数。
#      如果这样的偶数有多个，返回最小的那个。
#      如果不存在这样的偶数，返回 -1 。


# 数据限制：
#  1 <= nums.length <= 2000
#  0 <= nums[i] <= 10 ^ 5


# 输入： nums = [0,1,2,2,4,4,1]
# 输出： 2
# 解释： 偶数有 0, 2, 4 ，出现次数分别为 1, 2, 2 。
#       其中 2 和 4 出现次数最多，均为两次，返回最小的偶数 2 。

# 输入： nums = [4,4,4,9,2,4]
# 输出： 4
# 解释： 偶数有 4, 2 ，出现次数分别为 4, 1 。
#       其中 4 出现次数最多，返回 4 即可。

# 输入： nums = [29,47,21,41,13,37,25,7]
# 输出： -1
# 解释： 没有偶数，返回 -1 。


# 思路： Map
#
#      用一个名为 even_to_cnt 的 map 统计每个偶数出现的次数。
#
#      然后遍历 even_to_cnt 中的每个数字 num ，找到其中出现次数最多偶数 ans 即可。
#
#      【注意】如果出现次数相同，但遇到了更小的偶数，也需要更新 ans 。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 nums 中全部 O(n) 个数字
#          2. 需要遍历 even_to_cnt 中全部不同的偶数，最差情况下有 O(n) 个
#      空间复杂度：O(n)
#          1. 需要维护全部不同偶数的出现次数，最差情况下有 O(n) 个


class Solution:
    def mostFrequentEven(self, nums: List[int]) -> int:
        # 遍历 nums 数字，只统计偶数出现的次数
        even_to_cnt: Counter = Counter(num for num in nums if num & 1 == 0)

        # ans 表示出现次数最多的偶数，初始化为 -1 ，表示暂无偶数。
        # max_cnt 表示 ans 的出现次数，初始化为 0 ，表示暂无偶数。
        ans, max_cnt = -1, 0
        for even, cnt in even_to_cnt.items():
            if max_cnt < cnt:
                # 如果当前偶数 even 的出现次数更多，那么直接更新 ans 和 max_cnt
                ans = even
                max_cnt = cnt
            elif max_cnt == cnt and even < ans:
                # 如果当前偶数 even 的出现次数相同，但其更小，则更新 ans
                ans = even

        return ans
