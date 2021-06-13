# 链接：https://leetcode.com/problems/minimum-cost-to-change-the-final-value-of-expression/
# 题意：给定一个合法的逻辑表达式，只含有 0, 1, &, |, (, ) ，
#       有以下四种操作：
#           1. 将一个 0 变为 1
#           2. 将一个 1 变为 0
#           3. 将一个 & 变为 |
#           4. 将一个 | 变为 &
#       求最少执行多少次操作后，可以改变原表达式的结果？

# 数据限制：
#   1 <= expression.length <= 105
#   表达式只含有 0, 1, &, |, (, )
#   所有括号都正确匹配
#   没有空括号，即 "()" 不是表达式的子串

# 输入： expression = "1&(0|1)"
# 输出： 1
# 解释： "1&(0|1)" ->
#       "1&(0&1)"


# 输入： expression = "(0&0)&(0&0&0)"
# 输出： 3
# 解释： "(0&0)&(0&0&0)" ->
#       "(0|0)&(0&0&0)" ->
#       "(0|1)&(0&0&0)" ->
#       "(0|1)|(0&0&0)"

# 输入： expression = "(0|(1|0&1))"
# 输出： 1
# 解释： "(0|(1|0&1))" ->
#       "(0|(0|0&1))"

# 思路1： DFS + 栈
#
#       先通过栈计算出每一对括号的匹配位置，
#       定义函数 dfs(expression, mp, begin, end) ，
#       其返回值为 (val, num) ,
#           val 表示子表达式 expression[begin:end + 1] 的值
#           num 表示改变子表达式 expression[begin:end + 1] 的值所需最小的操作数
#
#       我们先分类讨论 p OP q 的不同情况（假设 p_val, p_num 和 q_val, q_num 已经计算好了）：
#           1. p_val = 0, OP = &, q_val = 0  -->  (0, 1 + min(p_num, q_num))
#               此时必定要将 OP 变为 | ，因为左右操作数都是 0 ，两者都改变不会优于只改变一个
#           2. p_val = 0, OP = &, q_val = 1  -->  (0, 1)
#              p_val = 1, OP = &, q_val = 0  -->  (0, 1)
#               此时必定要将 OP 变为 | ，因为左/右操作数变为 1 至少需要操作一次
#           3. p_val = 1, OP = &, q_val = 1  -->  (1, min(p_num, q_num))
#               此时必定是改变左右操作数中的一个，挑操作次数少的即可
#           4. p_val = 0, OP = |, q_val = 0  -->  (0, min(p_num, q_num))
#               此时必定是改变左右操作数中的一个，挑操作次数少的即可
#           5. p_val = 0, OP = |, q_val = 1  -->  (0, 1)
#              p_val = 1, OP = |, q_val = 0  -->  (0, 1)
#               此时必定要将 OP 变为 & ，因为左/右操作数变为 0 至少需要操作一次
#           6. p_val = 1, OP = |, q_val = 1  -->  (1, 1 + min(p_num, q_num))
#               此时必定要将 OP 变为 & ，因为左右操作数都是 1 ，两者都改变不会优于只改变一个
#
#       这六种情况再合并相同条件的可得：
#           1. p_val != q_val  -->  (0, 1)
#           2. p_val == q_val 时必定包含 min(p_num, q_num)
#           需要 +1 的情况：
#               (1) p_val == 0 and OP == &
#               (2) p_val == 1 and OP == |
#
#       以上情况是针对一个操作符而言，但所有的操作符号都是从左到右计算的，
#       所以我们可以把多个操作符号转换为以上的情况，再同时讨论其他情况即可：
#           1. 唯一一个数： 直接返回 (int(ch), 1)
#           2. 最外层是括号：直接返回 dfs(expression, mp, begin + 1, end - 1)
#           3. a OP_0 b ... OP q: 其中 OP 是最外层子表达式的最后一个操作符号，
#               即将前面的表达式 a OP_0 b ... 看成一个整体 p 进行处理，
#               这样就转换为前面讨论的情况了
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def minOperationsToFlip(self, expression: str) -> int:
        stack: List[int] = []
        end_to_begin: Dict[int, int] = {}
        for i, ch in enumerate(expression):
            if ch == '(':
                stack.append(i)
            elif ch == ')':
                begin = stack.pop()
                end_to_begin[i] = begin

        return self.dfs(expression, end_to_begin, 0, len(expression) - 1)[1]

    def dfs(self, expression: str, mp: Dict[int, int], begin: int, end: int) -> Tuple[int, int]:
        """返回表达式 expression[begin:end + 1] 的值，及改变该值所需的最小操作数"""
        # 相等时，则当前一定是 0 或 1
        if begin == end:
            # 前者表示表达式的值，后者表示改变所需的操作数
            return int(expression[begin]), 1
        # 找到最后一个子表达式的开始下标，
        # 如果最后一个字符不是后括号，那么最后一个字符就是最后一个子表达式
        last_sub_exp_begin = mp.get(end, end)
        # 如果最后一个子表达式的开始下标，就是当前表达式的开始下标，
        # 则当前表达式最外层是括号，直接返回括号内的值即可
        if begin == last_sub_exp_begin:
            return self.dfs(expression, mp, begin + 1, end - 1)
        # 现在将其转换成 p OP q 这种形式，前面所有的部分看出一个子表达式求解
        p_val, p_num = self.dfs(expression, mp, begin, last_sub_exp_begin - 2)
        q_val, q_num = self.dfs(expression, mp, last_sub_exp_begin, end)

        op = expression[last_sub_exp_begin - 1]
        val = p_val and q_val if op == '&' else p_val or q_val
        # 六种情况合并后的第一种，改变符号即可
        if p_val != q_val:
            # 根据操作符计算当前表达式的值
            return val, 1
        # 剩下的情况必定包含 min(p_num, q_num)
        num = min(p_num, q_num)
        # 如果是必定改变操作符的情况，则还要 +1
        if (p_val == 0 and op == '&') or (p_val == 1 and op == '|'):
            num += 1
        return val, num


