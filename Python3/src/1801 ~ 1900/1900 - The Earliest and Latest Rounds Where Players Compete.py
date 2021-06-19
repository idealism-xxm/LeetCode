# 链接：https://leetcode.com/problems/the-earliest-and-latest-rounds-where-players-compete/
# 题意：有 n 个运动员参加一项比赛，运动员下标从 1 开始计算，
#       每一轮比赛是正数第 i 个运动员和倒数第 i 个运动员比赛，胜者晋级到下一轮，
#       奇数个运动员时，最中间的那个运动员自动晋级下一轮，
#       现在求第 firstPlayer 个运动员和第 secondPlayer 个运动员
#       最早和最迟在第几轮比赛？

# 数据限制：
#   2 <= n <= 28
#   1 <= firstPlayer < secondPlayer <= n

# 输入： n = 11, firstPlayer = 2, secondPlayer = 4
# 输出： [3, 4]
# 解释： 一种最早比赛的序列：
#           第一轮: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
#           第二轮: 2, 3, 4, 5, 6, 11
#           第三轮: 2, 3, 4
#       一种最迟比赛的序列：
#           第一轮: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
#           第二轮: 1, 2, 3, 4, 5, 6
#           第三轮: 1, 2, 4
#           第四轮: 2, 4

# 输入： n = 5, firstPlayer = 1, secondPlayer = 5
# 输出： [1, 1]
# 解释： 第 1 号运动员和第 5 号运动员，第一轮就需要比赛

# 思路1： 递归 + 记忆化
#
#       可以发现数据很小，所以可以尝试使用枚举的方式
#       我们枚举每一轮胜负情况（保证指定的两个人必胜），
#       那么第一轮需要遍历 O(2 ^（(n // 2) - 2)) 次胜负情况，
#           然后将指定的两个人的编号转换成新一轮对应的编号，
#           这样就转换为子问题求解
#       那么第二轮需要遍历 O(2 ^ ((n // 4) - 2)) 次胜负情况，以此类推...
#       综上：时间复杂度为 O(2 ^（(n // 2) - 2) + (n // 4) - 2 + ... + (2 - 2))
#                       = O(2 ^ (n - 2 - 2ceil(logn))
#       由于 n 最大为 28 ，所以时间复杂度近似为： O(2 ^ (28 - 2 - 2ceil(log28))) = O(2 ^ 18)
#       可以直接遍历所有轮次胜负情况，并记忆化即可
#
#       这样可以使用三维数组 dp[n][first][second] 表示
#       有 n 个运动员，第 first 号运动员和第 second 号运动员最早和最迟的比赛时间
#
#       没想好如何使用迭代枚举胜负情况，就是用了递归的方式，
#       配合记忆化就可以
#
#       比赛时候想到的解法，似乎有点复杂，赛后才 AC
#
#       时间复杂度： O(2 ^ (n - 2 - 2ceil(logn)))
#       空间复杂度： O(n ^ 3)


dp = {}


class Solution:
    def earliestAndLatest(self, n: int, firstPlayer: int, secondPlayer: int) -> List[int]:
        arr = [True] * n
        return self.dfs(arr, 0, firstPlayer - 1, secondPlayer - 1)

    def dfs(self, arr: List[bool], cur: int, first: int, second: int) -> List[int]:
        n = len(arr)
        key = (n, first, second)
        # 只在刚开始时判断，如果第一次不成功，后续都不会成功
        if cur == 0:
            # 如果第一轮就就会遇上，则必定是第一轮就结束
            if first == n - second - 1:
                dp[key] = [1, 1]
            # 如果以前遇到过相关的情况，则直接返回
            if dp.get(key):
                return dp.get(key)

        # 如果已过半（奇数时则已到最中间），则当前所有的胜负情况已知晓
        # 可以开始下一轮
        if cur >= (n >> 1):
            # 计算 first 和 second 下一轮对应的位置
            nxt_first = sum((int(ok) for ok in arr[:first]))
            nxt_second = sum((int(ok) for ok in arr[:second]))
            # 计算下一轮比赛人数
            nxt_n = (n + 1) >> 1
            # 初始化胜负数组
            nxt_arr = [True] * nxt_n
            # 计算下一轮的结果，并记忆化
            earliest, latest = self.dfs(nxt_arr, 0, nxt_first, nxt_second)
            dp[(nxt_n, nxt_first, nxt_second)] = [earliest, latest]
            # 加上当前轮次
            return [earliest + 1, latest + 1]

        earliest, latest = n, 0
        # cur 既不是 first 也不是 second 时， cur 可以输
        if cur not in [first, second]:
            arr[cur] = False
            nxt_earliest, nxt_latest = self.dfs(arr, cur + 1, first, second)
            arr[cur] = True
            earliest = min(earliest, nxt_earliest)
            latest = max(latest, nxt_latest)
        # cur 的对手既不是 first 也不是 second 时， cur 的对手可以输
        if n - cur - 1 not in [first, second]:
            arr[n - cur - 1] = False
            nxt_earliest, nxt_latest = self.dfs(arr, cur + 1, first, second)
            arr[n - cur - 1] = True
            earliest = min(earliest, nxt_earliest)
            latest = max(latest, nxt_latest)

        return [earliest, latest]


