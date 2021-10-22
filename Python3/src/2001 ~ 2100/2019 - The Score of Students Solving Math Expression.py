# 链接：https://leetcode.com/problems/the-score-of-students-solving-math-expression/
# 题意：给定一个合法的表达式字符串 s ，只含有数字和 '+', '*' 。
#       给出一个整型数组 answers ，表示学生提交的答案列表，现在要求所有答案的分数和。
#       对于给定的 s 有正确答案为 correct_answer ，
#           (1) 若一个提交的答案 answer 与 correct_answer 相等，则得 5 分
#           (2) 若一个提交的答案 answer 与 s 按照某一错误的顺序得出得结果相等，则得 2 分
#           (3) 不满足前两个条件，则得 0 分

# 数据限制：
#   3 <= s.length <= 31
#   s 是一个合法的表达式字符串，只含有 0-9, '+', 和 '*'
#   s 中所有操作符两侧的数字范围在 [0,9] 内
#   操作符 ('+', '*') 的数量范围在 [0,15] 内
#   表达式 s 的正确结果范围在 [0, 1000] 内
#   n == answers.length
#   1 <= n <= 10 ^ 4
#   0 <= answers[i] <= 1000


# 输入： s = "7+3*1*2", answers = [20,13,42]
# 输出： 19
# 解释：
#   正确结果： 7+3*1*2 = 7+3*2 = 7 + 6 = 13
#   错误结果： (7+3)*1*2 = 10*1*2 = 20
#   
#   因此得分分别为： [2,5,0]

# 输入： s = "3+5*2", answers = [13,0,10,13,13,16,16]
# 输出： 19
# 解释：
#   正确结果： 3+5*2 = 3+10 = 13
#   错误结果： (3+5)*2 = 8*2 = 16
#
#   因此得分分别为： [5,0,0,5,5,2,2]

# 输入： s = "6+0*1", answers = [12,9,6,4,8,6]
# 输出： 10
# 解释：
#   正确结果： 6+0*1 = 6+0 = 6
#
#   因此得分分别为： [0,0,5,0,0,5]

# 输入： s = "1+2*3+4", answers = [13,21,11,15]
# 输出： 10
# 解释：
#   正确结果： 1+2*3+4 = 1+6+4 = 7+4 = 11
#   错误结果： (1+2)*3+4 = 3*3+4 = 9+4 = 13
#   错误结果： 1+2*(3+4) = 1+2*7 = 1+14 = 15
#   错误结果： (1+2)*(3+4) = 3*(3+4) = 3*7 = 21
#
#   因此得分分别为： [2,2,5,2]


# 思路： 调度场算法 + DP
#
#       我们既要计算正确的结果，还要计算所有可能的错误结果，
#       1. 正确结果：使用调度场算法计算即可
#       2. 错误结果：使用区间 DP ，计算每个区间内所有可能的结果即可
#           dp[i][j] 表示 s[i:j + 1] 表达式所有可能的结果
#           初始化： dp[i][i] = int(s[i]) , i 是偶数
#           状态转移：
#               (1) 先枚举区间长度 l ，从 3 开始，每次增加 2 ，直到长度为 n
#               (2) 其中枚举区间起始位置 i ，从 1 开始，每次增加 2 ，直到为 n - l ，
#                   并计算出结束位置 j = i + l - 1 ，我们将计算 s[i:j + 1] 内所有可能的结果
#               (3) 其中枚举最后一个操作符的位置 k ，从 i + 1 开始，每次增加 2 ，直到 j - 1
#                   此时根据操作符，对其左右两边集合内的数字计算，
#                   然后将小于等于 1000 的结果放入到 dp[i][j] 中即可
#
#       最后 dp[0][n - 1] 即为 s 中所有可能的结果，按照题意计算得分即可
#
#       时间复杂度： O((m ^ 3) * (k ^ 2)) ，其中 m = n // 2, k = 1000
#       空间复杂度： O((n ^ 2) * k)


OPERATOR_PRIORITY = {
    '$': -1,
    '+': 0,
    '*': 1,
}

OPERATOR_CALC_FUNC = {
    '+': lambda x, y: x + y,
    '*': lambda x, y: x * y,
}


class Solution:
    def scoreOfStudents(self, s: str, answers: List[int]) -> int:
        # 计算正确答案
        correct_answer = self.calculate(s)
        # 初始化 dp 数组
        n = len(s)
        dp = [[set() for _ in range(n)] for _ in range(n)]
        # 偶数位置，只有一个数字，直接计算
        for i in range(0, n, 2):
            dp[i][i].add(int(s[i]))

        # 先枚举要计算的区间长度（长度更小的在前面已经计算完成）
        for l in range(3, n + 1, 2):
            # 再枚举区间起始位置
            for i in range(0, n - l + 1, 2):
                # 计算结束位置
                j = i + l - 1
                # 再枚举最后一个操作符的位置
                for k in range(i + 1, j, 2):
                    # 枚举左侧区间可能的结果
                    for left in dp[i][k - 1]:
                        # 枚举右侧区间可能的结果
                        for right in dp[k + 1][j]:
                            # 计算当前区间的结果
                            result = OPERATOR_CALC_FUNC[s[k]](left, right)
                            if result <= 1000:
                                # 将结果放入到 dp 数组中
                                dp[i][j].add(result)

        # 计算得分
        ans = 0
        for answer in answers:
            if answer == correct_answer:
                ans += 5
            elif answer in dp[0][n - 1]:
                ans += 2
        return ans

    def calculate(self, s: str) -> int:
        s += '$'
        num_stack = []
        oprator_stack = []
        for ch in s:
            if ch in OPERATOR_PRIORITY:
                # 如果是操作符号，且其优先级 小于等于 符号栈顶操作符优先级别，则需要先计算
                while oprator_stack and OPERATOR_PRIORITY[ch] <= OPERATOR_PRIORITY[oprator_stack[-1]]:
                    # 弹出数字栈顶的两个数字，注意顺序，后弹出的在操作符左侧
                    b, a = num_stack.pop(), num_stack.pop()
                    # 弹出操作符
                    operator = oprator_stack.pop()
                    # 根据操作符计算，然后压入数字栈
                    num_stack.append(OPERATOR_CALC_FUNC[operator](a, b))

                # 当前操作符入栈
                oprator_stack.append(ch)
            else:
                # 如果是数字，则直接入栈即可
                num_stack.append(int(ch))
        
        # 最后数字栈只有一个数字，就是正确结果
        return num_stack[0]
