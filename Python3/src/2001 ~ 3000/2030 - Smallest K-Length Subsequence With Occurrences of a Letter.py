# 链接：https://leetcode.com/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter/
# 题意：给定一个字符串 s ，一个整数 k ，一个字母 letter 和一个整数 repetition ，
#       返回 s 中长度为 k 的子序列，这个子序列至少包含 repetition 个字母 letter 。
#       如果有多个符合要求的子序列，返回字典序最小的那个。

# 数据限制：
#   1 <= repetition <= k <= s.length <= 5 * 10 ^ 4
#   s 仅由英文小写字母组成
#   letter 是英文小写字母，且在 s 中至少出现了 repetition 次


# 输入： s = "leet", k = 3, letter = "e", repetition = 1
# 输出： "eet"
# 解释：
#   总共有四个子序列满足题意： "lee", "let", "let", "eet" ，
#   字典序最小的是 "eet"

# 输入： s = "leetcode", k = 4, letter = "e", repetition = 2
# 输出： "ecde"

# 输入： s = "bb", k = 2, letter = "b", repetition = 2
# 输出： "bb"
# 解释：
#   "bb" 是唯一满足题意的子序列


# 思路1： 贪心 + 栈
#
#       维护一个大小为 k 的结果栈 stack ，然后枚举 s 中的第 i 个字符 ch ，
#       每次我们可以先从 stack 中弹出一个字符，只要满足以下条件：
#           1. len(stack) > 0: 栈不为空
#           2. stack[-1] > ch: 栈顶字符比 ch 大
#           3. len(stack) + len(s) - i >= k: 弹出当前字符后，
#               栈中剩余的字符数 + s 中还未放入的字符数 >= k ，
#               保证最终能形成长度为 k 的子序列
#           4. stack[-1] != letter or cnt_letter - 1 + remain_letter >= repetition:
#               (1) 当前字符不是 letter ，则可以放心弹出，
#                   这样仍然会保证子序列中至少含有 repetition 个 letter
#               (2) 当前字符是 letter ，则需要 栈中剩余的 letter 数 + s 中还未放入的 letter 数 >= repetition ，
#                   保证最终形成的子序列中至少含有 repetition 个 letter
#      
#       然后我们决定是否放入当前字符 ch ，如果 len(stack) < k ，
#       则可能可以直接放入：
#           1. ch == letter: 必定可以放入，更新栈中的 letter 数即可
#           2. k - len(stack) > repetition - cnt_letter: 可以放入，保证剩余足够多的位置放下还差的 letter
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def smallestSubsequence(self, s: str, k: int, letter: str, repetition: int) -> str:
        # 初始化剩余的 letter 数
        remain_letter = sum(1 for ch in s if ch == letter)
        # 用数组模拟栈
        stack = []
        # 统计栈中的 letter 数
        cnt_letter = 0

        # 枚举要放入的第 i 个字符 ch
        for i, ch in enumerate(s):
            # 1. 如果栈不为空
            # 2. 栈顶字符比 ch 大
            # 3. 栈中剩余的字符数 + s 中还未放入的字符数 >= k
            # 4. 当前栈顶不是 letter 或者 栈中剩余的 letter 数 + s 中还未放入的 letter 数 >= repetition
            while stack and stack[-1] > ch and len(stack) - 1 + len(s) - i >= k and (stack[-1] != letter or cnt_letter - 1 + remain_letter >= repetition):
                # 弹出栈顶字符，如果该字符是 letter ，则减少 cnt_letter
                if stack.pop() == letter:
                    cnt_letter -= 1

            # 如果 stack 不足 k 个字符，则可能可以继续放入
            if len(stack) < k:
                # 如果 ch 是 letter 则必定能放入
                if ch == letter:
                    stack.append(ch)
                    cnt_letter += 1
                elif k - len(stack) > repetition - cnt_letter:
                    # 如果 ch 不是 letter ，则要留够 repetition - cnt_letter 个位置给 letter
                    stack.append(ch)
            
            # 更新剩余的 letter 数
            if ch == letter:
                remain_letter -= 1

        return ''.join(stack)
