# 链接：https://leetcode.com/problems/minimum-operations-to-convert-number/
# 题意：给定 n 个唯一的整数数组 nums 和两个整数 start 和 goal ，
#       现在可以对 x = start 执行任意次以下操作，求做少能变成 goal 的操作次数？
#       如果 0 <= x <= 1000 ，则可以选择一个数 nums[i] ，并设置 x 的值为：
#           1. x = x + nums[i]
#           2. x = x - nums[i]
#           3. x = x ^ nums[i]
#       如果 x < 0 || x > 1000 ，则不能进行任何操作。

# 数据限制：
#   1 <= nums.length <= 1000
#   -10 ^ 9 <= nums[i], goal <= 10 ^ 9
#   0 <= start <= 1000
#   start != goal
#   nums 中的所有数字都是唯一的

# 输入： nums = [1,3], start = 6, goal = 4
# 输出： 2
# 解释： 
#   - 6 ^ 1 = 7
#   - 7 ^ 3 = 4

# 输入： nums = [2,4,12], start = 2, goal = 12
# 输出： 2
# 解释： 
#   - 2 + 12 = 14
#   - 14 - 2 = 12

# 输入： nums = [3,5,7], start = 0, goal = -4
# 输出： 2
# 解释： 
#   - 0 + 3 = 3
#   - 3 - 7 = -4

# 输入： nums = [2,8,16], start = 0, goal = 1
# 输出： -1
# 解释： 
#   无法将 0 变为 1

# 输入： nums = [1], start = 0, goal = 3
# 输出： 3
# 解释： 
#   - 0 + 1 = 1 
#   - 1 + 1 = 2
#   - 2 + 1 = 3


# 思路： BFS
#
#       由于只有 x 在 [0, 1000] 之间时才可以进行操作，
#       那么可以往外扩展的状态只有 [0, 1000] ，这样可以直接 BFS 即可
#
#       由于所有的操作都有对应的相反操作，为了方便我们就从 goal 开始 BFS 。
#       初始队列 q 中只有 goal ，然后当 q 不为空时，一直如下处理：
#           拿出队首数字 cur 及变换到 cur 的最小操作次数 d ，
#           然后遍历选择的数 nums[i] ，分别执行三种操作得到计算结果 nxt ，
#           如果 nxt 在 [0, 1000] 内，
#           则可以记录从 goal 变化到 nxt 的最小操作次数为 d + 1 ，
#           并放入队列 q 中继续进行操作
#
#       时间复杂度： O(n * x) ，其中 x = 1000
#       空间复杂度： O(x)


class Solution:
    def minimumOperations(self, nums: List[int], start: int, goal: int) -> int:
        # 初始化队列，刚开始只有 goal 本身，无需任何操作
        q = deque([(goal, 0)])
        # 初始都没有使用过
        used = [False] * 1001
        # 如果 goal 在 [0, 1000] 内，则标记 goal 已使用过
        if 0 <= goal <= 1000:
            used[goal] = True
        
        # 当队列不为空时，可以继续遍历
        while len(q):
            # 拿出队首数字 cur 及变换到 cur 的最小操作次数 d
            cur, d = q.popleft()
            # 如果已经变换到 start ，
            # 则 d 就是从 start 变换到 goal 的最小操作次数
            if cur == start:
                return d

            # 遍历选择的数字
            for num in nums:
                # 计算下一个数字 cur + num ，
                # 如果在 [0, 1000] 内且没有访问过，则可以放入队列中
                nxt = cur + num
                if 0 <= nxt <= 1000 and not used[nxt]:
                    used[nxt] = True
                    q.append((nxt, d + 1))
                
                # 计算下一个数字 cur - num ，
                # 如果在 [0, 1000] 内且没有访问过，则可以放入队列中
                nxt = cur - num
                if 0 <= nxt <= 1000 and not used[nxt]:
                    used[nxt] = True
                    q.append((nxt, d + 1))
                
                # 计算下一个数字 cur ^ num ，
                # 如果在 [0, 1000] 内且没有访问过，则可以放入队列中
                nxt = cur ^ num
                if 0 <= nxt <= 1000 and not used[nxt]:
                    used[nxt] = True
                    q.append((nxt, d + 1))
        
        # 遍历完所有可能的状态还没有返回，则无法从 start 转换到 goal
        return -1
