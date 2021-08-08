# 链接：https://leetcode.com/problems/delete-characters-to-make-fancy-string
# 题意：给定一个字符串 s ，删除最少的字符后能形成的奇幻字符串。
#       奇幻字符串：不存在任意连续 3 个字符都相同的字符串。

# 数据限制：
#   2 <= s.length <= 10 ^ 5
#   s 由英文小写字母组成

# 输入： s = "leeetcode"
# 输出： "leetcode"
# 解释： 移除 "le(e)etcode" 括号中的一个字母 e 之后，剩余字符串 "leetcode" 符合要求。

# 输入： s = "aaabaaaa"
# 输出： "aabaa"
# 解释： 移除 "a(a)ab(a)(a)aa" 括号中的三个字母 a 之后，剩余字符串 "aabaa" 符合要求。

# 输入： s = "aab"
# 输出： s = "aab"


# 思路： 模拟
#
#       我们统计前一个字符 pre ，及其连续出现的次数，
#       如果当前字符 ch 与 pre 相同且 pre 连续出现的次数大于 2 次 ，
#       则当前字符不放入结果串中，
#       否则放入结果串中
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)

class Solution:
    def makeFancyString(self, s: str) -> str:
        ans = ''
        pre, cnt = '', 0
        for ch in s:
            if ch == pre:
                cnt += 1
                # 如果字符相同，且目前 ch 出现字符小于 3 次，
                # 则当前字符可以放入结果串中
                if cnt < 3:
                    ans += ch
            else:
                # 不相同，则放入结果串中
                ans += ch
                cnt = 1
                pre = ch

        return ans
