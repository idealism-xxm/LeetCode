# 链接：https://leetcode.com/problems/longest-substring-without-repeating-characters/
# 题意：给定一个字符串，求出不含重复字符的最长子串的长度？

# 输入： s = "abcabcbb"
# 输出： 3
# 解释： "abc" 的长度最长

# 输入： s = "bbbbb"
# 输出： 1
# 解释： "b" 的长度最长

# 输入： s = "pwwkew"
# 输出： 3
# 解释： "wke" 的长度最长

# 输入： s = ""
# 输出： 0
# 解释： "" 的长度最长

# 思路： 双指针
#
#       维护两个指针 l, r ，我们要保证 s[l: r + 1] 一直满足题意，且长度最长
#       最开始 l = r = 0 ，表示第一个字符一定满足题意
#       然后每次右移 r ，如果新增的字符 s[r] 已经存在了，则不断左移 l 直至 s[r] 不存在
#       此时 l 就是以 r 为右边界时，左边界能取的最小值，那么其对应的长度就是 r - l + 1
#       我们取所有这些长度的最大值即可
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        # 如果为空，则直接返回 0
        if not s:
            return 0

        # 第一个字符肯定复合要求
        ans: int = 1
        # 已经出现的字符
        existent_chars: Set[str] = {s[0]}
        # 接下来开始检查 s[l: r + 1] 是否符合要求
        l, r = 0, 1
        while r < len(s):
            # 当 s[r] 已经出现过时，不断移动 l ，直至 s[r] 未出现过
            while s[r] in existent_chars:
                existent_chars.remove(s[l])
                l += 1

            # 更新最大长度，加上新增的字符
            existent_chars.add(s[r])
            r += 1
            ans = max(ans, r - l)

        return ans
