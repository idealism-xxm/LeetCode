# 链接：https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/
# 题意：给定一个整数数组 nums ，求有多少对 (i, j) 满足 |nums[i] - nums[j]| == k ，其中 i < j 。


# 数据限制：
#   1 <= nums.length <= 200
#   1 <= nums[i] <= 100
#   1 <= k <= 99


# 输入： nums = [1,2,2,1], k = 1
# 输出： 4
# 解释：
#   [(1),(2),2,1]
#   [(1),2,(2),1]
#   [1,(2),2,(1)]
#   [1,2,(2),(1)]

# 输入： nums = [1,3], k = 3
# 输出： 0

# 输入： nums = [3,2,1,5,4], k = 2
# 输出： 3
# 解释：
#   [(3),2,(1),5,4]
#   [(3),2,1,(5),4]
#   [3,(2),1,5,(4)]


# 思路： Map
#
#       k 是一个定值，所以我们可以使用一个 map 记录每个数字出现的次数，
#       在遍历 nums 的时候计算满足题意的数对个数，并更新 map 。
#
#       遍历 nums 中的数字 num ，让其作为数对的 nums[j] ，
#       那么 nums[i] 的取值只有 num - k 和 num + k 。
#
#       将这两个数字的出现次数计入答案，然后更新 num 的出现次数。
#
#
#       时间复杂度： O(n)
#          1. 需要遍历 nums 中全部 O(n) 个数字
#       空间复杂度： O(n)
#          2. 需要维护 num_to_cnt 中全部不同的数字的出现次数，最差情况下有 O(n) 个


class Solution:
    def countKDifference(self, nums: List[int], k: int) -> int:
        # 统计每个数字出现的次数
        num_to_cnt: Dict[int, int] = defaultdict(int)
        # ans 维护所有满足题意的数对
        ans: int = 0
        for num in nums:
            # num - k 和 num + k 都能与 num 形成差为 k 的数对
            ans += num_to_cnt[num - k] + num_to_cnt[num + k]
            # 计入 num 出现次数
            num_to_cnt[num] += 1

        return ans
