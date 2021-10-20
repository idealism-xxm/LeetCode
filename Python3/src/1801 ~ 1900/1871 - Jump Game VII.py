# 链接：https://leetcode.com/problems/jump-game-vii/
# 题意：给定一个 01 串 s 和两个整数 minJump 和 maxJump ，
#       你最开始在 s[0] 处，且 s[0] 是 '0' ，
#       你可以从 s[i] 跳到 s[j] 处，只要满足以下条件即可：
#           1. i + minJump <= j <= (i + maxJump, len(s) - 1)
#           2. s[j] == '0'
#       判断最后是否能到达 len(s) - 1 处？

# 数据限制：
#   2 <= s.length <= 10 ^ 5
#   s[i] 是 '0' 或 '1'
#   s[0] == '0'
#   1 <= minJump <= maxJump < s.length

# 输入： s = "011010", minJump = 2, maxJump = 3
# 输出： true
# 解释：
#   "(0)11010" -> "011(0)10" -> "01101(0)"


# 输入： s = "01101110", minJump = 2, maxJump = 3
# 输出： false


# 思路： DP + 滑动窗口
#
#       如果 s[-1] == '1' ，则不满足题意，直接返回 False
#
#       我们用 dp[i] 表示从 i 能否到达 len(s) - 1 ，
#       初始化 dp[i] = False ， dp[-1] = True ，
#       这样我们可以从后往前更新。
#
#       对于位置 i 来说，它只能跳到 [i + minJump, i + maxJump] 内的 '0' 中，
#       由于这个范围的长度固定，每次只有 i 在改变，那么我们就可以用滑动窗口，
#       每次往前移动一个位置，并用 cnt 维护 dp[i + minJump, i + maxJump] 中的 True 的个数，
#       初始化 cnt = 0 ， l = len(s) - 1 + minJump ， r = len(s) - 1 + maxJump ，
#       每次循环先移动滑动窗口，然后更新 cnt ，
#           (1) 如果 l < len(s) ，说明 l 合法且更进入滑动窗口，
#               那么只要 dp[l] 是 True ，就需要对 cnt 加 1
#           (2) 如果 r < len(s) - 1 ，说明 r + 1 合法且更离开滑动窗口，
#               那么只要 dp[r + 1] 是 True ，就需要对 cnt 减 1
#       然后只要 cnt 大于 0 ，且 s[i] 是 '0' ，那么就可以从 i 跳到最后一个位置，
#       更新 dp[i] = True
#       
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def canReach(self, s: str, minJump: int, maxJump: int) -> bool:
        # 如果最后一个不是 '0' ，则必定不满足题意
        if s[-1] != '0':
            return False

        # dp[i] 表示从 i 能否到达 len(s) - 1
        dp = [False] * len(s)
        # s[-1] == '0' ，所以 dp[-1] = True
        dp[-1] = True
        # 滑动窗口的范围为 [i + minJump, i + maxJump]
        l, r = len(s) - 1 + minJump, len(s) - 1 + maxJump
        # 滑动窗口内 dp 是 True 的个数
        cnt = 0
        
        # 从后往前遍历
        for i in range(len(s) - 2, -1, -1):
            # 先移动左右边界
            l -= 1
            r -= 1
            # 如果左边界 l 合法，那么左边界 l 必定是刚刚进入的，
            # 只要 dp[l] 是 True ，滑动窗口内的 cnt 就需要 +1
            if l < len(s) and dp[l]:
                cnt += 1
            # 如果右边界 r 不是最后一个位置，
            # 那么其后的那个位置 r + 1 刚刚离开滑动窗口，
            # 只要 dp[r + 1] 是 True ，滑动窗口内的 cnt 就需要 -1
            if r < len(s) - 1 and dp[r + 1]:
                cnt -= 1
            
            # 如果此时 [l, r] 内还有可以抵达最后一个位置的位置，
            # 且 s[i] 是 '0' ，那么 i 也可以抵达最后一个位置
            if cnt and s[i] == '0':
                dp[i] = True

        return dp[0]
        