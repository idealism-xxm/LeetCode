# 链接：https://leetcode.com/problems/removing-stars-from-a-string/
# 题意：给定一个含有 * 的字符串 s ，每次操作可以进行如下处理：
#          1. 选择 s 中的一个 *
#          2. 删除当前 * 左侧第一个非 * 字符，然后再删除当前 *
#
#      返回所有 * 都被删除后的字符串。
#
#      注意：
#          1. 数据保证所有的操作都是合法的
#          2. 结果是唯一的


# 数据限制：
#  1 <= s.length <= 10 ^ 5
#  s 仅由英文小写字母和 * 组成
#  s 可以执行上述操作


# 输入： s = "leet**cod*e"
# 输出： "lecoe"
# 解释： 从左到右移除 * ：
#          1. lee(t*)*cod*e -> lee*cod*e
#          2. le(e*)cod*e   -> lecod*e
#          3. leco(d*)e     -> lecoe


# 输入： s = "erase*****"
# 输出： ""
# 解释： 从左到右移除 * ：
#          1. eras(e*)**** -> eras****
#          2. era(s*)***   -> era***
#          3. er(a*)**     -> er**
#          4. e(r*)*       -> e*
#          4. (e*)         -> ""


# 思路： 模拟
#
#      我们用一个数组 ans 维护结果字符串中的字符，初始化为空。
#
#      然后枚举 s 中的每一个字符 ch ：
#          1. ch == '*': 需要移除其左侧第一个字符，
#              而这个字符必定是 ans 中的最后一个字符，
#              所以直接删除 ans 中的最后一个字符即可
#          2. ch != '*': 直接放入 ans 中即可
#
#      最后 ans 就是满足题意的结果。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 s 中全部 O(n) 个字符一次
#      空间复杂度：O(n)
#          1. 需要维护结果字符串中的字符，最差情况下有 O(n) 个字符


class Solution:
    def removeStars(self, s: str) -> str:
        # ans 维护结果字符串中的字符
        ans: List[int] = []
        for ch in s:
            if ch == '*':
                # 如果当前字符是 * ，需要移除其左侧第一个字符，
                # 也就是 ans 中的最后一个字符
                ans.pop()
            else:
                # 如果当前字符非 * ，直接放入 ans 中即可
                ans.append(ch)

        return "".join(ans)
