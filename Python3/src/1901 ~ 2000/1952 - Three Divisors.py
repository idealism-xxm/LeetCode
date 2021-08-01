# 链接：https://leetcode.com/problems/three-divisors/
# 题意：给定一个数 n ，判断其是否仅能被三个数整除？


# 数据限制：
#   1 <= n <= 10 ^ 4

# 输入： n = 2
# 输出： false
# 解释： 2 只能被 1 和 2 整除

# 输入： n = 4
# 输出： true
# 解释： 4 能被 1, 2, 4 整除


# 思路： 模拟
#
#       比赛时开始想错了，结果 WA 了一次，
#       因为一个数肯定能被 1 和自身整除，所以第三个数一定是 sqrt(n) ，
#       但题目说的是仅有三个数能整除 n ，所以 sqrt(n) 必定是素数。
#
#       比赛时没想这么多，所以就枚举 [1, sqrt(n)] ，统计能整除 n 的数的个数 cnt ，
#       然后返回 cnt == 3 即可
#       这个时间复杂度是 O(sqrt(n))
#
#       我们还可以判断 sqrt(n) 是否为素数，这样时间复杂度就降为 O(n ^ (1/4))
#       判断是否为素数的方法可以枚举 [2, sqrt(sqrt(n))]
#
#       时间复杂度： O(n ^ (1/4))
#       空间复杂度： O(1)

class Solution:
    def isThree(self, n: int) -> bool:
        if n < 4:
            return False

        root = int(math.sqrt(n))
        # 如果 n 不是平方数，则必定不满足题意
        if root * root != n:
            return False

        # 判断 root 是否为素数
        for i in range(2, int(math.sqrt(root)) + 1):
            # 如果能整除一个数，就不是素数，直接返回 False
            if root % i == 0:
                return False

        return True
