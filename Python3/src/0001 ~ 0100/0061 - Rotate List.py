# 链接：https://leetcode.com/problems/rotate-list/
# 题意：给定一个单链表的头结点 head ，
#      将这个单链表向右循环移动 k 次，
#      返回移动后的单链表头结点。


# 数据限制：
#   链表中的结点数在 [0, 500] 内
#   -100 <= Node.val <= 100
#   0 <= k <= 2 * 10 ^ 9


# 输入： head = [1,2,3,4,5], k = 2
# 输出： [4,5,1,2,3]
# 解释： 1 -> 2 -> 3 -> 4 -> 5
#                 ↓
#       5 -> 1 -> 2 -> 3 -> 4
#                 ↓
#       4 -> 5 -> 1 -> 2 -> 3

# 输入： head = [0,1,2], k = 4
# 输出： 4
# 解释： 0 -> 1 -> 2
#            ↓
#       2 -> 0 -> 1
#            ↓
#       1 -> 2 -> 0
#            ↓
#       0 -> 1 -> 2
#            ↓
#       2 -> 0 -> 1


# 思路： 模拟
#
#		可以发现 k 的值很大，而链表的结点数 n 最多只有 500 个，
#       而循环移动每执行 n 次必定会恢复原状。
#
#       所以先统计链表的结点数 n ，然后将 k 对 n 取余，即 k %= n 。
#
#       1. 如果此时 k 为 0 ，那么直接返回原链表。
#       2. 如果此时 k 不为 0 ，那么就将前 n - k 个结点取出，
#           插入到后 k 个结点后面。
#          
#           即先将链表在第 n - k 个结点处断开，
#           再将 head 插入到 tail 后。
#
#
#      时间复杂度：O(n)
#           1. 需要遍历全部 O(n) 个结点
#      空间复杂度：O(1)
#           1. 只需要使用常数个额外变量


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        # 如果链表为空或者 k 为 0 ，直接返回原链表
        if not head or k == 0:
            return head

        # 统计链表的结点数 n
        n: int = 0
        # 记录尾结点 tail
        tail: Optional[ListNode] = head
        # 当前结点还存在时，继续统计
        while tail:
            # 当前结点纳入统计
            n += 1
            # 移动到下一个结点
            tail = tail.next

        # 将 k 对 n 取余
        k %= n
        # 如果 k 为 0 ，直接返回原链表
        if k == 0:
            return head

        # 找到第一部分的尾结点
        first_tail: Optional[ListNode] = head
        # 第一部分只有 n - k 个结点，
        # 所以 first_tail 只需要往后移动 n - k - 1 即可。
        for _ in range(n - k - 1):
            # 移动到下一个结点
            first_tail = first_tail.next

        # 获取第二部分的头结点
        second_head: Optional[ListNode] = first_tail.next
        # 断开链表
        first_tail.next = None
        # 记录第二部分的尾结点 second_tail
        second_tail: Optional[ListNode] = second_head
        # 当下一个结点还存在时，继续移动
        while second_tail.next:
            # 移动到下一个结点
            second_tail = second_tail.next

        # 将第一部分插入到第二部分尾部即可
        second_tail.next = head

        # 第二部分的头结点就是结果链表的头结点
        return second_head
