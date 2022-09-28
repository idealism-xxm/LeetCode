# 链接：https://leetcode.com/problems/remove-nth-node-from-end-of-list/
# 题意：给定一个链表，删除第 n 个元素。
#
#      进阶：只遍历一次。


# 数据限制：
#  链表有 sz 个结点
#  1 <= sz <= 30
#  0 <= Node.val <= 100
#  1 <= n <= sz


# 输入： head = [1,2,3,4,5], n = 2
# 输出： [1,2,3,5]
# 解释： 1 -> 2 -> 3 -> 4 -> 5
#                 ↓
#       1 -> 2 -> 3    ->   5

# 输入： head = [1], n = 1
# 输出： []


# 输入： head = [1,2], n = 1
# 输出： [1]


# 思路2：快慢指针
#
#      对于链表的题目，一般都可以使用一个哨兵结点。
#
#      本题使用哨兵结点方便处理移除头结点这种边界情况。
#
#      我们将快慢指针 fast 和 slow 都初始化为哨兵结点，先将快指针 fast 移动 n 次。
#
#      然后将快慢指针每次同时往后移动一个结点，直到快指针 fast 到达尾结点。
#
#      此时，慢指针 slow 指向待移除结点的前一个结点，将其移除即可。
#
#
#      快慢指针也可以用来快速找到中间的元素，
#      只要快指针每次移动两个结点，慢指针每次移动一个结点即可。
#
#
#      时间复杂度： O(|head|)
#          1. 需要遍历链表中的全部 O(|head|) 个结点
#      空间复杂度： O(1)
#          1. 只需要维护常数个额外变量


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        # 使用一个哨兵结点，方便处理移除头结点这种边界情况
        head_pre: ListNode = ListNode(0, head)
        # 快指针先网后移动 n 个结点
        fast: ListNode = head_pre
        # 先移动快指针 n 次
        for _ in range(n):
            fast = fast.next

        # 然后同时移动快慢指针，直到快指针到达尾结点
        slow: ListNode = head_pre
        while fast.next:
            fast = fast.next
            slow = slow.next

        # 此时， slow 指向待移除结点的前一个结点，将其移除即可
        slow.next = slow.next.next

        # 返回结果链表的头结点
        return head_pre.next
