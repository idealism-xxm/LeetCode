# 链接：https://leetcode.com/problems/count-ways-to-build-rooms-in-an-ant-colony/
# 题意：给定一棵以 0 为根的树，求从根开始构建整个树的不同构建顺序的数量？
#       这个数量要模 1e9 + 7

# 数据限制：
#   n == prevRoom.length
#   2 <= n <= 10 ^ 5
#   prevRoom[0] == -1
#   0 <= prevRoom[i] < n (1 <= i < n)
#   每个和和根连起来节点都是可达的

# 输入： prevRoom = [-1,0,1]
# 输出： 1
# 解释： 有 1 种构建方式
#       0 → 1 → 2

# 输入： prevRoom = [-1,0,0,1,2]
# 输出： 6
# 解释： 有 6 种构建方式
#       0 → 1 → 3 → 2 → 4
#       0 → 2 → 4 → 1 → 3
#       0 → 1 → 2 → 3 → 4
#       0 → 1 → 2 → 4 → 3
#       0 → 2 → 1 → 3 → 4
#       0 → 2 → 1 → 4 → 3

# 思路： 树形 DP + 排列组合
#
#       先根据 preRoom 构建邻接表，然后开始树形 DP ，
#       dfs(i) 返回 Tuple[int, int] ，
#       第一个数 cnt 表示子树的大小，第二个数 num 表示子树不同构建顺序的数量。
#
#       dfs 中，先求出所有子树的大小数组 cnts，及其构建顺序数量数组 nums，
#       然后计算当前子树的大小 cnt = total + 1 = sum(cnts) + 1 ，
#       然后计算对应的数量 num = product(C(sum(cnts[i:]), cnts[i]) * nums[i]) % MOD 。
#
#       使用排列组合计算每个子树构建顺序：
#       先确定第一棵子树每个节点的位置： 从 total 中选取 cnts[0] 个位置，
#           即 C(total, cnts[0]) ，每一种选取方式都有 nums[0] 个构建顺序
#       再确定第二棵子树每个节点的位置： 从剩余 total - cnts[0] 中选取 cnts[1] 个位置，
#           即 C(total - cnts[0], cnts[1])  ，每一种选取方式都有 nums[1] 个构建顺序
#       以此类推
#
#       比赛中有一个大致的思路，但是组合数取模已经太久没使用了，只能使用模版，也没搞明白。。。
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


MOD = 1000000007
N = 100005
fac = [0] * N
ifac = [0] * N


def quick_pow(x: int, y: int) -> int:
    res = 1
    while y > 0:
        if y & 1:
            res = (res * x) % MOD
        x = (x * x) % MOD
        y >>= 1

    return res


def C(n: int, k: int) -> int:
    return fac[n] * ifac[k] % MOD * ifac[n - k] % MOD


def init():
    fac[0] = ifac[0] = 1
    for i in range(1, N):
        fac[i] = fac[i - 1] * i % MOD
        ifac[N - 1] = quick_pow(fac[N - 1], MOD - 2)
    for i in range(N - 2, 0, -1):
        ifac[i] = ifac[i + 1] * (i + 1) % MOD


init()


class Solution:
    def waysToBuildRooms(self, prevRoom: List[int]) -> int:
        n = len(prevRoom)
        # 构建邻接表
        edges: List[List[int]] = [None] * n
        for i in range(n):
            edges[i] = []
        for cur, pre in enumerate(prevRoom[1:], start=1):
            edges[pre].append(cur)

        return self.dfs(edges, 0)[1]

    def dfs(self, edges: List[List[int]], u: int) -> Tuple[int, int]:
        # 先计算所有子树的结果
        cnts, nums = [0] * len(edges[u]), [0] * len(edges[u])
        for i, v in enumerate(edges[u]):
            cnts[i], nums[i] = self.dfs(edges, v)

        # 再计算当前子树的结果
        total, num = sum(cnts), 1
        cnt = total + 1
        for i in range(len(edges[u])):
            # 这里分成两步进行乘法
            num = num * nums[i] % MOD
            num = num * C(total, cnts[i]) % MOD
            total -= cnts[i]

        return cnt, num
