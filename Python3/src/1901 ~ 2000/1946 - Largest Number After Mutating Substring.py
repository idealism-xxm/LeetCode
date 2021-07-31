# 链接：https://leetcode.com/problems/largest-number-after-mutating-substring/
# 题意：给定一个数字字符串 num 和一个可替换的数位列表 list ，
#       现在可以替换其中一个子串，求替换后的最大数字字符串？


# 数据限制：
#   1 <= num.length <= 105
#   num 只由 0-9 组成
#   change.length == 10
#   0 <= change[d] <= 9

# 输入： num = "132", change = [9,8,5,0,3,6,4,2,6,8]
# 输出： 832
# 解释： 
#       替换子串 "1"
#       1 被换成 change[1] = 8
#
#       "132" -> "832"

# 输入： num = "021", change = [9,4,3,5,7,2,1,9,0,6]
# 输出： 934
# 解释： 
#       替换子串 "021"
#       0 被换成 change[0] = 9
#       2 被换成 change[2] = 3
#       1 被换成 change[1] = 4
#
#       "021" -> "934"

# 输入： num = "5", change = [1,4,7,5,3,2,5,6,9,4]
# 输出： 5


# 思路： 贪心
#
#       按照题意，我们最多替换一个子串，考虑空串的情况下，
#       最后会分为三部分：
#       第一部分：所有的数位都不小于替换后的数位
#       第二部分：从第一个替换后的数位变大的位置开始，到后续第一个替换后的数位变小的位置之前结束
#       第三部分：剩余的所有数位
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)

class Solution:
    def maximumNumber(self, num: str, change: List[int]) -> str:
        ans = ""
        i = 0
        # 先将所有数位都不小于替换后的数位的数放入 ans
        while i < len(num):
            digit = int(num[i])
            if digit < change[digit]:
                break

            ans += num[i]
            i += 1
        
        # 再将替换后可以变得不会更小的数位替换后放入 ans
        while i < len(num):
            digit = int(num[i])
            if digit > change[digit]:
                break

            ans += str(change[digit])
            i += 1
        
        # 最后再将剩余数位放入 ans
        return ans + num[i:]
