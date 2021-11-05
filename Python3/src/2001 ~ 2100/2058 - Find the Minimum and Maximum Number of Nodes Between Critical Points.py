# 链接：https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
# 题意：给定一个单链表，找出两个关键节点最近距离和最远距离，没有则返回 [-1, -1] ，
#       关键节点就是局部最大值或局部最小值，
#           一个节点是局部最大值就是这个节点的值比他左右两边节点的值都大，
#           一个节点是局部最小值就是这个节点的值比他左右两边节点的值都小。

# 数据限制：
#   链表的节点数范围为 [2, 10 ^ 5]
#   1 <= Node.val <= 10 ^ 5

# 输入： nums = [0,1,2]
# 输出： 0
# 解释： 
#   i=0: 0 mod 10 = 0 == nums[0]
#   i=1: 1 mod 10 = 1 == nums[1]
#   i=2: 2 mod 10 = 2 == nums[2]

# 输入： nums = [4,3,2,1]
# 输出： 2
# 解释： 
#   i=0: 0 mod 10 = 0 != nums[0]
#   i=1: 1 mod 10 = 1 != nums[1]
#   i=2: 2 mod 10 = 2 == nums[2]
#   i=3: 3 mod 10 = 3 != nums[3]

# 输入： nums = [2,1,3,5,2]
# 输出： 1
# 解释： 
#   i=0: 0 mod 10 = 0 != nums[0]
#   i=1: 1 mod 10 = 1 == nums[1]
#   i=2: 2 mod 10 = 2 != nums[2]
#   i=3: 3 mod 10 = 3 != nums[3]
#   i=4: 4 mod 10 = 4 != nums[4]


# 思路： 枚举
#
#       维护两个值 left_critical 和 right_critical ，
#           分别表示最左侧的关键节点的下标和最右侧的关键节点的下标，
#       这样我们可以在遍历链表的时候就求出最小距离和最大距离，
#       当遍历到一个新的关键节点时，
#           最小距离： min(当前关键节点的下标 - 最右侧的关键节点的下标)
#           最大距离： max(当前关键节点的下标 - 最左侧的关键节点的下标)
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def nodesBetweenCriticalPoints(self, head: Optional[ListNode]) -> List[int]:
        # 记录最左侧的关键节点的下标和最右侧的关键节点的下标
        left_critical, right_critical = None, None
        # 记录前一个节点的值和当前节点
        pre_val, cur = head.val, head.next
        # 记录当前遍历到的节点下标
        i = 1
        # 记录相邻两个关键节点的最小距离
        min_ans, max_ans = 1000001, -1
        # 如果下一个节点存在，则当前节点可能是关键节点
        while cur.next:
            nxt = cur.next
            # 如果当前节点是局部最大值 或者 局部最小值，则是关键节点
            if (pre_val < cur.val and cur.val > nxt.val) or (pre_val > cur.val and cur.val < nxt.val):
                # 如果当前最右侧的关键节点存在，则更新相邻两个关键节点最小距离
                if right_critical is not None:
                    min_ans = min(min_ans, i - right_critical)
                # 更新最右侧的关键节点的下标
                right_critical = i
                # 如果最左侧的关键节点不存在，则更新为当前节点
                if left_critical is None:
                    left_critical = i
                else:
                    # 如果最左侧的关键节点存在，则可以更新两个关键节点最大距离
                    max_ans = max(max_ans, i - left_critical)

            # 移动到下一个节点
            pre_val = cur.val
            cur = nxt
            i += 1
        
        # 如果最小距离不存在，则设置为 -1
        if min_ans == 1000001:
            min_ans = -1

        return [min_ans, max_ans]
