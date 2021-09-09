# 链接：https://leetcode.com/problems/operations-on-tree/
# 题意：给定一颗树， parent[i] 表示 i 节点的父节点， parent[0] = -1 表示 0 节点时根，
#       现在要对这棵树实现以下三种操作：
#           1. lock(num, user): 用户 user 锁定节点 num
#               num 节点未锁定时，可以进行锁定，返回 true
#               num 节点已锁定时，不进行处理，返回 false
#           2. unlock(num, user): 用户 user 解锁节点 num
#               num 节点被 user 锁定时，可以解锁，返回 true
#               num 节点未被 user 锁定时，不进行处理，返回 false
#           3. upgrade(num, user): 用户 user 升级节点 num
#               在满足以下 3 个条件的情况下，锁定 num 节点，并解锁所有子节点，然后返回 true
#                   (1) num 节点
#                   (2) num 节点的子节点至少存在一个被锁定（不限 user ）
#                   (3) num 节点的所有祖先节点未被锁定
#               否则返回 false

# 数据限制：
#   n == parent.length
#   2 <= n <= 2000
#   0 <= parent[i] <= n - 1 for i != 0
#   parent[0] == -1
#   0 <= num <= n - 1
#   1 <= user <= 10 ^ 4
#   parent 能表示一颗合法的树
#   lock, unlock, 和 upgrade 操作最多共 2000 个

# 输入： ["LockingTree", "lock", "unlock", "unlock", "lock", "upgrade", "lock"]
#       [[[-1, 0, 0, 1, 1, 2, 2]], [2, 2], [2, 3], [2, 2], [4, 5], [0, 1], [0, 1]]
# 输出： [null, true, false, true, true, true, false]
# 解释： 
#       LockingTree lockingTree = new LockingTree([-1, 0, 0, 1, 1, 2, 2]);
#       lockingTree.lock(2, 2);    // 返回 true ，因为 2 节点未被锁定
#                                  //   2 节点现在将被用户 2 锁定
#       lockingTree.unlock(2, 3);  // 返回 false ，因为用户 3 不能解锁用户 2 锁定的节点
#       lockingTree.unlock(2, 2);  // 返回 true ，因为用户 2 可以自己锁定的节点
#                                  //   2 节点现在将被用户 2 解锁
#       lockingTree.lock(4, 5);    // 返回 true ，因为 4 节点未被锁定
#                                  //   4 节点现在将被用户 5 锁定
#       lockingTree.upgrade(0, 1); // 返回 true ，因为 0 节点未被锁定，且存在子节点 4 被锁定
#                                  //   0 节点现在将被用户 1 锁定，并且 4 节点将被解锁
#       lockingTree.lock(0, 1);    // 返回 false ，因为 0 节点已被锁定


# 思路： 模拟
#
#       按照题意模拟即可，
#       我们可以发现最多有 2000 个节点，而最多也只有 2000 个操作，
#       所以我们每个操作的时间复杂度在 O(n) 内即可通过
#
#       lock 和 unlock 只涉及当前节点，
#       所以我们可以使用数组 lock_user[i] 表示 i 节点的锁定状态：
#           lock_user[i] == -1: 未锁定
#           lock_user[i] > 0: 当前被用户 lock_user[i] 锁定
#       这样我们就可以在 O(1) 内实现 lock 和 unlock 操作
#
#       upgrade 有三个条件，判断当前节点以及祖先节点是否被锁定，
#       可以直接不断往上找到跟节点，每次判断即可
#
#       判断子节点与解锁子节点可以合并在一起进行处理：
#           我们定义一个函数 _unlock_descendants(num: int) -> int ，
#           表示它能解锁 num 节点的所有子节点，并返回解锁的数量
#       那么在 upgrade 操作时，直接调用，最后判断返回值是否大于 0 即可
#
#       时间复杂度： O(n ^ 2))
#       空间复杂度： O(n ^ 2)


class LockingTree:

    def __init__(self, parent: List[int]):
        # self.parent[i] 表示 i 节点的父节点
        self.parent = parent
        # self.lock_user[i] 表示 i 节点的锁定状态
        self.lock_user = [-1] * len(parent)
        # self.children[i] 存储 i 节点的所有直接子节点列表
        self.children = defaultdict(list)
        for child in range(1, len(parent)):
            self.children[parent[child]].append(child)

    def lock(self, num: int, user: int) -> bool:
        # 如果 num 节点未被锁定，即可进行锁定
        if self.lock_user[num] == -1:
            self.lock_user[num] = user
            return True

        return False

    def unlock(self, num: int, user: int) -> bool:
        # 如果 num 节点被用户 user 锁定，即可进行解锁
        if self.lock_user[num] == user:
            self.lock_user[num] = -1
            return True
        
        return False

    def upgrade(self, num: int, user: int) -> bool:
        # 如果有锁定的祖先节点，则直接返回 False
        if self._has_locked_ancestor(num):
            return False

        # 尝试解锁所有子节点
        cnt = self._unlock_descendants(num)
        # 如果成功解锁过子节点，则可以锁定当前节点
        if cnt:
            self.lock_user[num] = user
            return True

        return False
    
    def _has_locked_ancestor(self, num: int) -> bool:
        while num != -1:
            # 如果有一个祖先节点被锁定，则直接返回 True
            if self.lock_user[num] != -1:
                return True
            num = self.parent[num]
        # 所有祖先在节点均未被锁定，返回 False
        return False
    
    def _unlock_descendants(self, num: int) -> int:
        cnt = 0
        # 如果当前节点被锁定，则进行解锁，并记录解锁数量
        if self.lock_user[num] != -1:
            cnt += 1
            self.lock_user[num] = -1
        
        # 递归处理所有子节点，并记录解锁数量
        for child in self.children[num]:
            cnt += self._unlock_descendants(child)
        return cnt

# Your LockingTree object will be instantiated and called as such:
# obj = LockingTree(parent)
# param_1 = obj.lock(num,user)
# param_2 = obj.unlock(num,user)
# param_3 = obj.upgrade(num,user)
