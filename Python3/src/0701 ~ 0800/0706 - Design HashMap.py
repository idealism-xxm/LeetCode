# 链接：https://leetcode.com/problems/design-hashmp/
# 题意：不用任何内置哈希表相关的库，自己实现一个哈希表。
#      该哈希表需要支持以下操作：
#          1. void put(int key, int value): 将键值对 (key, value) 插入到哈希表中
#          2. int get(int key): 返回 key 对应的值，不存在则返回 -1
#          3. void remove(int key): 将 key 及其对应的值从哈希表中移除


# 数据限制：
#  0 <= key, value <= 10 ^ 6
#  put, get, remove 最多被调用 10 ^ 4 次


# 输入： ["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
#       [[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
# 输出： [null, null, null, 1, -1, null, 1, null, -1]
# 解释： MyHashMap myHashMap = new MyHashMap();
#       myHashMap.put(1, 1); # map = [[1,1]]
#       myHashMap.put(2, 2); # map = [[1,1], [2,2]]
#       myHashMap.get(1);    # 返回 1
#       myHashMap.get(3);    # 返回 -1 （未找到）
#       myHashMap.put(2, 1); # map = [[1,1], [2,1]] （更新已存在的值）
#       myHashMap.get(2);    # 返回 1
#       myHashMap.remove(2); # map = [[1,1]]
#       myHashMap.get(2);    # return -1 （未找到）


# 思路： 模拟
#
#		由于 key 的范围是 [0, 10 ^ 6] ，
#      所以我们可以直接定义一个长度为 10 ^ 6 + 1 的 int 数组 mp ，
#      初始化为 -1 ， mp[key] 就表示 key 哈希表中 key 对应的值。
#
#
#      时间复杂度：O(1)
#          1. 全部都是数组的直接读写操作，所以时间复杂度是 O(1)
#      空间复杂度：O(v)
#          1. 需要分配大小为 O(v) 的 int 数组， 其中 v 为哈希表中的最大值


# 定义 mp 的最大长度
MAX_SIZE: int = 1_000_001


class MyHashMap:

    def __init__(self):
        # 初始化长度为 MAX_SIZE 的 int 数组，
        # 全部设置为 -1
        self.mp = [-1] * MAX_SIZE

    def put(self, key: int, value: int) -> None:
        # 将 mp[key] 标记为 value
        self.mp[key] = value

    def get(self, key: int) -> int:
        # mp[key] 就表示 key 在哈希表中 key 对应的值
        return self.mp[key]

    def remove(self, key: int) -> None:
        # 将 mp[key] 设置为 -1
        self.mp[key] = -1


# Your MyHashMap object will be instantiated and called as such:
# obj = MyHashMap()
# obj.put(key,value)
# param_2 = obj.get(key)
# obj.remove(key)
