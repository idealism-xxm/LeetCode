# 链接：https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
# 题意：有一个值 X 初始为 0 ，对其进行一些操作后，返回最后的值。
#       "X++" 和 "++X": X += 1
#       "X--" 和 "--X": X -= 1

# 数据限制：
#   1 <= operations.length <= 100
#   operations[i] 是 "++X", "X++", "--X", "X--" 四个之一

# 输入： operations = ["--X","X++","X++"]
# 输出： 1
# 解释：
#   初始： X = 0
#   --X: X =  0 - 1 = -1
#   X++: X = -1 + 1 =  0
#   X++: X =  0 + 1 =  1

# 输入： operations = ["++X","++X","X++"]
# 输出： 3
# 解释：
#   初始： X = 0
#   ++X: X = 0 + 1 = 1
#   ++X: X = 1 + 1 = 2
#   X++: X = 2 + 1 = 3

# 输入： operations = ["X++","++X","--X","X--"]
# 输出： 0
# 解释：
#   初始： X = 0
#   X++: X = 0 + 1 = 1
#   ++X: X = 1 + 1 = 2
#   --X: X = 2 + 1 = 1
#   X--: X = 1 - 1 = 0


# 思路： 模拟
#
#       按照题意对 X 操作即可
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def finalValueAfterOperations(self, operations: List[str]) -> int:
        ans = 0
        for ope in operations:
            # 如果是加法操作，则执行 +1 ，否则执行 -1
            if ope == "X++" or ope == "++X":
                ans += 1
            else:
                ans -= 1
        return ans
