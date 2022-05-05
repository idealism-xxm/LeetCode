# 链接：https://leetcode.com/problems/implement-stack-using-queues/
# 题意：用队列模拟栈，支持以下 4 种操作：
#          1. void push(int x): 将 x 压入栈中
#          2. int pop(): 弹出栈顶元素
#          3. int top(): 返回栈顶元素
#          4. boolean empty(): 判断栈是否为空


# 数据限制：
#  1 <= x <= 9
#  push, pop, top 和 empty 会调用 100 次
#  所有的 pop 和 top 调用都是合法的


# 输入： ["MyStack", "push", "push", "top", "pop", "empty"]
#       [[], [1], [2], [], [], []]
# 输出： [null, null, null, 2, 2, false]
# 解释： MyStack myStack = new MyStack();
#       myStack.push(1);
#       myStack.push(2);
#       myStack.top();   # 返回 2
#       myStack.pop();   # 返回 2
#       myStack.empty(); # 返回 False


# 思路： 模拟
#
#      由于队列是先进先出，无论怎么都无法换方向，
#      所以我们有两种选择：
#          1. 在入栈时切换方向，那么 push 时候为 O(n) ，
#              pop/top 时候为 O(1)
#          2. 在出栈时切换方向，那么 push 时候为 O(1) ，
#              pop/top 时候为 O(n)
#
#      第一种方法可以只使用一个队列，也更容易处理。
#
#      我们将队列维护成栈的顺序，即后进的元素在前面，先进的元素在后面。
#
#      在 push 的时候，先记录此时的队列长度 length ，
#      然后将元素放入队尾，最后将队首的 length 个元素依次弹出并放入队尾。
#
#      这样在 pop/top 的时候只用返回队首元素即可。
#
#
#      时间复杂度：push 是 O(n), pop/top/empty 是 O(1)
#          1. push: 需要操作队列中全部 O(n) 个数字
#          2. pop/top: 只需要操作队首的数字一次
#          3. empty: 只需要判断队列是否为空
#      空间复杂度：总空间复杂度为 O(n) ，每个操作都是 O(1)
#          1. push/pop/top/empty: 只需要使用常数个额外变量


class MyStack:

    def __init__(self):
        # 模拟所需的队列
        self.q = deque()

    def push(self, x: int) -> None:
        # 获取需要移动的元素个数
        length: int = len(self.q)
        # 将元素放入队尾
        self.q.append(x)
        # 将队列前 length 个元素依次弹出并放入队尾
        while length > 0:
            # 弹出队首元素
            num: int = self.q.popleft()
            # 将元素放入队尾
            self.q.append(num)
            # 需要移动的元素个数减 1
            length -= 1

    def pop(self) -> int:
        return self.q.popleft()

    def top(self) -> int:
        return self.q[0]

    def empty(self) -> bool:
        # 判断队列是否为空
        return not len(self.q)


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()
