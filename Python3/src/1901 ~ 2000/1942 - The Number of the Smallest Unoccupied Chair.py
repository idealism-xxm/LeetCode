# 链接：https://leetcode.com/problems/the-number-of-the-smallest-unoccupied-chair/
# 题意：有 n 把椅子，有 n 个人会来坐，现在给出这 n 个人的到达和离开时间数组 times ，
#       当一个人到达时，会坐到编号最小的椅子上，当一个人离开时，会立刻释放椅子。
#       求第 targetFriend 个人坐在哪把椅子上？


# 数据限制：
#   n == times.length
#   2 <= n <= 10 ^ 4
#   times[i].length == 2
#   1 <= arrival_i < leaving_i <= 10 ^ 5
#   0 <= targetFriend <= n - 1
#   arrival_i 都是唯一的

# 输入： times = [[1,4],[2,3],[4,6]], targetFriend = 1
# 输出： 1
# 解释： 
#       人 0 在时刻 1 坐在椅子 0 上，
#       人 1 在时刻 2 坐在椅子 1 上，   
#       人 1 在时刻 3 离开椅子 1 ，
#       人 0 在时刻 4 离开椅子 0 ，
#       人 2 在时刻 4 坐在椅子 0 上。   
#       所以答案是人 1 坐在椅子 1 上。

# 输入： times = [[3,10],[1,5],[2,6]], targetFriend = 0
# 输出： 2
# 解释： 
#       人 1 在时刻 1 坐在椅子 0 上，
#       人 2 在时刻 2 坐在椅子 1 上，
#       人 0 在时刻 3 坐在椅子 2 上，   
#       人 1 在时刻 5 离开椅子 0 ，
#       人 2 在时刻 6 离开椅子 1 ，
#       人 0 在时刻 10 坐在椅子 2 上。   
#       所以答案是人 0 坐在椅子 2 上。


# 思路： 优先队列
#
#       用优先队列维护当前可用的椅子，再用一个数组维每个时间点可以释放的椅子列表，
#
#       我们对 times 按照到达时间升序排序，然后遍历，
#       每一次优先将当前时间点以前所有的椅子释放，
#       然后从优先队列中获取编号最小的椅子，
#       并放入对应离开时间点的释放列表中，
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(n)

from collections import defaultdict
import queue

class Solution:
    def smallestChair(self, times: List[List[int]], targetFriend: int) -> int:
        # 由于到达时间点都不同，所以将人编号转换为到达时间点
        target_arrival = times[targetFriend][0]

        # 按照到达时间升序排序
        def cmp(a, b):
            return a[0] - b[0]
        times.sort(key=cmp_to_key(cmp))

        # 记录上次释放椅子的时间点
        last_time = 0
        # 优先队列维护所有可用的椅子
        available = queue.PriorityQueue()
        for i in range(len(times)):
            available.put(i)
        # 记录每个时间点可以释放的椅子
        occupied = defaultdict(list)
        # 遍历每个人列表
        for arrival, leave in times:
            # 将当前时间点及以前的椅子全部释放
            while last_time <= arrival:
                for chair_num in occupied[last_time]:
                    available.put(chair_num)
                last_time += 1
            
            # 获取最小的椅子编号
            chair_num = available.get()
            if arrival == target_arrival:
                return chair_num
            
            # 放入待释放的椅子列表中
            occupied[leave].append(chair_num)
        
        # 不会走到这里
        return None