# 思路2： DP + 调度场算法
#
#       有了前面 DFS + 栈解法的分析，我们可以将其转换为 DP + 栈，
#       即在出入栈的过程中使用调度场算法，计算表达式的值，同时更新 DP 值
#
#       由于需要和调度场算法结合，其符号栈和操作数栈必备，
#       我们就将刚刚操作数栈改造一下，用 dp[-1] = (val, num) 替代原有的操作数栈
#           val 表示子表达式 expression[begin:end + 1] 的值
#           num 表示改变子表达式 expression[begin:end + 1] 的值所需最小的操作数
#
#       我们先分类讨论 p OP q 的不同情况（假设 p_val, p_num 和 q_val, q_num 已经计算好了）：
#           1. p_val = 0, OP = &, q_val = 0  -->  (0, 1 + min(p_num, q_num))
#               此时必定要将 OP 变为 | ，因为左右操作数都是 0 ，两者都改变不会优于只改变一个
#           2. p_val = 0, OP = &, q_val = 1  -->  (0, 1)
#              p_val = 1, OP = &, q_val = 0  -->  (0, 1)
#               此时必定要将 OP 变为 | ，因为左/右操作数变为 1 至少需要操作一次
#           3. p_val = 1, OP = &, q_val = 1  -->  (1, min(p_num, q_num))
#               此时必定是改变左右操作数中的一个，挑操作次数少的即可
#           4. p_val = 0, OP = |, q_val = 0  -->  (0, min(p_num, q_num))
#               此时必定是改变左右操作数中的一个，挑操作次数少的即可
#           5. p_val = 0, OP = |, q_val = 1  -->  (0, 1)
#              p_val = 1, OP = |, q_val = 0  -->  (0, 1)
#               此时必定要将 OP 变为 & ，因为左/右操作数变为 0 至少需要操作一次
#           6. p_val = 1, OP = |, q_val = 1  -->  (1, 1 + min(p_num, q_num))
#               此时必定要将 OP 变为 & ，因为左右操作数都是 1 ，两者都改变不会优于只改变一个
#
#       这六种情况再合并相同条件的可得：
#           1. p_val != q_val  -->  (0, 1)
#           2. p_val == q_val 时必定包含 min(p_num, q_num)
#           需要 +1 的情况：
#               (1) p_val == 0 and OP == &
#               (2) p_val == 1 and OP == |
#
#       此时我们是按照调度场算法从做往右计算的，
#       所以只要符号栈栈顶不是前括号，都需要计算
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def minOperationsToFlip(self, expression: str) -> int:
        stack: List[int] = []
        dp: List[Tuple[int, int]] = []
        for i, ch in enumerate(expression):
            # 如果是 (, &, | ，则放入符号栈，处理下一个字符
            if ch in '(&|':
                stack.append(ch)
                continue
            # 如果是后括号，则将符号栈顶的符号移除（必定是前括号）
            if ch == ')':
                stack.pop()
            else:
                # 此时 ch 是 0 或 1 ，放入操作数栈 dp
                dp.append((int(ch), 1))

            # 此时如果符号栈有符号，且栈顶不是前括号，则需要进行计算
            if len(stack) > 0 and stack[-1] != '(':
                op = stack.pop()
                q_val, q_num = dp.pop()
                p_val, p_num = dp.pop()
                val = p_val and q_val if op == '&' else p_val or q_val
                # 六种情况合并后的第一种，改变符号即可
                if p_val != q_val:
                    # 根据操作符计算当前表达式的值
                    num = 1
                else:
                    # 剩下的情况必定包含 min(p_num, q_num)
                    num = min(p_num, q_num)
                    # 如果是必定改变操作符的情况，则还要 +1
                    if (p_val == 0 and op == '&') or (p_val == 1 and op == '|'):
                        num += 1
                # 将新的状态入栈
                dp.append((val, num))

        # 此时操作数栈只有一个元素，就是最终算出来的结果
        return dp[0][1]


