# 链接：https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
# 题意：给定一个二叉树，将其转换成链表形式？

# 输入： root = [1,2,5,3,4,null,6]
# 输出： [1,null,2,null,3,null,4,null,5,null,6]
# 解释：
#     1             1
#    / \             \
#   2   5             2
#  / \   \      →      \
# 3   4   6             3
#                        \
#                         4
#                          \
#                           5
#                            \
#                             6

# 输入： root = []
# 输出： []

# 输入： root = [0]
# 输出： [0]


# 思路1：递归
#
#		对于子树 root 来说，可以递归调用处理成三部分
#			1. 当前根结点
#			2. 左子树形成的链表（可能为空链表）
#			3. 右子树形成的链表（可能为空链表）
#
#		然后将三部分按顺序连接起来即可。
#
#      最后返回链表的头结点和尾结点，方便上层处理。
#
#
#		时间复杂度： O(n)
#          1. 需要遍历全部 O(n) 个结点一次
#		空间复杂度： O(n)
#          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        Solution.dfs(root)

    @staticmethod
    def dfs(root: Optional[TreeNode]) -> Tuple[Optional[TreeNode], Optional[TreeNode]]:
        if not root:
            return None, None

        # 递归处理左右结点，并获取对应链表的头结点和尾结点
        left_head, left_tail = Solution.dfs(root.left)
        right_head, right_tail = Solution.dfs(root.right)
        # 清空左子结点
        root.left = None

        # 当前结点目前既是头结点，也是尾结点
        head, tail = root, root
        # 将左半部分挂在链表尾部
        if left_head:
            tail.right = left_head
            tail = left_tail
        # 将右半部分挂在链表尾部
        if right_head:
            tail.right = right_head
            tail = right_tail

        # 返回当前子树转换成的链表头结点和尾结点
        return head, tail
