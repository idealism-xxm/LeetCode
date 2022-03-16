# 链接：https://leetcode.com/problems/validate-stack-sequences/
# 题意：给定 pushed 和 popped 两个序列，每个序列中的值都不重复，
#      如果在空栈上用 pushed 中的值执行 push 和 pop 操作，
#      最终 pop 出的序列能得到 popped 时，返回 true；
#      否则，返回 false 。


# 数据限制：
#   1 <= pushed.length <= 1000
#   0 <= pushed[i] <= 1000
#   pushed 的所有元素 互不相同
#   popped.length == pushed.length
#   popped 是 pushed 的一个排列


# 输入： pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
# 输出： true
# 解释： 我们可以按以下顺序执行：
#       push(1), push(2), push(3), push(4), pop() -> 4,
#       push(5), pop() -> 5, 
#       pop() -> 3, 
#       pop() -> 2, 
#       pop() -> 1

# 输入： pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
# 输出： false
# 解释： 1 不能在 2 之前弹出


# 思路： 栈
#
#       本题有一个很关键的条件：序列中的值都不重复，
#       所以我们可以模拟栈一个栈，
#       每次判断栈顶数字和 popped 需要出栈的数字是否相等，
#       来决定如何处理。
#   
#       我们可以维护一个栈 stack ，
#       并维护当前 pushed 中已入栈的数字个数 pushed_cnt 。
#   
#       那么我们按顺序遍历 popped 中需要出栈的数字 cur ：
#           1. 如果当前栈顶数字不是 cur ，
#               则当前只能不断入栈 pushed 中的下一个数字，
#               直至栈顶的数字是 cur 或者 pushed 中的数字全部入栈。
#           2. 结束前面的循环后，此时栈顶数字 top 出栈，有两种情况，
#               (1) top == cur ，则说明出栈成功，进入下一个循环
#               (2) top != cur ，则说明出栈失败，
#                   并且 pushed 中的所有数字都已入栈，
#                   popped 必定是一个不合法的栈序列，
#                   直接返回 False
#
#       遍历完 popped 中的所有数字后，
#       表明每一个需要出栈的数字都能找到对应的入栈数字，
#       则 popped 是一个合法的栈序列，返回 True
#
#      
#       时间复杂度：O(n)
#           1. 需要遍历 pushed 中全部 O(n) 个数字
#           2. 需要遍历 popped 中全部 O(n) 个数字
#       空间复杂度：O(n)
#           1. 需要维护一个栈 stack ，最差情况下需要存储全部 O(n) 个数字


class Solution:
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        n: int = len(pushed)
        # 定义一个栈，用来模拟当前栈的状态
        stack: List[int] = []
        # 维护 pushed 已经入栈的数字个数
        pushed_cnt = 0
        # 遍历需要出栈的数字
        for cur in popped:
            # 如果当前栈顶数字不是当前需要出栈的数字，则当前只能入栈新数字。
            # 如果此时 pushed 中还有数字可以入栈时，则需要循环处理。
            while (not stack or stack[-1] != cur) and pushed_cnt < n:
                # 将 pushed[pushed_cnt] 入栈
                stack.append(pushed[pushed_cnt])
                # 令 pushed_cnt 加 1
                pushed_cnt += 1

            # 如果当前栈顶数字不是当前需要出栈的数字，
            # 那么 popped 不是一个合法的栈序列，
            # 直接返回 False
            if stack.pop() != cur:
                return False

        # 此时，每一个需要出栈的数字都能找到对应的入栈数字，
        # 则 popped 是一个合法的栈序列，返回 True
        return True
