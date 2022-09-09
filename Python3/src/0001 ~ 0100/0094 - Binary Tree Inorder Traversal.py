# 链接：https://leetcode.com/problems/binary-tree-inorder-traversal/
# 题意：给定一个二叉树，返回中序遍历的结果？
#
#      进阶：使用循环的方式进行中序遍历。


# 数据限制：
#  二叉树的结点范围为 [0, 100]
#  -100 <= Node.val <= 100


# 输入： root = [1,null,2,3]
# 输出： [1,3,2]
# 解释： 1
#        \
#         2
#        /
#       3

# 输入： root = []
# 输出： []

# 输入： root = [1]
# 输出： [1]


# 思路1：递归/DFS
#
#      递归版本的中序遍历非常简单，核心就是要保证收集顺序如下：
#          1. 先收集左子树的值
#          2. 再收集当前结点的值
#          3. 最后收集右子树的值
#
#
#      时间复杂度：O(n)
#          1. 需要遍历全部 O(n) 个结点
#      空间复杂度：O(n)
#          1. 栈递归深度就是树高 O(h) ，
#              最差情况下，全部 O(n) 个结点在一条链上


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        # ans 用于收集中序遍历的结果，传入 dfs 中递归处理即可
        ans: List[int] = []
        Solution.dfs(root, ans)

        return ans

    def dfs(root: Optional[TreeNode], ans: List[int]):
        # 如果当前子树为空，则直接返回
        if not root:
            return

        # 中序遍历就是需要先收集左子树的值，再收集当前结点的值，最后收集右子树的值
        Solution.dfs(root.left, ans)
        ans.append(root.val)
        Solution.dfs(root.right, ans)


# 思路2：迭代
#
#      将递归转成迭代需要记录栈信息，我们可以用 stack 来记录栈中待处理的结点。
#
#      如果当前结点 cur 非空 或 栈 stack 非空，
#      则还有待处理的结点，继续迭代处理：
#          1. 中序遍历要先遍历左子树的所有结点，所以如果 cur 非空时，
#              要将其入栈，先处理其左子结点，直至 cur 为空
#          2. 此时 stack 顶部就是当前能处理的最左侧结点，将其值放入 ans 中
#          3. 最后需要对 cur.right 同样执行以上操作，
#              那么令 cur = cur.right ，进入下一轮迭代
#
#
#      时间复杂度：O(n)
#          1. 需要遍历全部 O(n) 个结点
#      空间复杂度：O(n)
#          1. 模拟的栈 stack 需要记录树高 O(h) 个待处理的结点，
#              最差情况下，全部 O(n) 个结点在一条链上


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        # ans 用于收集中序遍历的结果
        ans: List[int] = []
        # stack 为迭代模拟递归时需要的栈
        stack: deque = deque()
        # cur 为迭代模拟递归时的当前结点
        cur: Optinal[TreeNode] = root
        # 如果当前结点非空 或 栈不为空，则可以继续循环
        while cur or stack:
            # 中序遍历要先遍历左子树的所有结点，所以如果 cur 非空时，
            # 要将其入栈，先处理其左子结点，直至 cur 为空
            while cur:
                stack.append(cur)
                cur = cur.left

            # 此时 stack 顶部就是当前能处理的最左侧结点，将其值放入 ans 中
            cur = stack.pop()
            ans.append(cur.val)
            # 接下来该处理 cur 的右子树
            cur = cur.right

        return ans
