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


# 思路2： 贪心
#
#       我们维护每个字符所有出现的位置，以及每个位置及其后面的 letter 数
#       然后枚举答案中的第 i 个位置，贪心从 'a' 开始枚举 ch ，直至 'z' ，
#       假设 ch 最左侧的位置为 j ，
#       那么枚举的字符能放入第 i 个位置的条件要满足以下条件：
#           1. len(s) - j >= k - i: 保证最终子序列长度能达到 k
#           2. remain_letter[j] >= repetition - cnt_letter: 
#               保证最终子序列中至少含有 repetition 个 letter
#           3. ch == letter or k - i - 1 >= repetition - cnt_letter: 
#               保证还有足够的位置放下 repetition - cnt_letter 个 letter
#
#       时间复杂度： O(n + k * t) ，其中 t = 26
#       空间复杂度： O(n)


class Solution:
    def smallestSubsequence(self, s: str, k: int, letter: str, repetition: int) -> str:
        # remain_letter[i] 表示 s[i:] 中的 letter 数
        remain_letter = [0] * len(s)
        # ch_to_postions[ch] 表示 ch 在 s 中的位置列表
        ch_to_postions = defaultdict(list)
        ch_to_postions[s[-1]].append(len(s) - 1)
        if s[-1] == letter:
            remain_letter[-1] = 1
        # 初始化 remain_letter 和 ch_to_postions
        for i in range(len(s) - 2, -1, -1):
            ch_to_postions[s[i]].append(i)
            if s[i] == letter:
                remain_letter[i] = remain_letter[i + 1] + 1
            else:
                remain_letter[i] = remain_letter[i + 1]

        ans = [''] * k
        # ans 中最后一个字符的位置
        last_position = -1
        # ans 中 letter 数
        cnt_letter = 0

        # 枚举答案的第 i 个位置
        for i in range(k):
            # 枚举答案中的第 i 个位置的字符（保证字典序）
            for ch in string.ascii_lowercase:
                # 如果 positions 的最后一个位置小于等于 last_position ，
                # 则最后一个位置需要丢弃，因为子序列的下标单调递增
                positions = ch_to_postions[ch]
                while positions and positions[-1] <= last_position:
                    positions.pop()
                if not positions:
                    continue
                
                # 获取 ch 可用的最小位置
                j = positions[-1]
                # 1. len(s) - j >= k - i: 保证最终子序列长度能达到 k
                # 2. remain_letter[j] >= repetition - cnt_letter: 
                #       保证最终子序列中至少含有 repetition 个 letter
                # 3. ch == letter or k - i - 1 >= repetition - cnt_letter: 
                #       保证还有足够的位置放下 repetition - cnt_letter 个 letter
                if len(s) - j >= k - i and remain_letter[j] >= repetition - cnt_letter and (ch == letter or k - i - 1 >= repetition - cnt_letter):
                    # 当前位置放入 ch
                    ans[i] = ch
                    # 并记录 ans 中最后一个字符的位置是 j
                    last_position = j
                    # 如果当前字符是 letter ，则 cnt_letter 加 1
                    if ch == letter:
                        cnt_letter += 1
                    
                    # 跳出处理下一个位置
                    break

        return ''.join(ans)
