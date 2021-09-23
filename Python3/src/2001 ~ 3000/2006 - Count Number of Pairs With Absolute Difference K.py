# 链接：https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/
# 题意：给定一个整数数组 nums ，求有多少对 (i, j) 满足 |nums[i] - nums[j]| <= k ，其中 i < j 。

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
#       最简单就是使用 O(n ^ 2) 的枚举判断，但 k 是一个定值，
#       所以我们可以使用一个 map 记录，然后每次遍历的时候直接计算即可
#
#       时间复杂度： O(n)
#       空间复杂度： O(C), 其中 C 表示整型数字的范围


class Solution:
    def countKDifference(self, nums: List[int], k: int) -> int:
        # 统计每个数字出现的次数
        cnt = defaultdict(int)
        ans = 0
        for num in nums:
            # num - k 和 num + k 都能与 num 形成差为 k 的数对
            ans += cnt[num - k] + cnt[num + k]
            # 计入 num 出现次数
            cnt[num] += 1
        return ans
