# 链接：https://leetcode.com/problems/backspace-string-compare/
# 题意：给定两个字符串 s 和 t ，其中 '#' 代表退格。
#      现在将 s 与 t 输入到文本编辑器中，
#      如果最后的文本一致，则返回 true ，否则返回 false 。


# 数据限制：
#  1 <= s.length, t.length <= 200
#  s 和 t only 仅含有英文小写字母和 '#'


# 输入： s = "ab#c", t = "ad#c"
# 输出： true
# 解释： s 和 t 都会变为 "ac"

# 输入： s = "ab##", t = "c#d#"
# 输出： true
# 解释： s 和 t 都会变为 ""

# 输入： s = "a#c", t = "b"
# 输出： false
# 解释： s 会变成 "c", 而 t 会变成 "b"


# 思路： 栈
#
#      由于这个文本编辑器的过程只有两个操作：输入字符和退格，
#      即后进先出，所以可以用栈来模拟这个过程。
#
#      初始化一个空栈 stack ，
#      然后遍历字符串中的每一个字符 ch ：
#          1. ch == '#': 则将 stack 中的栈顶元素弹出，
#              如果栈为空，则不做处理
#          2. ch != '#': 则将 ch 压入 stack 中
#
#      最后栈 stack 中的字符就是字符串对应的文本。
#
#      获取 s 和 t 对应的文本，然后返回两个文本的比较结果即可。
#
#
#      时间复杂度：O(n + m)
#          1. 需要遍历 s 中全部 O(n) 个字符
#          2. 需要遍历 t 中全部 O(m) 个字符
#      空间复杂度：O(n + m)
#          1. 需要维护 s 对应的文本，最差情况下有 O(n) 个字符
#          2. 需要维护 t 对应的文本，最差情况下有 O(m) 个字符         


class Solution:
    def backspaceCompare(self, s: str, t: str) -> bool:
        # 获取 s 和 t 对应的文本，然后返回比较结果
        return Solution.get_text(s) == Solution.get_text(t)

    @staticmethod
    def get_text(s: str) -> str:
        # 初始化一个空栈
        stack: List[str] = []
        # 遍历字符串中的每一个字符
        for ch in s:
            if ch == '#':
                # 如果是 '#' ，则将栈顶元素弹出
                if stack:
                    stack.pop()
            else:
                # 否则，将字符压入栈中
                stack.append(ch)

        # 最后将所有字符转换为字符串
        return "".join(stack)
