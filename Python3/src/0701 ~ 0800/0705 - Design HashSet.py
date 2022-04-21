# 链接：https://leetcode.com/problems/design-hashset/
# 题意：不用任何内置哈希表相关的库，自己实现一个哈希集合。
#      该哈希集合需要支持以下操作：
#          1. void add(key): 将 key 插入到哈希集合中
#          2. void remove(key): 将 key 从哈希集合中移除
#          3. bool contains(key): 如果哈希集合包含 key 则返回 true ，
#              否则返回 false


# 数据限制：
#  0 <= key <= 10 ^ 6
#  add, remove, contains 最多被调用 10 ^ 4 次


# 输入： ["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
#       [[], [1], [2], [1], [3], [2], [2], [2], [2]]
# 输出： [null, null, null, true, false, null, true, null, false]
# 解释： MyHashSet myHashSet = new MyHashSet();
#       myHashSet.add(1);      # set = [1]
#       myHashSet.add(2);      # set = [1, 2]
#       myHashSet.contains(1); # 返回 True
#       myHashSet.contains(3); # 返回 False ，没有找到
#       myHashSet.add(2);      # set = [1, 2]
#       myHashSet.contains(2); # 返回 True
#       myHashSet.remove(2);   # set = [1]
#       myHashSet.contains(2); # 返回 False ，已被移除


# 思路： 模拟
#
#      由于 key 的范围是 [0, 10 ^ 6] ，
#      所以我们可以直接定义一个长度为 10 ^ 6 + 1 的 bool 数组 set ，
#      set[key] 就表示 key 是否存在于哈希集合中。
#
#
#      时间复杂度：O(1)
#          1. 全部都是数组的直接读写操作，所以时间复杂度是 O(1)
#      空间复杂度：O(v)
#          1. 需要分配大小为 O(v) 的 bool 数组， 其中 v 为哈希集合中的最大值


# 定义 set 的最大长度
MAX_SIZE: int = 1_000_001


class MyHashSet:

    def __init__(self):
        # 初始化长度为 MAX_SIZE 的 bool 数组，
        # 全部设置为 false
        self.set = [False] * MAX_SIZE

    def add(self, key: int) -> None:
        # 将 set[key] 标记为 True
        self.set[key] = True

    def remove(self, key: int) -> None:
        # 将 set[key] 标记为 False
        self.set[key] = False
        

    def contains(self, key: int) -> bool:
        # set[key] 就表示 key 是否存在于 set 中
        return self.set[key]
        


# Your MyHashSet object will be instantiated and called as such:
# obj = MyHashSet()
# obj.add(key)
# obj.remove(key)
# param_3 = obj.contains(key)
