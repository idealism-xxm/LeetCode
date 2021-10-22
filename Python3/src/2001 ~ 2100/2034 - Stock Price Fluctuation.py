# 链接：https://leetcode.com/problems/stock-price-fluctuation/
# 题意：有一个股票的价格纪录列表 records ，
#       每一个记录 record = (timestamp, price) 表示该股票在 timestamp 价格为 price ，
#       这个记录可能不会按顺序到来，也有可能同一个 timestamp 对应不同的 price ，
#       我们认为后到达的记录是正确的，将会更新先到达的价格。
#       现在实现一个算法支持如下操作：
#           1. update(timestamp, price): 更新 timestamp 时的价格为 price
#           2. current(): 获取当前股票的最新价格（即 timestamp 最大的价格）
#           3. maximum(): 获取当前所有记录的最大价格
#           4. minimum(): 获取当前所有记录的最小价格

# 数据限制：
#   1 <= timestamp, price <= 10 ^ 9
#   总共最多有 10 ^ 5 次调用这些函数： update, current, maximum, minimum.
#   current, maximum, minimum 会在至少调用一次 update 后调用

# 输入： ["StockPrice", "update", "update", "current", "maximum", "update", "maximum", "update", "minimum"]
#       [[], [1, 10], [2, 5], [], [], [1, 3], [], [4, 2], []]
# 输出： [null, null, null, 5, 10, null, 5, null, 2]
#       
# 解释：
#   - StockPrice stockPrice = new StockPrice();
#   - stockPrice.update(1, 10); // Timestamps 是 [1] ，对应的价格为 [10]
#   - stockPrice.update(2, 5);  // Timestamps 是 [1,2] ，对应的价格为 [10,5].
#   - stockPrice.current();     // 返回 5, 最新的 timestamp 是 2 ，对的价格是 5
#   - stockPrice.maximum();     // 返回 10, 最大的价格是 10 对应的 timestamp 是 1
#   - stockPrice.update(1, 3);  // 前一个为 1 的 timestamp 是错误的价格，所以其价格会被更新为 3
#                               // Timestamps 是 [1,2] ，对应的价格为 [3,5].
#   - stockPrice.maximum();     // 返回 5, 最大的价格是 5 对应的 timestamp 是 2
#   - stockPrice.update(4, 2);  // Timestamps 是 [1,2,4] ，对应的价格为 [3,5,2].
#   - stockPrice.minimum();     // 返回 2, 最小的价格是 2 对应的 timestamp 是 4


# 思路： map + 堆
#
#       mp 维护每个时间戳对应的最新价格的 map
#       mx 维护价格的最大值的堆
#       mn 维护价格的最小值的堆
#       cnt 维护每个价格出现次数的 map
#       last 维护最新时间戳及其价格
#
#       update(timestamp, price):
#           1. 如果当前时间戳更小，则先更新 last = (timestamp, price)
#           2. 先查询该时间戳的旧价格 pre_price ，如果存在，则 cnt[pre_price] -= 1
#           3. 再更新当前时间戳的最新价格， mp[timestamp] = price
#           4. 最新价格出现次数 +1 ， cnt[price] += 1
#           5. 分别放入价格最大值和最小值的堆
#
#       current(): 直接返回 last[1]
#
#       maximum(): 从 mx 中查出最大价格，如果该价格的出现次数为 0 ，
#                   则从堆中移除，继续获取下一个最大值，直至价格的出现次数不为 0 ，
#                   返回这个最大价格即可
#
#       minimum(): 从 mn 中查出最小价格，如果该价格的出现次数为 0 ，
#                   则从堆中移除，继续获取下一个最小值，直至价格的出现次数不为 0 ，
#                   返回这个最小价格即可
#           
#
#       时间复杂度： 每个操作平均时间复杂度为 O(logn)
#       空间复杂度： O(n)


class StockPrice:

    def __init__(self):
        # 维护价格的最大值的堆， self.mx[0] 为最大值
        self.mx = []
        # 维护价格的最小值的堆， self.mn[0] 为最小值
        self.mn = []
        # 维护每个价格的次数
        self.cnt = defaultdict(int)
        # 维护每个时间戳的价格
        self.mp = {}
        # 维护最新的时间戳和价格
        self.last = (0, 0)

    def update(self, timestamp: int, price: int) -> None:
        # 如果当前时间戳更大，则更新 self.last
        # 这里取等号，是为了更新该时间戳的最新价格
        if self.last[0] <= timestamp:
            self.last = (timestamp, price)

        # 获取该时间戳的前一次价格
        pre_price = self.mp.get(timestamp)
        # 如果前一次价格存在，则改价格次数 -1
        if pre_price is not None:
            self.cnt[pre_price] -= 1
        # 当前价格次数 +1
        self.cnt[price] += 1
        # 更新该时间戳的价格
        self.mp[timestamp] = price

        # 更新价格的最大的值堆和最小值的堆
        heapq.heappush(self.mx, -price)
        heapq.heappush(self.mn, price)

    def current(self) -> int:
        # 当前价格在 self.last 中维护
        return self.last[1]

    def maximum(self) -> int:
        # 如果最大的价格已经不存在了，则从最大值的堆中移除
        while self.cnt[-self.mx[0]] == 0:
            heapq.heappop(self.mx)

        # 当前最大价格就是还存在的最大价格
        return -self.mx[0]
        

    def minimum(self) -> int:
        # 如果最小的价格已经不存在了，则从最小值的堆中移除
        while self.cnt[self.mn[0]] == 0:
            heapq.heappop(self.mn)

        # 当前最小价格就是还存在的最小价格
        return self.mn[0]


# Your StockPrice object will be instantiated and called as such:
# obj = StockPrice()
# obj.update(timestamp,price)
# param_2 = obj.current()
# param_3 = obj.maximum()
# param_4 = obj.minimum()
