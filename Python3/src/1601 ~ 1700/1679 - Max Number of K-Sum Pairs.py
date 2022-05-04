# 链接：https://leetcode.com/problems/max-number-of-k-sum-pairs/
# 题意：给定一个整数数组 nums 和一个整数 k 。
#      每次操作可以将 nums 中和为 k 的两个数删除，
#      求最多能执行多少次这样的操作？


# 数据限制：
#  1 <= nums.length <= 10 ^ 5
#  1 <= nums[i] <= 10 ^ 9
#  1 <= k <= 10 ^ 9


# 输入： nums = [1,2,3,4], k = 5
# 输出： 2
# 解释： (1) 移除数字 1 和 4, nums 变为 [2,3]
#       (2) 移除数字 2 和 3, nums 变为 []
#       此时 nums 中没有和为 5 的两个数，所以最多操作 2 次

# 输入： nums = [3,1,3,4,3], k = 6
# 输出： 1
# 解释： (1) 移除数字 3 和 3, nums 变为 [1,4,3]
#       此时 nums 中没有和为 6 的两个数，所以最多操作 1 次


# 思路： Map
#
#      可以维护一个名为 num_to_cnt 的 map，
#      num_to_cnt[num] 表示 num 的可使用次数。
#
#      然后遍历 nums 数字中的数字 num ，
#      根据 k - num 的可使用次数处理：
#          1. num_to_cnt[num] > 0: 则可以执行一次操作，
#              移除 num 和一个 k - num 。
#              即 num_to_cnt[k - num] -= 1; and += 1;
#          2. num_to_cnt[num] == 0: 则不能执行操作，
#              令 num 的可使用次数加 1 ，
#              即 num_to_cnt[num] += 1;
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 nums 中全部 O(n) 个数字
#      空间复杂度：O(n)
#          1. 需要维护一个大小为 O(n) 的 map


class Solution:
    def maxOperations(self, nums: List[int], k: int) -> int:
        # num_to_cnt[num] 表示 num 还可以使用的次数
        num_to_cnt: Dict[int, int] = defaultdict(int)
        # ans 表示操作次数
        ans: int = 0
        # 遍历 nums 中的所有数字
        for num in nums:
            if num_to_cnt[k - num] > 0:
                # 如果 k - num 的可使用次数大于 0 ，
                # 则可以执行一次操作，移除 num 和一个 k - num 
                ans += 1
                # k - num 可使用次数减 1
                num_to_cnt[k - num] -= 1
            else:
                # 如果可使用次数等于 0 ，则不能操作，
                # 令 num 的可使用次数加 1
                num_to_cnt[num] += 1

        return ans
