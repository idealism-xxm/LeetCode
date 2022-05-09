# 链接：https://leetcode.com/problems/letter-combinations-of-a-phone-number/
# 题意：给定一个数字串，每一位数字范围在 [2, 9] 内，
#      在九宫格键盘下一次按下相应的数字键，
#      求所有可能打出来的英文字符串？


# 数据限制：
#  0 <= digits.length <= 4
#  digits[i] 是一个 ['2', '9'] 范围内的数位


# 输入： digits = "23"
# 输出： ["ad","ae","af","bd","be","bf","cd","ce","cf"]

# 输入： digits = ""
# 输出： []

# 输入： digits = "2"
# 输出： ["a","b","c"]


# 思路： 递归
#
#      我们使用 dfs(digits, index, cur, ans) 遍历收集所有可能的字符串，其中：
#          1. digits: 输入的数字串
#          2. index: 当前遍历到的下标
#          3. cur: 当前已遍历的数字串的一个可能的字符串
#          4. ans: 当前收集到的所有可能的字符串的列表
#
#      在 dfs 中，我们按照如下逻辑处理即可：
#          1. index == len(digits) ，则表明已经遍历完数字串，
#              此时 cur 就是一个可能的字符串，将其加入到 ans 中。
#          2. index != len(digits) ，则表明还需要继续遍历数字串，
#              遍历 digits[index] 对应的字母列表串中的字符 ch ，
#              将 cur[index] 设置为 ch ，然后递归调用 dfs 。
#
#
#      时间复杂度：O(4 ^ n)
#          1. 需要遍历全部可能的字符串，最差情况下所有的数字键都是 4 个字母的，
#              共有 O(4 ^ n) 个可能的字符串
#      空间复杂度：O(4 ^ n)
#          1. 需要收集全部可能的字符串，最差情况下所有的数字键都是 4 个字母的，
#              共有 O(4 ^ n) 个可能的字符串


# 定义每个数位对应的字母列表
DIGIT_TO_LETTERS: List[str] = ["", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"];


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        # 如果没有按任何数字键，则返回空列表
        if not digits:
            return []

        # ans 用于收集所有可能的字符串
        ans: List[str] = []
        # cur 表示当前按键下能形成的某个字符列表
        cur: List[str] = [''] * len(digits)
        # 递归收集所有可能的字符串
        Solution.dfs(digits, 0, cur, ans)

        return ans

    @staticmethod
    def dfs(digits: str, index: int, cur: List[str], ans: List[str]):
        # 如果已按下全部数字键，则 cur 就是一个可能的字符串，收集后返回
        if index == len(digits):
            ans.append("".join(cur))
            return

        # 遍历 digits[index] 下对应的的所有字母
        for ch in DIGIT_TO_LETTERS[ord(digits[index]) - ord('0')]:
            # 将当前字母放入 cur 中
            cur[index] = ch
            # 递归收集下一个字母
            Solution.dfs(digits, index + 1, cur, ans)
