# 链接：https://leetcode.com/problems/longest-subsequence-repeated-k-times/
# 题意：给定一个字符串 s 和一个整数 k，找到 s 中最长的子序列 seq ，
#       且这个 seq 重复 k 次后仍能在 s 中找到对应的子序列，
#       如果长度相同，则返回字典序最大的。

# 数据限制：
#   n == s.length
#   2 <= k <= 2000
#   2 <= n < k * 8
#   s 由英文小写字母组成

# 输入： s = "letsleetcode", k = 2
# 输出： "let"
# 解释：
#   "let" 和 "ete" 都是最长的子序列，且重复 2 次后仍能在 s 中找到对应的子序列。
#   但 "let" 的字典序时最大的，因此返回 "let" 。


# 思路： DFS
#
#       比赛时还以为是什么特殊算法，注意到这个子序列最长是 7 个字符，但没有往暴力的地方想
#
#       比赛后看到可以直接枚举这个子序列，才恍然大悟，
#       还是要通过计算时间复杂度来推测，而不是无脑排除暴力
#
#       从最长长度的最大字典序开始枚举子序列，找到第一个满足要求的即可
#
#       时间复杂度： O(floor(n / k)! * n)
#       空间复杂度： O(n * floor(n / k))


class Solution:
    def longestSubsequenceRepeatedK(self, s: str, k: int) -> str:
        # 统计每个字母出现的次数
        cnt = defaultdict(int)
        for ch in s:
            cnt[ch] += 1
        
        # 找到出现次数 >= k 的字母，如果出现次数 >= 2k ，则需要放入两次
        chars = [ch * (freq // k) for ch, freq in cnt.items() if freq >= k]
        # 为了优先找到字典序大的，还需要降序排序
        chs = "".join(sorted(chars, reverse=True))

        # 挂在当前实例上，方便处理
        self.s = s
        self.chs = chs
        self.used = [False] * len(self.chs)
        self.k = k
        # 初始化答案为空串，用满足题意的字符串不断更新即可
        self.ans = ""

        # dfs 枚举可能的答案
        self.dfs("")
        return self.ans

    def dfs(self, cur: str):
        # 如果 cur * k 不是 s 的子序列，则直接返回
        if len(cur) and not self.is_subsequence(cur * self.k):
            return
        
        # 如果 cur 的长度更大，则更新 ans
        # 由于我们是按照字典序降序枚举，所以某个长度第一次满足题时，则其字典序必定最大
        if len(cur) > len(self.ans):
            self.ans = cur

        # 枚举下一个字母
        for i in range(len(self.chs)):
            # 如果第 i 个字符没有被使用过，则可以用它
            if not self.used[i]:
                self.used[i] = True
                self.dfs(cur + self.chs[i])
                self.used[i] = False
            
    
    def is_subsequence(self, seq: str) -> bool:
        # 如果使用 iter 的话，能够将时间从 9000 ms 减少到 1500 ms
        #
        # s_iter = iter(self.s)
        # return all(c in s_iter for c in seq)

        i, j = 0, 0
        while i < len(seq):
            # 找到下一个和 seq[i] 相同的字符
            while j < len(self.s) and self.s[j] != seq[i]:
                j += 1
            # 如果遍历完 s 都找不到，则直接返回 False
            if j == len(self.s):
                return False
            
            # 下面继续处理下一个字符
            i += 1
            j += 1
        
        # 最后如果遍历完了 seq ，则其是 s 的子序列，返回 True
        return True
