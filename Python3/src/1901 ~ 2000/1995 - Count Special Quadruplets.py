# 链接：https://leetcode.com/problems/count-special-quadruplets/
# 题意：给定一个整型数组 nums ，求有多少个四元组满足以下要求？
#       1. nums[a] + nums[b] + nums[c] == nums[d]
#       2. a < b < c < d

# 数据限制：
#   4 <= nums.length <= 50
#   1 <= nums[i] <= 100

# 输入： nums = [1,2,3,6]
# 输出： 1
# 解释：
#   - (0, 1, 2, 3): 1 + 2 + 3 == 6

# 输入： nums = [3,3,6,4,5]
# 输出： 0

# 输入： nums = [1,1,1,3,5]
# 输出： 4
# 解释：
#   - (0, 1, 2, 3): 1 + 1 + 1 == 3
#   - (0, 1, 3, 4): 1 + 1 + 3 == 5
#   - (0, 2, 3, 4): 1 + 1 + 3 == 5
#   - (1, 2, 3, 4): 1 + 1 + 3 == 5


# 思路： 枚举
#
#       最朴素的方法就是直接枚举四个数，然后判断是否满足条件即可
#       这样时间复杂度是 O(n ^ 4) ，空间复杂度是 O(1)
#
#       如果要降低时间复杂度，那么就需要额外空间存储
#       由于四个数是有序的，我们可以倒序枚举前两个数 a, b 计算结果，
#       并倒序枚举后两个数 (c, d) ，并将所有 nums[d] - nums[c] 的结果存储在计数 map 中，
#       这样在枚举 (a, b) 时就可以在 O(1) 内统计出满足要求的数总数
#
#       时间复杂度： O(n ^ 2)
#       空间复杂度： O(n ^ 2)


class Solution:
    def countQuadruplets(self, nums: List[int]) -> int:
        ans = 0
        # 计数 dict ，统计 nums[d] - nums[c] 出现的次数
        cnt = defaultdict(int)
        # 初始化先计算最后两个数的差值
        cnt[nums[-1] - nums[-2]] = 1
        # 枚举 a, b
        for b in range(len(nums) - 3, 0, -1):
            for a in range(b - 1, -1, -1):
                # 目前 cnt 已包含后面所有二元组 (c, d) 形成的 nums[d] - nums[c] 的统计结果
                ans += cnt[nums[a] + nums[b]]

            # 在下一轮中，此时的 b 就相当于需要统计的 c
            # 对于其他 c 来说，已经在前面的循环中处理过了
            for d in range(b + 1, len(nums)):
                cnt[nums[d] - nums[b]] += 1
        return ans