# 思路3： DP + 调度场算法
#
#       还有另一种状态定义方式，和前面的方式类似，但不需要讨论太多的情况，
#       直接根据当前的值进行状态转移即可
#       （比赛时也想到了这种状态定义，但是没有想到结合调度场算法，怎么处理都时 O(n ^ 2)）
#
#       由于需要和调度场算法结合，其符号栈和操作数栈必备，
#       我们就将刚刚操作数栈改造一下，用 dp[-1] = (zero_num, one_num) 替代原有的操作数栈
#           zero_num 表示让子表达式的值为 0 时所需的最小操作数
#           one_num  表示让子表达式的值为 1 时所需的最小操作数
#
#       我们先分类讨论 p OP q 的不同情况：
#       （假设 p_zero_num, p_one_num 和 q_zero_num, q_one_num 已经计算好了）
#           1. OP == &:
#               # 要得到 0 ，只需要左右任意一个表达式为 0 即可
#               zero_num = min(p_zero_num, q_zero_num)
#               # 要得到 1
#               # (1) 左右表达式都变为 1
#               # (2) 操作符变为 | ，左右任意一个表达式为 0 即可
#               one_num  = min(p_one_num + q_one_num, 1 + min(p_one_num, q_one_num))
#           2. OP == |:
#               # 要得到 1
#               # (1) 左右表达式都变为 0
#               # (2) 操作符变为 & ，左右任意一个表达式为 0 即可
#               zero_num = min(p_zero_num + q_zero_num, 1 + min(p_zero_num, q_zero_num))
#               # 要得到 1 ，只需要左右任意一个表达式为 1 即可
#               one_num  = min(p_one_num, q_one_num)
#
#       此时我们是按照调度场算法从做往右计算的，
#       所以只要符号栈栈顶不是前括号，都需要计算
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def minOperationsToFlip(self, expression: str) -> int:
        stack: List[int] = []
        dp: List[Tuple[int, int]] = []
        for i, ch in enumerate(expression):
            # 如果是 (, &, | ，则放入符号栈，处理下一个字符
            if ch in '(&|':
                stack.append(ch)
                continue
            # 如果是后括号，则将符号栈顶的符号移除（必定是前括号）
            if ch == ')':
                stack.pop()
            elif ch == '0':
                dp.append((0, 1))
            else:
                dp.append((1, 0))

            # 此时如果符号栈有符号，且栈顶不是前括号，则需要进行计算
            if len(stack) > 0 and stack[-1] != '(':
                op = stack.pop()
                q_zero_num, q_one_num = dp.pop()
                p_zero_num, p_one_num = dp.pop()
                if op == '&':
                    # 要得到 0 ，只需要左右任意一个表达式为 0 即可
                    zero_num = min(p_zero_num, q_zero_num)
                    # 要得到 1
                    # (1) 左右表达式都变为 1
                    # (2) 操作符变为 | ，左右任意一个表达式为 0 即可
                    one_num = min(p_one_num + q_one_num, 1 + min(p_one_num, q_one_num))
                else:
                    # 要得到 1
                    # (1) 左右表达式都变为 0
                    # (2) 操作符变为 & ，左右任意一个表达式为 0 即可
                    zero_num = min(p_zero_num + q_zero_num, 1 + min(p_zero_num, q_zero_num))
                    # 要得到 1 ，只需要左右任意一个表达式为 1 即可
                    one_num = min(p_one_num, q_one_num)
                # 将新的状态入栈
                dp.append((zero_num, one_num))

        # 此时操作数栈只有一个元素，就是最终算出来的结果，其中一个必定是 0
        return max(*dp[0])
