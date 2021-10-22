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


# 思路2： 滑动窗口
#
#       换个想法，这个问题其实就能转换得比较简单了，
#       是下面这个题目的简单版本：
#       https://leetcode.com/problems/longest-repeating-character-replacement/
#
#
#       我们思路 1 中二分时判断的时候用到了固定长度的滑动窗口，
#       如果我们使用一个动态的滑动窗口，那么可以在 O(n) 内得处结果。
#       这个动态的滑动窗口内要维护 'T' 或 'F' 的个数 <= k 。
#
#       我们维护滑动窗口内的长度 length ，以及其中每个字符的长度 cnt ，
#       每次我们先放入当前字符，
#       如果 length - max(cnt) > k ，则需要将滑动窗口最前面的字符移除，
#       然后更新 ans = max(ans, length)
#
#       但这样如果将 'T' 和 'F' 替换成所有英文字母的话，时间复杂度的常数会增大，
#       所以需要继续挖掘其中可以利用的隐藏关联。
#
#       1. 我们可以发现 ans 只会越来越大，那么我们每次可以无脑放入当前字符，
#       保证滑动窗口的长度不会减小，即每次放入字符后，最多只移除滑动窗口前面的一个字符，
#       这样不会使得结果更差，而且保证 length 就是 ans 。
#
#       2. 我们再观察什么时候需要移除滑动窗口前面的一个字符。
#           就是当 length - max(cnt) > k 时，由于现在 length 不会减小，
#           如果 max(cnt) 不变得比上次更大的话，那么还是必须移除滑动窗口前面的一个字符，
#           所以 max(cnt) 也是不会减小的，那么我们可以用 max_cnt 维护这个最大次数，
#           也不会使得结果更差。
#           
#           现在如何更新 max_cnt 呢？可以发现只有刚加入的那个字符可能成为出现次数最多的那个，
#           所以我们更新 max_cnt = max(max_cnt, cnt[ch]) 即可。
#
#       通过这两个优化，我们就能将时间复杂度的常数去除，每次在循环中如下处理即可：
#
#           (1) length + 1 - max_cnt <= k ：当前字符加入后，
#               除去最多的 max_cnt 个字符后，还有不超过 k 个字符，不用去除滑动窗口最前面的字符
#           (2) length + 1 - max_cnt > k ：当前字符加入后，
#               除去最多的 max_cnt 个字符后，有超过 k 个字符，需要把滑动窗口最前面的字符移除后
#
#
#       时间复杂度： O(n)
#       空间复杂度： O(t) ，其中 t 时字符的种类数量


class Solution:
    def maxConsecutiveAnswers(self, answerKey: str, k: int) -> int:
        # 记录滑动窗口的长度
        length = 0
        # 统计滑动窗口内每个字符出现的次数
        cnt = defaultdict(int)
        # 记录滑动窗口内最多的出现次数
        max_cnt = 0
        for i, ch in enumerate(answerKey):
            # ch 加入到滑动窗口中
            cnt[ch] += 1
            # 更新滑动窗口字符最多的出现次数，只有 ch 有可能增大 max_cnt
            max_cnt = max(max_cnt, cnt[ch])
            # 如果现在滑动窗口中，除去最多的 max_cnt 个字符后，
            # 还有不超过 k 个字符，不用去除滑动窗口最前面的字符，
            # 那么滑动窗口的大小可以增大为 length + 1
            if length + 1 - max_cnt <= k:
                length += 1
            else:
                # 否则，必须移除滑动窗口最前面的字符
                cnt[answerKey[i - length]] -= 1
        
        return length
