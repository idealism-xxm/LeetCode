# 链接：https://leetcode.com/problems/find-greatest-common-divisor-of-array/
# 题意：给定一个整型数组，返回最大值和最小值的最大公约数。

# 数据限制：
#   2 <= nums.length <= 1000
#   1 <= nums[i] <= 1000

# 输入： nums = [2,5,6,9,10]
# 输出： 2
# 解释： 
#       最小值为 2 ，最大值为 10 ，所以它们的最大公约数是 2 。

# 输入： nums = [7,5,6,8,3]
# 输出： 1
# 解释： 
#       最小值为 3 ，最大值为 8 ，所以它们的最大公约数是 1 。

# 输入： nums = [3,3]
# 输出： 3
# 解释： 
#       最小值为 3 ，最大值为 3 ，所以它们的最大公约数是 3 。


# 思路： 模拟
#
#       找到数组中的最大值 mx 和最小值 mn
#
#       直接计算 mx 和 mn 的最大公约数 gcd 即可
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)

class Solution:
    def findGCD(self, nums: List[int]) -> int:
        mn, mx = 10000, 1
        for num in nums:
            mn = min(mn, num)
            mx = max(mx, num)
        return gcd(mn, mx)
