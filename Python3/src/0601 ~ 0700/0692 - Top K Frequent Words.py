# 链接：https://leetcode.com/problems/top-k-frequent-words/
# 题意：给定一个字符串数组 words 和一个正整数 k ，返回 k 个出现次数最多的字符串。
#      结果按照出现次数降序排序，出现次数相同时按字典序升序排序。
#
#      进阶：使用时间复杂度为 O(nlogk) ，空间复杂度为 O(n) 的算法。


# 数据限制：
#  1 <= words.length <= 500
#  1 <= words[i].length <= 10
#  words[i] 仅由英文小写字母组成
#  1 <= k <= words 中不同字符串的数量


# 输入： words = ["i","love","leetcode","i","love","coding"], k = 2
# 输出： ["i","love"]
# 解释： "i" 和 "love" 是出现次数最多的两个，其中 "i" 字典序更小

# 输入： words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
# 输出： ["the","is","sunny","day"]
# 解释： "the", "is", "sunny", "day" 是出现次数最多的四个，出现次数分别为 4, 3, 2, 1


# 思路： Map + 优先队列/堆
#
#      先用 word_to_cnt 统计 wrods 不同字符串的出现次数。
#
#      然后用一个优先队列/堆 top 维护这些字符串及出现次数，
#      并保持堆最多有 k 个元素。
#
#      堆底元素是第 1 多的，堆顶顶元素是第 k 多的，
#      这样当堆内元素超过 k 个时，可以快速移除不满足题意的。
#
#
#      时间复杂度：O(nlogk)
#          1. 需要遍历 word 中的全部 O(n) 个字符串
#          2. 需要遍历 word_to_cnt 中全部不同的字符串，
#              最差情况下有 O(n) 个。
#              每次都需要进行至多两次堆操作，时间复杂度为 O(logk)
#      空间复杂度：O(n + k)
#          1. 需要维护 word_to_cnt 中全部不同的字符串，
#              最差情况下有 O(n) 个。
#          2. 需要维护 top 中全部 O(k) 个元素


class Solution:
    def topKFrequent(self, words: List[str], k: int) -> List[str]:
        # word_to_cnt 统计每个字符串的出现次数
        word_to_cnt: Counter = Counter(words)

        # top 维护前 k 个出现次数最多的字符串及其出现次数，
        # 堆顶元素是出现次数第 k 多的，方便移除不满足题意的元素
        top: List[int] = []
        for word, cnt in word_to_cnt.items():
            # 将当前字符串及其出现次数放入 top 中
            heapq.heappush(top, (cnt, ReverseStr(word)))
            # 保持 top 只有前 k 个出现次数最多的
            if len(top) > k:
                heapq.heappop(top)

        # 从 top 中取出所有字符串并收集成一个列表
        ans: List[int] = []
        while top:
            ans.append(heapq.heappop(top)[1].s)

        # ans 的顺序是次数少的在前，所以要反向
        ans.reverse()
        return ans


class ReverseStr:
    def __init__(self, s: str):
        self.s = s

    def __lt__(self, other) -> bool:
        # 反转字符串的大小关系，因为 heapq 默认是最小堆
        return self.s > other.s

    def __eq__(self, other) -> bool:
        return self.s == other.s
