# 链接：https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
# 题意：给定一个单链表的头结点，交换正数第 k 个结点和倒数第 k 个结点的值，
#      返回结果单链表的头结点。


# 数据限制：
#  链表中的结点数量为 n
#  1 <= k <= n <= 10 ^ 5
#  0 <= Node.val <= 100


# 输入： head = [1,2,3,4,5], k = 2
# 输出： [1,4,3,2,5]
# 解释： 1 ->  2  -> 3 ->  4  -> 5
#                   ↓
#       1 -> (4) -> 3 -> (2) -> 5


# 输入： head = [7,9,6,6,7,8,3,0,9,5], k = 5
# 输出： [7,9,6,6,8,7,3,0,9,5]
# 解释： 7 -> 9 -> 6 -> 6 ->  7  ->  8  -> 3 -> 0 -> 9 -> 5
#                               ↓
#       7 -> 9 -> 6 -> 6 -> (8) -> (7) -> 3 -> 0 -> 9 -> 5


# 思路： 双指针
#
#      我们找到正数第 k 个结点 front 。
#
#      再定义当前结点 cur = front ，并令 back = head ，
#      然后不断同时向后移动 cur 和 back ，直至 cur.next 为空，
#      那么此时 back 就指向倒数第 k 个结点。
#
#      最后交换 front 和 back 的值即可。
#
#      
#      时间复杂度：O(n)
#          1. 需要遍历链表全部 O(n) 个节点
#      空间复杂度：O(n)
#          1. 只需要维护常数个额外变量


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapNodes(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        # front 初始指向正数第 1 个结点
        front: Optional[ListNode] = head
        # front 向后移动 k - 1 次，指向正数第 k 个结点
        for _ in range(k - 1):
            # front 向后移动一个结点
            front = front.next

        # 定义当前指针 cur ，初始化为 front
        cur: Optional[ListNode] = front
        # back 初始指向正数第 1 个结点
        back: Optional[ListNode] = head
        # 当 cur.next 为空时， back 就指向倒数第 k 个结点
        while cur.next:
            # cur 向后移动一个结点
            cur = cur.next
            # back 向后移动一个结点
            back = back.next

        # 交换 front 和 back 的值
        front.val, back.val = back.val, front.val

        # 返回结果单链表的头结点
        return head
