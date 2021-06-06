# 链接：https://leetcode.com/problems/process-tasks-using-servers/
# 题意：给定一组服务器的权重 servers, servers[i] 表示第 i 个服务器的权重，
#      还有一组待执行的任务 tasks, tasks[j] 表示第 j 个任务执行用时，
#      第 j 个任务最早可在第 j 秒执行，每个任务可执行时，
#      会从所有服务器中选择权重最大的，权重相同时选择下标最小的。
#      同一秒有多个任务可执行时，优先执行下标最小的。
#      求每个任务在哪台服务器上执行？

# 输入： servers = [3,3,2], tasks = [1,2,3,2,1,2]
# 输出： [2,2,0,2,1,2]
# 解释： 0 秒：任务 0 在服务器 2 上执行，直到 1 秒时结束
#       1 秒：服务器 2 可用
#             任务 1 在服务器 2 上执行，直到 3 秒时结束
#       2 秒：任务 2 在服务器 0 上执行，直到 5 秒时结束
#       3 秒：服务器 2 可用
#             任务 3 在服务器 2 上执行，直到 5 秒时结束
#       4 秒：任务 4 在服务器 1 上执行，直到 5 秒时结束
#       5 秒：所有服务器均可用
#             任务 5 在服务器 2 上执行，直到 7 秒时结束

# 输入： servers = [5,1,4,3,2], tasks = [2,1,2,4,5,2,1]
# 输出： [1,4,1,4,1,3,2]
# 解释： 0 秒：任务 0 在服务器 1 上执行，直到 2 秒时结束
#       1 秒：任务 1 在服务器 4 上执行，直到 2 秒时结束
#       2 秒：服务器 1 和 4 可用
#             任务 2 在服务器 1 上执行，直到 4 秒时结束
#       3 秒：任务 3 在服务器 4 上执行，直到 7 秒时结束
#       4 秒：服务器 1 可用
#             任务 4 在服务器 1 上执行，直到 9 秒时结束
#       5 秒：任务 5 在服务器 3 上执行，直到 7 秒时结束
#       6 秒：任务 6 在服务器 2 上执行，直到 7 秒时结束

# 思路： 优先队列
#
#       按照题意模拟即可，有两点需要注意：
#           1. 需要用优先队列维护 可用的服务器 和 使用中的服务器，
#               方便每次在 O(logn) 内找到符合的服务器
#           2. 维护当前秒数，需要跳跃式更新当前秒数，不能每次 +1 ，否则会超时
#
#       优先队列 available_servers 存储当前可用的服务器，按照 (权重，下标) 存储，
#           这样可以每次在 O(logn) 内找到满足要求的服务器，进行操作
#       优先队列 running_servers 存储当前正在使用的服务器，按照 (可用时的时间，下标) 存储
#           这样每次开始时可以在 O(logn) 内找到运行完的服务器，添加到 available_servers 中
#
#       然后遍历任务即可，每次还要维护当前秒数 now = max(now, j) ，
#           如果没有可用服务器，则 now 为第一个服务器运行完成的时间点
#
#       时间复杂度： O(mlogn)
#       空间复杂度： O(n)


class Solution:
    def assignTasks(self, servers: List[int], tasks: List[int]) -> List[int]:
        available_servers = []
        running_servers = []
        ans = [0] * len(tasks)
        # 开始时，所有的服务器都可用
        for i, server in enumerate(servers):
            heapq.heappush(available_servers, (server, i))

        now = 0
        for i, task in enumerate(tasks):
            # 第 i 个任务，至少要第 i 秒才能执行
            now = max(now, i)
            # 找到所有运行完成的服务器，放入可用服务器列表中
            while len(running_servers) > 0:
                # 找到第一个将可用的服务器
                info = heapq.heappop(running_servers)
                # 如果还未到时间点，则放回，跳出循环
                if info[0] > now:
                    heapq.heappush(running_servers, info)
                    break
                # 到时间点了，则放入可用服务器列表中
                heapq.heappush(available_servers, (servers[info[1]], info[1]))

            # 如果当前时间点无可用服务器，则需要跳跃到服务器可用的第一个时间点，
            # 将此时所有可用的服务器取出来
            if len(available_servers) == 0:
                # 跳跃到下一个时间点
                info = heapq.heappop(running_servers)
                heapq.heappush(running_servers, info)
                now = info[0]

                # 找到此时所有运行完成的服务器，放入可用服务器列表中
                while len(running_servers) > 0:
                    # 找到第一个将可用的服务器
                    info = heapq.heappop(running_servers)
                    # 如果还未到时间点，则放回，跳出循环
                    if info[0] > now:
                        heapq.heappush(running_servers, info)
                        break
                    # 到时间点了，则放入可用服务器列表中
                    heapq.heappush(available_servers, (servers[info[1]], info[1]))

            # 找到第一个可用的服务器，跑任务 i
            weight, index = heapq.heappop(available_servers)
            ans[i] = index
            # 将当前服务器的信息放入 running_servers 中
            heapq.heappush(running_servers, (now + task, index))

        return ans
