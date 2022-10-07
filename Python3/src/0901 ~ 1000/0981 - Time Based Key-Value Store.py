# 链接：https://leetcode.com/problems/time-based-key-value-store/
# 题意：设计一个基于时间的键值数据结构，能存储同一个 key 的多个不同时间点的值，
#      并能获取某个时间点 key 对应的 value 。
#
#      实现以下方法：
#          1. TimeMap(): 初始化该数据结构
#          2. void set(String key, String value, int timestamp): 
#              在 timestamp 时存储 key 的值为 value
#          3. String get(String key, int timestamp):
#              获取 timestamp 及之前 key 最后一次存储的 value ，不存在则返回 ""


# 数据限制：
#  1 <= key.length, value.length <= 100
#  key 和 value 仅由英文小写字母和数字组成
#  1 <= timestamp <= 10 ^ 7
#  调用 set 方法的所有 timestamp 都是严格单调递增的
#  set 和 get 最多会调用 2 * 10 ^ 5 次


# 输入： ["TimeMap", "set", "get", "get", "set", "get", "get"]
#       [[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
# 输出： [null, null, "bar", "bar", null, "bar2", "bar2"]
# 解释： TimeMap timeMap = new TimeMap();
#       timeMap.set("foo", "bar", 1);  # 在时间点 1 时存储 "foo" 的值为 "bar"
#       timeMap.get("foo", 1);         # 返回 "bar"
#       timeMap.get("foo", 3);         # 返回 "bar" ，因为 "foo" 在时间点 2 和 3 没有存储任何值
#       timeMap.set("foo", "bar2", 4); # 在时间点 4 时存储 "foo" 的值为 "bar2"
#       timeMap.get("foo", 4);         # 返回 "bar2"
#       timeMap.get("foo", 5);         # 返回 "bar2"


# 思路： 二分
#
#      题目数据保证调用 set 方法的所有 timestamp 都是严格单调递增的，
#      所以对于同一个 key 下的所有值来说，对应的时间点都是升序的。
#
#      那么我们可以将同一个 key 的值按照 set 顺序放入同一个数组 values 中，
#      在查询时，先找到 values 中时间点第一个大于 timestamp 的值的下标 k 。
#
#      如果 k 为 0 ，则表示不存在这样的值，则返回 "" 。
#      如果 k 不为 0 ，则 values[k - 1] 的值是 timestamp 及之前最后一次存储的，返回该值即可。
#
#      由于存在多个 key ，我们还要将 values 数组维护在一个 map 中。
#
#      
#      时间复杂度：set - O(1) | get - O(logn)
#          1. set: 只需要常数次操作即可
#          2. get: 需要对 values 数组进行二分。
#              最差情况下，所有操作的 key 都一样，前一半是 set 操作，那么 values 数组的大小为 O(n) 。
#              后一半是 get 操作，那么每次都需要二分 O(n) 个元素，二分时间复杂度为 O(logn) 
#      空间复杂度：O(n)
#          1. 需要存储全部 O(n) 次 set 操作设置的值和 timestamp


class TimeMap:

    def __init__(self):
        self.key_to_values = defaultdict(list)

    def set(self, key: str, value: str, timestamp: int) -> None:
        # 将 timestamp 和 value 放入列表中
        self.key_to_values[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        # 获取 key 对应的 value 列表
        values: List[Tuple[int, str]] = self.key_to_values[key]
        # 找到 values 中时间点第一个大于 timestamp 的值的下标
        k: int = bisect.bisect_right(values, timestamp, key=lambda value: value[0])
        # k == 0 表示不存在这样的值，则返回 ""
        if k == 0:
            return ""
        # 此时 values[k - 1] 的值是 timestamp 及之前最后一次存储的
        return values[k - 1][1]


# Your TimeMap object will be instantiated and called as such:
# obj = TimeMap()
# obj.set(key,value,timestamp)
# param_2 = obj.get(key,timestamp)