# 思路2： 递归 + 记忆化
#
#       赛后看到有更巧妙的解法，虽然也是枚举，但是合并了相同位置的情况，时间复杂度降低为：
#       O(n ^ 4 * logn)
#
#       我们定义 dfs(n, l, r) 表示 n 个运动员，第一个运动员距离左侧第 l 位，
#       第二个运动员是距离右侧第 r 位，
#       返回值为这种情况下最早和最迟遇见的轮数，
#
#       内部我们按照如下方式进行遍历，在下一轮中，有如下情况：
#           第一个运动员可以是左侧第 nxt_l in [1, l] 位，
#           第二个运动员可以是右侧第 nxt_r in [1, r] 位，
#               同时由于每一对只能胜出一个，所以第一个运动员的 nxt_l
#               可以直接进一步限制第二个运动员 nxt_r 的上下限
#               第一个运动员下次序号为 nxt_l ，
#               则表明其左侧晋级了 nxt_l - 1 人，淘汰了 l - nxt_l + 1 人，
#               即第二个运动员右侧至少晋级了 l - nxt_l + 1 人，至少淘汰了 nxt_l - 1 人
#               即 nxt_r 的下限为 l - nxt_l + 1 ，上限为 r - nxt_l + 1
#       注意还要保证:
#           总共晋级人数不超过 nxt_n ，即 nxt_l + nxt_r <= nxt_n
#           总共淘汰人数不超过 n - nxt_n ，
#               即 l - nxt_l + r - nxt_r <= n - nxt_n
#                   => l + r <= n - nxt_n + nxt_l + nxt_r
#                   => l + r + nxt_n - n <= nxt_l + nxt_r
#       综上： l + r + nxt_n - n <= nxt_l + nxt_r <= nxt_n
#       为了方便计算，并减少一点时间复杂度，我们可以令 l <= r
#
#       时间复杂度： O(n ^ 4 * logn)
#       空间复杂度： O(n ^ 3)


class Solution:
    def earliestAndLatest(self, n: int, firstPlayer: int, secondPlayer: int) -> List[int]:
        return self.dfs(n, firstPlayer, n - secondPlayer + 1)

    @lru_cache(None)
    def dfs(self, n: int, l: int, r: int) -> List[int]:
        # 保证 l <= r
        if l > r:
            return self.dfs(n, r, l)

        # 如果次序一样，则本轮就需要比赛
        if l == r:
            return [1, 1]

        nxt_n = (n + 1) >> 1
        earliest, latest = n, 0
        # 枚举第一个运动员下次序号
        for nxt_l in range(1, l + 1):
            # 枚举第二个运动员下次序号，注意 nxt_l 决定了其上下限
            for nxt_r in range(l - nxt_l + 1, r - nxt_l + 1):
                if l + r + nxt_n - n <= nxt_l + nxt_r <= nxt_n:
                    # 获取下一轮的结果
                    nxt_earliest, nxt_latest = self.dfs(nxt_n, nxt_l, nxt_r)
                    # 比较更新时，需要加上本轮
                    earliest = min(earliest, nxt_earliest + 1)
                    latest = max(latest, nxt_latest + 1)

        return [earliest, latest]


# 思路3： 贪心 TODO
#
#       看到有更好的解法，直接使用贪心就可以在 O(logn) 内解出来：
#       1. https://github.com/wisdompeak/LeetCode/tree/master/Dynamic_Programming/1900.The-Earliest-and-Latest-Rounds-Where-Players-Compete
#       2. https://github.com/wisdompeak/lc-score-board/tree/gh-pages/generateEXCEL/Data/Cup/Round_4
#
#       不过这种解法需要考虑的情况太多，就不涉及了（详情可看以上解法）
#
#       为了方便操作：我们可以让 first 在中点左侧（如果在右侧，则可以对称转换），
#           且 first 离边界比 second 离边界更近
#       1. 特殊情况： first 和 second 第一轮就相遇，直接返回
#       2. 找到最迟比赛轮数
#           为了让 first 和 second 最迟比赛，那么最优策略下有以下两种情况：
#               (1) 初始状态 second 在中点左侧（包含中点），则会剩余 2 人
#                   操作如下：让 second 左侧的人都输掉，然后 first 和 second 排在第一和第二，直至剩下 2 人
#                   计算方式：ceil(logn)
#               (2) 初始状态 second 在中点右侧，则会剩余至少 3 人
#                   操作如下： first 左侧还有人时，第一轮其左侧人全输，
#                       ① 从第一轮开始，每一轮让 second 右侧的人输掉一个，
#                           直至 second 右侧全部输掉，所以至多需要 n - second + 1 轮
#                       ② second 比较靠近中点，所以最后可以剩下 3 人，
#                           所以最多需要 ceil(logn) 轮
#                   计算方式： min(n - second + 1, ceil(logn))
#
#           针对第 (1) 种情况， n - second + 1 >= ceil(logn)
#           综上：最迟比赛轮数 = min(n - second + 1, ceil(logn))
#
#       3. 找到最早比赛轮数
#           太难了，选择放弃
#
#
#
#       时间复杂度： O(logn)
#       空间复杂度： O(1)


class Solution:
    def earliestAndLatest(self, n: int, firstPlayer: int, secondPlayer: int) -> List[int]:
        # 特殊情况，第一轮就相遇
        if firstPlayer + secondPlayer == n + 1:
            return [1, 1]

        # 保证 first 在左侧，且 first 离边界比 second 离边界更近
        if firstPlayer + secondPlayer > n + 1:
            firstPlayer, secondPlayer = n - secondPlayer + 1, n - firstPlayer + 1
            pass