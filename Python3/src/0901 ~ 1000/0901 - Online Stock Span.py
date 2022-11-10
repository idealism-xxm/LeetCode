# 链接：https://leetcode.com/problems/online-stock-span/
# 题意：设计一个算法，支持收集一个股票每天的价格 price ，并返回当天价格的跨度。
#
#      当日股票的跨度是股票价格小于或等于今天价格的最大连续日数
#      （从今天开始往回数，包括今天）。例如：
#
#          如果股票连续 7 天的价格为 [100,80,60,70,60,75,85] ，
#          那么股票跨度将是 [1,1,1,2,1,4,6]
#
#      实现 StockSpanner 类的两个方法：
#          1. StockSpanner(): 初始化该类的一个实例
#          2. int next(int price): 给定今天的价格 price ，并返回今天价格的跨度


# 数据限制：
#  1 <= price <= 10 ^ 5
#  最多会调用 10 ^ 4 次 next


# 输入： ["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
#       [[], [100], [80], [60], [70], [60], [75], [85]]
# 输出： [null, 1, 1, 1, 2, 1, 4, 6]
# 解释： StockSpanner stockSpanner = new StockSpanner();
#       stockSpanner.next(100); # 返回 1
#       stockSpanner.next(80);  # 返回 1
#       stockSpanner.next(60);  # 返回 1
#       stockSpanner.next(70);  # 返回 2
#       stockSpanner.next(60);  # 返回 1
#       stockSpanner.next(75);  # 返回 4 ，因为过去 4 天（包含今天的价格 75 ） 的价格小于等于今天的价格
#       stockSpanner.next(85);  # 返回 6


# 思路： 单调栈
#
#      朴素的想法肯定维护全部天的价格，并从今天 i 往前遍历，
#      直到遇到第一个比今天价格大的 j ，那么 i 的跨度就是 i - j 。
#
#      这是一个 O(n ^ 2) 的思路，在题目给定的数据范围内无法通过，
#      所以需要一个更优的解法。
#
#
#      这样处理没有利用已知的信息，存在大量重复的遍历。
#      其实只要把每次的跨度也维护起来，就能避免重复遍历的情况。
#
#      我们维护数组 arr ， arr[i] = [price_i, span_i] ，
#      其中 price_i 表示第 i 天的价格， span_i 表示第 i 天的跨度。
#
#      当调用 next(price) 时，第 i 天的价格为 price ，初始化跨度 span 为 1 ，
#      令 j 为前一天 i - 1 ，然后按照如下方式进行处理：
#          1. arr[j].price <= price: 
#              则从第 j 天开始往前 arr[j].span 天的价格都小于等于 price ，
#              直接将 arr[j].span 计入 span 中，即 span += arr[j].span 。
#              我们也无需再重复遍历这几天，即 j -= arr[j].span ，然后重复这个流程。
#          2. arr[j].price > price: 则第 i 天的跨度到此为止，不再重复
#
#      执行完上述流程后， span 就是第 i 天的跨度，放入数组 arr 中，然后返回 span 即可。
#
#
#      此时可以进一步优化，可以注意到如果第 j 天的跨度是 arr[j].span ，
#      那么前面两个分支都不会遍历从第 j 天开始往前的 arr[j].span 天。
#
#      那么就不需要再维护这些信息，去除这些天的信息后， arr 就变成了单调递减栈 stack 。
#      流程也有所变化：
#          1. stack.top().price <= price:
#              则从这天天开始往前 stack.top().span 天的价格都小于等于 price ，
#              直接将 stack.top().span 计入 span 中，即 span += stack.top().span 。
#              那么 stack.top() 的跨度将不会再被遍历到，出栈即可，然后重复这个流程。
#          2. stack.top().price > price: 则第 i 天的跨度到此为止，不再重复
#
#
#		时间复杂度： 平均 O(1)
#          1. 总共调用 O(n) 次 next ，全部 O(n) 天的信息会入栈 1 次，
#              这些信息至多出栈 1 次。
#              平摊下来每次调用 next 会有 1 次入栈，不超过 1 次出栈，
#              平均时间复杂度为 O(1) 。
#		空间复杂度： O(n)
#          1. 需要维护价格递减的那些天的信息，
#              最差情况下，全部 O(n) 天价格都是递减的


class StockSpanner:

    def __init__(self):
        # 单调递减栈， stack[i] = (price, span) ，
        #  price: 某天（记作 x ）的价格，严格单调递减
        #  span: 跨度，即 x 过去 span 天（包含 x ）的价格均小于等于 price
        self.stack: List[Tuple] = []

    def next(self, price: int) -> int:
        # span 维护今天价格的跨度，初始化为 1 ，表示今天价格必定满足题意
        span: int = 1
        while self.stack:
            pre_price, pre_span = self.stack[-1][0], self.stack[-1][1]
            if pre_price <= price:
                # 如果栈顶价格小于等于 price ，
                # 则其过去 pre_span 天的价格均小于等于 price ，
                # 可以直接计入 span 中，并将栈顶元素出栈
                span += pre_span
                self.stack.pop()
            else:
                # 如果栈定价格大于 price ，则跨度到此为止
                break

        # 此时，栈为空 或 栈顶价格大于 price ，将今天的信息入栈
        self.stack.append((price, span))
        
        # 返回今天价格的跨度
        return span


# Your StockSpanner object will be instantiated and called as such:
# obj = StockSpanner()
# param_1 = obj.next(price)
