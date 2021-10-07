# 链接：https://leetcode.com/problems/maximize-the-confusion-of-an-exam/
# 题意：给定一个只含有 'T' 和 'F' 的字符串 s ，
#       求最多进行 k 次如下操作后，最长连续的 'T' 或 'F' 子串的长度是多少？
#       操作：将 s 中的一个 'T' 换成 'F' ，或者将 'F' 换成 'T' 。

# 数据限制：
#   n == answerKey.length
#   1 <= n <= 5 * 10 ^ 4
#   answerKey[i] 是 'T' 或 'F'
#   1 <= k <= n


# 输入： answerKey = "TTFF", k = 2
# 输出： 4
# 解释： "TTFF" -> "TTTT"

# 输入： answerKey = "TFFT", k = 1
# 输出： 3
# 解释： "TFFT" -> "FFFT"

# 输入： answerKey = "TTFTTFTT", k = 1
# 输出： 5
# 解释： "TTFTTFTT" -> "TTTTTFTT"


# 思路1： 二分
#
#       我们二分最长长度 mid ，然后判断是否存在 i ，
#       使得 s[i:i + mid] 最多只含有 k 个 'T' 或 'F' ，
#       即 'T' 和 'F' 个数的较小值 <= k
#
#       如果存在，则下一次二分右半边： l = mid + 1
#       如果不存在，则下一次二分左半边： r = mid - 1
#
#       当 l <= r 时，继续上述二分，最终答案就是 r
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(1)


class Solution:
    def maxConsecutiveAnswers(self, answerKey: str, k: int) -> int:
        def check(length: int) -> bool:
            # 先统计 s[0:length] 中 'T' 的个数
            cnt_t = sum(1 for i in range(length) if answerKey[i] == 'T')
            # 如果 'T' 或者 'F' 的个数小于等于 k ，
            # 则长度为 legnth 的子串经过最多 k 次替换后，可以变成连续的 'T' 或 'F'
            if cnt_t <= k or length - cnt_t <= k:
                return True

            # 枚举长度为 legnth 的子串开始下标
            for i in range(1, len(answerKey) - length + 1):
                # 每次从 s[i - 1:i - 1 + length] 变到 s[i:i + length] 时，
                # 只用更新 s[i - 1] 和 s[i - 1 + length] 对 cnt_t 造成的影响即可
                if answerKey[i - 1] == 'T':
                    cnt_t -= 1
                if answerKey[i - 1 + length] == 'T':
                    cnt_t += 1
                # 如果 'T' 或者 'F' 的个数小于等于 k ，
                # 则长度为 legnth 的子串经过最多 k 次替换后，可以变成连续的 'T' 或 'F'
                if cnt_t <= k or length - cnt_t <= k:
                    return True
            
            # 所有起始位置长度为 length 的子串都不满足，则当前长度不行
            return False

        # 二分区间为 [1, n]
        l, r = 1, len(answerKey)
        while l <= r:
            # 判断 s[i:i + mid] 最多只含有 k 个 'T' 或 'F'
            mid = (l + r) >> 1
            if check(mid):
                # 下一次二分右半边
                l = mid + 1
            else:
                # 下一次二分左半边
                r = mid - 1
        
        return r
