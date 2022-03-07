# 链接：https://leetcode.com/problems/count-all-valid-pickup-and-delivery-options/
# 题意：给定 n 笔快递订单，每笔订单都需要先收件，再配送，
#      求总共有多少种收件/配送序列的方案数？


# 数据限制：
#  1 <= n <= 500


# 输入： n = 1
# 输出： 1
# 解释： 只有一种序列 (P1, D1) ，
#       订单 1 的配送（D1）在订单 1 的收件（P1）后。

# 输入： n = 2
# 输出： 6
# 解释： 所有可能的序列包括：
#      (P1,P2,D1,D2), (P1,P2,D2,D1), (P1,D1,P2,D2),
#      (P2,P1,D1,D2), (P2,P1,D2,D1), (P2,D2,P1,D1).
#
#      (P1,D2,P2,D1) 是一个无效的序列，
#      因为订单 2 的收件（P2）不应在订单 2 的配送（D2）之后。

# 输入： n = 3
# 输出： 90


# 思路： 排列组合
#
#		第一反应就是高中的排列组合：
#       最开始，我们的方案序列中有 2 * n 个位置可选，
#       每次选取其中的 2 个位置用于放置订单的收件 P 和配送 D 。
#   
#       我们按顺序枚举第 i 笔订单，前 [0, i) 个订单此时已经放置好了，
#       那现在方案序列中还有 remain_position = 2 * (n - i) 个位置可选，
#       需要从中选取两个位置，其中靠前的位置必定放 P_i ，靠后的位置必定放 D_i 。
#       方案数为： C(remain_position, 2)
#               = (remain_position * (remain_position - 1)) / 2
#   
#       那么总方案数就是所有这些订单的方案数之积，因为不同订单的方案相互独立。
#   
#
#       时间复杂度：O(n)
#           1. 需要枚举全部 O(n) 笔订单
#           2. 每笔订单的方案都能在 O(1) 内算出
#       空间复杂度：O(1)
#           1. 只需要使用常数个额外的变量

MOD: int = 1000000007

class Solution:
    def countOrders(self, n: int) -> int:
        # ans 统计所有可能的方案数，因为每一步的方案都相互独立，
        # 所以后续都是乘法，使用乘法单位元 1
        ans: int = 1
        for i in range(n):
            # 现在选择第 i 笔订单的收件 P_i 和配送 D_i 的位置，
            # 还剩 2 * (n - i) 个位置可选
            remain_position = 2 * (n - i)
            # 需要从中选择 2 个位置，
            # 总共有 cur = C(remain_position, 2) 种选法，
            # 因为 P_i 必须在 D_i 之前，所以选定 2 个位置后，
            # P_i 必定在前一个位置， D_i 必定在后一个位置
            cur = remain_position * (remain_position - 1) // 2
            # 第 i 笔订单的方案数 cur 与已有的方案数 ans 相互独立，
            # 可以再进行组合，所以要用乘法
            ans = (ans * cur) % MOD

        return ans
