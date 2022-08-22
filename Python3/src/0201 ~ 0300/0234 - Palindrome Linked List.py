# 链接：https://leetcode.com/problems/palindrome-linked-list/
# 题意：给定一个单链表，判断是不是回文的？
#
#      进阶：使用时间复杂度为 O(n) 空间复杂度为 O(1) 的算法求解。


# 数据限制：
#  链表的结点数范围为 [1, 10 ^ 5]
#  0 <= Node.val <= 9


# 输入： head = [1,2,2,1]
# 输出： true

# 输入： head = [1,2]
# 输出： false


# 思路： 快慢指针
#
#      先用快慢指针找到后半部分，然后将后半部分翻转，再对比前后是否相等即可
#
#      时间复杂度： O(n)
#          1. 只需要遍历全部 O(n) 个结点
#      空间复杂度： O(1)
#          1. 只需要使用常数个额外变量即可


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        # 1. 先用快慢指针找到后半部分
        fast: Optional[ListNode] = head
        slow: Optional[ListNode] = head
        while fast:
            # 快指针每次走两步
            fast = fast.next
            if fast:
                fast = fast.next

            # 慢指针每次走一步
            slow = slow.next

        # 此时慢指针就指向后半部分的头结点
        #（如果链表结点数是奇数，那么此时必定是正中间结点的后一个）

        # 2. 翻转后半部分
        slow: Optional[ListNode] = Solution.reverse_list(slow)

        # 3. 对比前后是否相等即可
        l: Optional[ListNode] = head
        r: Optional[ListNode] = slow
        while l and r:
            # 如果值不相等，则必定不是回文链表，直接返回 false
            if l.val != r.val:
                return False

            # 如果值相等，则都移动至下一个结点继续对比
            l = l.next
            r = r.next

        # 所有值都相等，则是回文链表
        return True

    @staticmethod
    def reverse_list(head: Optional[ListNode]) -> Optional[ListNode]:
        # 使用头插法翻转链表
        head_pre: Optional[ListNode] = ListNode(0)
        # 若 head 不是空结点，则继续处理
        while head:
            # 先获取下一个结点
            next: Optional[ListNode] = head.next
            # 再将 head 用头插法放入结果链表中
            head.next = head_pre.next
            head_pre.next = head
            # 接下来处理下一个结点
            head = next

        # 返回翻转后链表的头结点
        return head_pre.next
