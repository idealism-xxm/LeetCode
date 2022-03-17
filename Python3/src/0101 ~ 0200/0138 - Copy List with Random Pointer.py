# 链接：https://leetcode.com/problems/copy-list-with-random-pointer/
# 题意：给一个单链表，每一个结点除了有 next 指向下一个结点，
#		还有一个 random 指向链表中随机的一个结点，
#		现在深拷贝这个链表，并返回结果链表的头结点。

# 数据限制：
#	0 <= n <= 1000
#	-(10 ^ 4) <= Node.val <= 10 ^ 4
#	Node.random 为空或者链表中存在的某个结点


# 输入： head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
# 输出： [[7,null],[13,0],[11,4],[10,2],[1,0]]

# 输入： head = [[1,1],[2,1]]
# 输出： [[1,1],[2,1]]

# 输入： head = [[3,null],[3,0],[3,null]]
# 输出： [[3,null],[3,0],[3,null]]


# 思路1： 递归 + Map
#
#      	相同的结点只会被拷贝一次，因此我们可以维护一个 Map 来记录，
#      	origin_to_cloned[origin] 表示结点 origin 对应的已经拷贝过的结点。
#
#		从给定的结点开始递归拷贝即可，
#			1. 如果当前结点 cur 已拷贝过，则直接进行引用并返回
#			2. 如果当前结点 cur 未拷贝过，则创建一个拷贝的结点 cloned ，
#          	将其放入到 origin_to_cloned 中
#
#          	然后递归处理 cur.next 结点。
#
#			再直接从 origin_to_cloned 中，
#			取出 cur.random 对应的拷贝结点，
#			放入到 cloned.random 中，
#			因为此时链表中的所有结点已经被深拷贝过。
#
#
#		关联题目： LeetCode 133 - 克隆图
#
#
#		时间复杂度： O(n)
#			1. 需要遍历链表中的全部 O(n) 个结点
#		空间复杂度： O(n)
#			1. 需要存储拷贝链表中的全部 O(n) 个结点
#			2. 需要维护一个 O(n) 的 Map 来记录已经拷贝过的结点


"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

from typing import Optional


class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        # 维护每个结点下标对应的拷贝结点
        origin_to_cloned: Dict[Node, Node] = {}

        def dfs(cur: Optional[Node]) -> Optional[Node]:
            # 如果当前结点 cur 为空，则直接返回即可
            if not cur:
                return None

            # 如果当前结点 cur 已拷贝过，则直接取出返回
            if cur in origin_to_cloned:
                return origin_to_cloned[cur]

            # 此时需要创建当前结点 cur 的拷贝结点 cloned
            cloned: Node = Node(cur.val)
            # 将 cloned 放入到 origin_to_cloned 中
            origin_to_cloned[cur] = cloned
            # 递归处理 cur.next 结点
            cloned.next = dfs(cur.next)
            # 此时链表中所有的结点都已拷贝过，
            # 直接从 origin_to_cloned 中取出即可
            if cur.random:
                # 取出 cur.random 对应的拷贝结点
                cloned.random = origin_to_cloned[cur.random]

            # 返回当前结点的拷贝结点 cloned
            return cloned

        return dfs(head)


# 思路2： 三次迭代
#
#      	在思路 1 中，我们使用了 Map 记录了原始结点对应的克隆结点，
#       如果不使用 Map ，那似乎就没有很好的方法处理了。
#
#       看了讨论区后，发现可以充分利用链表的特性，
#       使用三次迭代处理，这样就可以避免使用 Map 了。
#
#       1. 第一次迭代，先对原链表的每个结点都进行克隆，
#           并将克隆结点 cloned 插入到原结点 origin 后面。
#
#           现在链表的变为 A -> a -> B -> b ...  的形式，
#           （其中大写字母代表原结点 origin ，小写字母代表对应的克隆结点 cloned ）
#           这样只要找到了原结点 origin ，就能找到对应的克隆结点 cloned = origin.next
#
#       2. 第二次迭代，将克隆链表中的每个结点的 random ，
#           指向对应的克隆结点 random.next 。
#
#       3. 第三次迭代，先保存克隆链表的头结点 cloned_head = head.next 。
#           然后将原链表和克隆链表分开，即复原原链表，并建立克隆链表。
#
#
#		时间复杂度： O(n)
#			1. 需要遍历链表中的全部 O(n) 个结点
#		空间复杂度： O(n)
#			1. 需要存储拷贝链表中的全部 O(n) 个结点


"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        # 如果原链表为空，则直接返回即可
        if not head:
            return None

        # 第一次迭代，为原链表中的每个结点创建一个克隆结点，并插入到原结点之后。
        origin: Optional[Node] = head
        # 如果原链表还有结点，那么继续克隆
        while origin:
            # 克隆当前结点
            cloned: Node = Node(origin.val, origin.next, origin.random)
            # 将克隆结点 cloned 插入到当前结点 origin 之后
            origin.next = cloned
            # 移动至下一个原始结点
            origin = cloned.next
        
        # 现在链表的变为 A -> a -> B -> b ...  的形式，
        # （其中大写字母代表原结点 origin ，小写字母代表对应的克隆结点 cloned ）
        # 这样只要找到了原结点 origin ，就能找到对应的克隆结点 cloned = origin.next

        # 第二次迭代，将克隆链表中的每个结点的 random 指向对应的克隆结点 random.next
        cloned: Optional[Node] = head.next
        # 如果克隆链表还有结点，那么继续处理其 random
        while cloned:
            # 将克隆结点的 random 指向对应的克隆结点 random.next 。
            # 注意 random 不为空时，才有对应的克隆结点
            if cloned.random:
                cloned.random = cloned.random.next
            # 移动至下一个原结点
            cloned = cloned.next
            # 如果还有原结点，则移动至对应的克隆结点
            if cloned:
                cloned = cloned.next

        # 记录克隆链表的头结点，拆分完成后返回
        cloned_head: Optional[Node] = head.next
        # 第三次迭代，将原链表和克隆链表分开，即复原原链表，并建立克隆链表
        origin: Optional[Node] = head
        cloned: Optional[Node] = cloned_head
        # 如果原链表还有结点，那么继续进行拆分
        while origin:
            # 将原结点的 next 指向下一个原结点
            origin.next = origin.next.next
            # 如果还有原结点，则将克隆结点的 next 指向下一个克隆结点
            if cloned.next:
                cloned.next = cloned.next.next

            # 原结点移动至下一个原结点
            origin = origin.next
            # 克隆结点移动至下一个克隆结点
            cloned = cloned.next

        # 返回克隆链表的头结点
        return cloned_head
