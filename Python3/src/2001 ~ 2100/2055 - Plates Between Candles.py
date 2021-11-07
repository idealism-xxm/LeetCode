# 链接：https://leetcode.com/problems/plates-between-candles/
# 题意：给定一个只含有 '*' 和 '|' 的字符串 s ，
#       其中 '*' 表示蜡烛， '|' 表示盘子。
#       给定一组查询 queries ， queries[i] = [left_i, right_i] ，
#       针对每个查询求 s[left_i:right_i] 内（含边界）中蜡烛间的盘子数量。

# 数据限制：
#   3 <= s.length <= 10 ^ 5
#   s 只含有 '*' 和 '|'
#   1 <= queries.length <= 10 ^ 5
#   queries[i].length == 2
#   0 <= left_i <= right_i < s.length

# 输入： s = "**|**|***|", queries = [[2,5],[5,9]]
# 输出： [2,3]
# 解释： 
#   - queries[0]: "|**|" 中蜡烛之间有 2 个盘子
#   - queries[1]: "|***|" 中蜡烛之间有 3 个盘子

# 输入： s = "***|**|*****|**||**|*", queries = [[1,17],[4,5],[14,17],[5,11],[15,16]]
# 输出： [9,0,0,0,0]
# 解释： 
#   - queries[0]: "**|**|*****|**||*" 中蜡烛之间有 9 个盘子
#   其他的蜡烛之间都没有盘子


# 思路： 前缀和
#
#       维护一个数组 plate_cnt ，其中 plate_cnt[i] 表示 s[0:i] 中的盘子数量，
#       维护两个数组 left_candle_postion 和 right_candle_position ，
#           left_candle_postion 表示 s[i] 及其左侧的第一个蜡烛的位置，
#           right_candle_position 表示 s[i] 及其右侧的第一个蜡烛的位置。
#
#       然后我们针对每次查询找到范围内左边界及其右侧第一根蜡烛的位置 left ，
#           和右边界及起左侧左侧第一根蜡烛的位置 right 。
#       如果 left 和 right 都在查询范围内，那么它们之间的
#           盘子数量 = plate_cnt[right] - plate_cnt[left]
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def platesBetweenCandles(self, s: str, queries: List[List[int]]) -> List[int]:
        n = len(s)
        # 初始化 s[:i] 中的盘子数量
        plate_cnt = [0] * n
        # 初始化 s[i] 及其左侧的第一个蜡烛的位置
        left_candle_postion = [-1] * n
        # 初始化 s[i] 及其右侧的第一个蜡烛的位置
        right_candle_position = [n + 1] * n
        # 如果第一个是盘子，则 s[:0] 中有 1 个盘子
        if s[0] == '*':
            plate_cnt[0] = 1
        else:
            # 如果第一个是蜡烛，则 s[i] 及其左侧的第一个蜡烛的位置是 0
            left_candle_postion[0] = 0
        # 如果最后一个是蜡烛，则 s[i] 及其右侧的第一个蜡烛的位置是 n - 1
        if s[n - 1] == '|':
            right_candle_position[n - 1] = n - 1

        # 遍历中间的进行更新
        for i in range(1, n):
            # 如果当前是盘子
            if s[i] == '*':
                # s[:i] 中有 plate_cnt[i - 1] + 1 个盘子
                plate_cnt[i] = plate_cnt[i - 1] + 1
                # s[i] 及其左侧的第一个蜡烛的位置是 left_candle_postion[i - 1]
                left_candle_postion[i] = left_candle_postion[i - 1]
            else:
                # 如果当前是蜡烛，则 s[:i] 中有 plate_cnt[i - 1] 个盘子
                plate_cnt[i] = plate_cnt[i - 1]
                # s[i] 及其左侧的第一个蜡烛的位置是 i
                left_candle_postion[i] = i
            
            # 从右往左遍历，更新 right_candle_position
            # 如果当前是盘子
            j = n - 1 - i
            if s[j] == '*':
                # s[i] 及其左侧的第一个蜡烛的位置是 right_candle_position[j + 1]
                right_candle_position[j] = right_candle_position[j + 1]
            else:
                # s[j] 及其左侧的第一个蜡烛的位置是 j
                right_candle_position[j] = j
        
        ans = [0] * len(queries)
        for i in range(len(queries)):
            l, r = queries[i]
            # 找到 s[l:r] 内最左边和最右边的蜡烛的位置
            left, right = right_candle_position[l], left_candle_postion[r]
            # 如果蜡烛都在范围内，则可以计算出盘子数量
            if left < r and right > l:
                ans[i] = plate_cnt[right] - plate_cnt[left]

        return ans
