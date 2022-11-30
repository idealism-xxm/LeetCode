# 链接：https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/
# 题意：实现一个数据结构 RandomizedCollection ，支持在一个可重复集合中插入、删除数，
#      以及随机返回集合中的一个数，需要实现以下操作:
#          1. RandomizedCollection(): 初始化一个 RandomizedCollection 对象
#          2. bool insert(int val): 往集合中插入 val 。
#              如果 val 原本不在集合中，则返回 true ；否则返回 false 。
#          3. bool remove(int val): 从集合中删除 val 。
#              如果 val 原本在集合中，则返回 true ；否则返回 false 。
#          4. int getRandom(): 返回集合中随机一个数。（调用时保证集合中至少存在一个数）
#              需要保证每个数字被返回的可能性与其数量线性相关。
#
#      进阶：集合的三个方法的平均时间复杂度都为 O(1) 。


# 数据限制：
#  -(2 ^ 31) <= val <= 2 ^ 31 - 1
#  最多调用 insert, remove, getRandom 共 2 * 10 ^ 5 次。
#  调用 getRandom 时保证集合中至少存在一个数


# 输入： ["RandomizedCollection", "insert", "insert", "insert", "getRandom", "remove", "getRandom"]
#       [[], [1], [1], [2], [], [1], []]
# 输出： [null, true, false, true, 2, true, 1]
# 解释： RandomizedCollection randomizedCollection = new RandomizedCollection();
#       randomizedCollection.insert(1);   # 往集合中插入 1 。 1 原本不在集合中，返回 true 。集合变为 [1]
#       randomizedCollection.insert(1);   # 往集合中插入 1 。 1 原本在集合中，返回 false 。集合变为 [1,1]
#       randomizedCollection.insert(2);   # 往集合中插入 2 。 2 原本不在集合中，返回 true 。集合变为 [1,1,2]
#       randomizedCollection.getRandom(); # 应该 2/3 概率返回 1 ， 1/3 概率返回 2
#       randomizedCollection.remove(1);   # 从集合中移除 1 。 1 原本在集合中，返回 true 。集合变为 [1,2]
#       randomizedCollection.getRandom(); # 应该 1/2 概率返回 1 ， 1/2 概率返回 2


# 思路： Map + Set
#
#      本题是 LeetCode 380 的加强版，允许集合中的数字可重复，
#      所以 map 的值需要使用 set 来维护。
#
#
#      如果数据结构题需要用到 O(1) 的查询操作，一般都要使用 map 来辅助。
#
#      我们可以用 nums 维护集合中的数。
#
#      注意到使用类似 random 的库函数，
#      能在 O(1) 内随机获得范围 [0, nums.len()) 中的一个数 i ，
#      那么也就能在 O(1) 内随机返回集合中的数 nums[i] 。
#
#      同时我们也需要在 O(1) 内从 nums 中删除一个数 val ，
#      那么就需要知道 val 在 nums 中的下标，这就需要使用 map 来维护这个关系。
#
#      由于数字可重复出现， map 的值需要使用 set 来维护，以保证能在 O(1) 内对值进行操作。
#
#      删除 val 时，如果 val 在 map 中，则获取对应下标集合中的任意一个下标 index ，
#      将最后一个数移动至 index （注意更新下标）。
#      并从对应的下标集合中删除 index ，再删除 nums 中最后一个数。
#
#      插入 val 时，如果 val 不在 map 中，
#      则将 val 放入 nums 末尾，并将其下标插入 map 中。
# 
#
#      时间复杂度： O(1)
#          1. 集合的三个方法都只需要执行常数次操作
#      空间复杂度： O(n)
#          1. 需要维护全部 O(n) 个数字​


class RandomizedCollection:

    def __init__(self):
        # nums 维护集合中的全部数字，用于 O(1) 随机返回一个数字
        self.nums: List[int] = []
        # num_to_index_set[num] 表示 num 在 nums 中的下标集合，用于 O(1) 插入/删除一个数字
        self.num_to_index_set: Dict[int, Set[int]] = defaultdict(set)

    def insert(self, val: int) -> bool:
        # 获取 val 所在的下标集合
        index_set: Set[int] = self.num_to_index_set[val]
        
        # 否则将 val 放入 nums 中，并记录其在 nums 中的下标
        index_set.add(len(self.nums))
        self.nums.append(val)

        # 如果只有一个下标，则说明 val 原本不在集合中，需要返回 true
        return len(index_set) == 1

    def remove(self, val: int) -> bool:
        # 如果 val 原本不在集合中，则直接返回 false
        if val not in self.num_to_index_set or len(self.num_to_index_set[val]) == 0:
            return False

        # 获取 val 的任意一个下标 index
        index_set: Set[int] = self.num_to_index_set[val]
        index: int = next(iter(index_set))
        # 先将 index 从下标集合中移除，防止最后一个数也是 val 时多删下标
        index_set.remove(index)
        last_index: int = len(self.nums) - 1
        # 更新最后一个数的下标
        last_index_set: Set[int] = self.num_to_index_set[self.nums[last_index]]
        # 要先加 index ，再删除 last_index ，防止 index == last_index 时没删下标
        last_index_set.add(index)
        last_index_set.remove(last_index)
        # 将最后一个数移动至 index 处
        self.nums[index] = self.nums[last_index]

        # O(1) 移除最后一个数字（即 val ）
        self.nums.pop()
        return True

    def getRandom(self) -> int:
        # 随机返回 nums 中的一个数
        return random.choice(self.nums)


# Your RandomizedCollection object will be instantiated and called as such:
# obj = RandomizedCollection()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()
