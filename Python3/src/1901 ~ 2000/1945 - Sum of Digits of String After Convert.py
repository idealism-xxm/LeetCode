# 链接：https://leetcode.com/problems/sum-of-digits-of-string-after-convert/
# 题意：给定一个字符串 s ，先将其转换成数字后拼接在一起，
#       然后计算所有数位的和，重复这样的过程 k 次，返回最后的数？


# 数据限制：
#   1 <= s.length <= 100
#   1 <= k <= 10
#   s 只由英文小写字母组成

# 输入： s = "iiii", k = 1
# 输出： 36
# 解释： 
#       转换： "iiii" ➝ "(9)(9)(9)(9)" ➝ "9999" ➝ 9999
#       计算： #1: 9999 ➝ 9 + 9 + 9 + 9 ➝ 36

# 输入： s = "leetcode", k = 2
# 输出： 6
# 解释： 
#       转换： "leetcode" ➝ "(12)(5)(5)(20)(3)(15)(4)(5)" ➝ "12552031545" ➝ 12552031545
#       计算： #1: 12552031545 ➝ 1 + 2 + 5 + 5 + 2 + 0 + 3 + 1 + 5 + 4 + 5 ➝ 33
#             #2: 33 ➝ 3 + 3 = 6

# 输入： s = "zbax", k = 2
# 输出： 2


# 思路： 模拟
#
#       按照题意计算所数位之和 k 次即可
#
#       时间复杂度： O(kn)
#       空间复杂度： O(n)

class Solution:
    def getLucky(self, s: str, k: int) -> int:
        s = "".join([str(string.ascii_lowercase.index(ch) + 1) for ch in s])
        while k > 0:
            k -= 1
            s = str(sum([int(ch) for ch in s]))
        return int(s)
