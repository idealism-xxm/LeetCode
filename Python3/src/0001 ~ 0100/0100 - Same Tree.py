# 链接：https://leetcode.com/problems/same-tree/
# 题意：判断两个二叉树是否一样？


# 数据限制：
#  树的结点数范围为 [0, 100]
#  -(10 ^ 4) <= Node.val <= 10 ^ 4


# 输入： p = [1,2,3], q = [1,2,3]
# 输出： true
# 解释： 1             1
#      / \           / \
#     2   3         2   3

# 输入： p = [1,2], q = [1,null,2]
# 输出： false
# 解释： 1             1
#      /               \
#     2                 2

# 输入： p = [1,2,1], q = [1,1,2]
# 输出： false
# 解释： 1             1
#      / \           / \
#     2   1         1   2

# 思路：递归
#
#      递归处理即可
#          1. 若当前结点的值相同，则左子树和右子树都相同才返回 true
#          2. 若当前结点都值不同，则直接返回 false
#      
#      时间复杂度： O(n)
#          1. 需要遍历全部 O(n) 个结点
#      空间复杂度： O(h)
#          1. 栈递归深度就是树高 h ，最差情况下，全部 O(n) 个结点在一条链上


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        # 都是空结点，则当前子树相同，直接返回 true
        if not p and not q:
            return True

        # 如果当前结点都值不同，则当前子树不同，直接返回 false
        if not p or not q or p.val != q.val:
            return False

        # 都是非空结点，且当前结点值相同，则需要左右子树是否相同
        return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
