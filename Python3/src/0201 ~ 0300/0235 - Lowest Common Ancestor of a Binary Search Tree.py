# 链接：https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
# 题意：给定一个二叉搜索树，找到两个结点的最近公共祖先？


# 数据限制：
#  树的结点数在 [2, 10 ^ 5] 之间
#  -(10 ^ 9) <= Node.val <= 10 ^ 9
#  所有的 Node.val 都各不相同
#  p != q
#  p 和 q 必定在树中


# 输入： root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
# 输出： 6
# 解释： 结点 2 和 8 的 LCA 是 6
#       6
#    /      \
#   2        8
#  / \      / \
# 0   4    7   9
#    / \
#   3   5

# 输入： root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
# 输出： 2
# 解释： 结点 2 和 4 的 LCA 是 2
#       6
#    /      \
#   2        8
#  / \      / \
# 0   4    7   9
#    / \
#   3   5


# 思路： 迭代
#
#      本题是 LeetCode 236 的简化版，充分利用二叉搜索树的特性即可。
#
#      因为给定的树是二叉搜索树，所以要寻找的两个结点第一次分开的结点就是它们的 LCA 。
#
#      首先为了方便处理，令 p.val < q.val ，然后迭代判断即可：
#          1. p.val <= cur.val && p.val >= cur.val ，则两者有不同的走向，
#             那么 cur 就是 p 和 q 的 LCA ，直接返回。
#             【注意】这里为了方便处理，同时考虑了 p 或 q 是 LCA 的情况，所以条件取等号了
#          2. p.val < q.val < cur.val ，则两者的 LCA 在左子树
#          3. cur.val < p.val < q.val ，则两者的 LCA 在右子树
#
#
#      时间复杂度： O(h)
#          1. 只需要遍历树高 h 个结点，最差情况下，全部 O(n) 个结点在一条链上
#      空间复杂度： O(1)
#          1. 只需要使用常数个额外变量即可


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        # 为了方便处理，令 p.val < q.val
        if p.val > q.val:
            p, q = q, p

        cur: Optional[TreeNode] = root
        while True:
            # 如果 p.val <= cur.val && p.val >= cur.val ，则两者有不同的走向，
            # 那么 cur 就是 p 和 q 的 LCA ，直接返回
            # 【注意】这里为了方便处理，同时考虑了 p 或 q 是 LCA 的情况，所以条件取等号了
            if p.val <= cur.val and cur.val <= q.val:
                return cur

            if p.val <= cur.val:
                # 此时有 p.val < q.val < cur.val ，则两者的 LCA 在左子树
                cur = cur.left
            else:
                # 此时有 cur.val < p.val < q.val ，则两者的 LCA 在左子树
                cur = cur.right

        # p 和 q 在树中，所以必定存在 LCA ，不会走到这里
