# 链接：https://leetcode.com/problems/check-if-an-original-string-exists-given-two-encoded-strings/
# 题意：一个仅有英文小写字母组成的原始字符串可以按照以下步骤编码：
#       1. 将这个字符串分成任意个非空子串
#       2. 选择任意个子串，并替换成这个子串的长度
#       3. 连接所有的子串，形成编码后的串
#      给定两个编码后的串 s1 和 s2 ，判断一个原始字符串能否编码成这两个串？

# 数据限制：
#   1 <= s1.length, s2.length <= 40
#   s1 和 s2 仅由数字 (1-9) 、小写英文字母组成
#   s1 和 s2 中连续的数字长度不超过 3

# 输入： s1 = "internationalization", s2 = "i18n"
# 输出： true
# 解释： 
#   "internationalization" -> ["internationalization"] -> ["internationalization"] -> "internationalization"
#   "internationalization" -> ["i", "nternationalizatio", "n"] -> ["i", "18", "n"] -> "i18n"

# 输入： s1 = "l123e", s2 = "44"
# 输出： true
# 解释： 
#   "l123e" -> ["l", "e", "et", "cod", "e"] -> ["l", "1", "2", "3", "e"] -> "l123e"
#   "l123e" -> ["leet", "code"] -> ["4", "4"] -> "44"


# 输入： s1 = "a5b", s2 = "c5b"
# 输出： false
# 解释： 
#   编码成 s1 的字符串必须以 'a' 开始
#   编码成 s2 的字符串必须以 'c' 开始

# 输入： s1 = "112s", s2 = "g841"
# 输出： true
# 解释： 
#   "gaaaaaaaaaaaas" -> ["g", "aaaaaaaaaaaa", "s"] -> ["1", "12", "s"] -> "l123e"
#   "gaaaaaaaaaaaas" -> ["g", "aaaaaaaa", "aaaa", "s"] -> ["g", "8", "4", "1"] -> "g841"

# 输入： s1 = "ab", s2 = "a2"
# 输出： false
# 解释： 
#   编码成 s1 的字符串长度为 2
#   编码成 s2 的字符串长度为 3


# 思路： DP
#
#       设 dp[i][j] 表示 s1[:i] 和 s2[:j] 的匹配时，原串可能的长度差集合。
#       初始化： dp[0][0] = {0}
#       状态转移：
#           遍历 dp[i][j] 中的每一个 d ，
#           1. s1[i:p] 都是数字： dp[p][j] 中含有长度差 d + int(s1[i:p])
#           2. s2[j:q] 都是数字： dp[i][q] 中含有长度差 d - int(s2[j:q])
#           3. s1[i] 是字母且 d < 0 ： dp[i+1][j] 中含有长度差 d + 1
#               （此时 s1[i] 可以被 s2 中多余的长度消耗）
#           4. s2[j] 是字母且 d > 0 ： dp[i][j+1] 中含有长度差 d - 1
#               （此时 s2[j] 可以被 s1 中多余的长度消耗）
#           5. s1[i] == s2[j] 且都是字母且长度差 d == 0 ： dp[i+1][j+1] 中含有长度差 0
#
#       如果最后 dp[n][m] 中含有长度差 0 ，则表示 s1 和 s2 可以成功匹配
#
#       时间复杂度： O(n * m * D * 10 ^ D) ，其中 D = 3
#       空间复杂度： O(n * m * 10 ^ D)


class Solution:
    def possiblyEquals(self, s1: str, s2: str) -> bool:
        n, m = len(s1), len(s2)
        # 初始化 dp 数组
        dp = [[set() for _ in range(m + 1)] for _ in range(n + 1)]
        dp[0][0].add(0)

        # 遍历 s1 已匹配的字符长度
        for i in range(n + 1):
            # 遍历 s2 已匹配的字符长度
            for j in range(m + 1):
                # 遍历 dp[i][j] 中的每一个长度差 d
                for d in dp[i][j]:
                    # 情况 1 ： s1[i:p] 都是数字时，状态转移至 dp[p][j]
                    num = 0
                    for p in range(i, n):
                        # 第一个非数字时跳出循环
                        if not s1[p].isdigit():
                            break

                        # 将 s1[i:p + 1] 转换为数字 num
                        num = num * 10 + int(s1[p])
                        # 将 dp[i][j] 转移而来的长度差 d + num 添加到 dp[p + 1][j] 中
                        dp[p + 1][j].add(d + num)

                    # 情况 2 ： s2[j:q] 都是数字时，状态转移至 dp[i][q]
                    num = 0
                    for q in range(j, m):
                        # 第一个非数字时跳出循环
                        if not s2[q].isdigit():
                            break

                        # 将 s2[j:q + 1] 转换为数字 num
                        num = num * 10 + int(s2[q])
                        # 将 dp[i][j] 转移而来的长度差 d - num 添加到 dp[i][q + 1] 中
                        dp[i][q + 1].add(d - num)
                    
                    # 情况 3 ：s1[i] 是字母且 d < 0
                    if i < n and s1[i].isalpha() and d < 0:
                        # 将 dp[i][j] 转移而来的长度差 d + 1 添加到 dp[i + 1][j] 中
                        dp[i + 1][j].add(d + 1)
                    
                    # 情况 4 ：s2[j] 是字母且 d > 0
                    if j < m and s2[j].isalpha() and d > 0:
                        # 将 dp[i][j] 转移而来的长度差 d - 1 添加到 dp[i][j + 1] 中
                        dp[i][j + 1].add(d - 1)
                    
                # 情况 5 ： 长度差 0 在 dp[i][j] 中 且 s1[i] == s2[j] 且都是字母
                if i < n and j < m and 0 in dp[i][j] and s1[i] == s2[j] and s1[i].isalpha():
                    # 将 dp[i][j] 转移而来的长度差 0 添加到 dp[i + 1][j + 1] 中
                    dp[i + 1][j + 1].add(0)
                    
        # 如果 dp[-1][-1] 中含有长度差 0 ，则 s1 和 s2 匹配成功
        return 0 in dp[n][m]
