# 链接：https://leetcode.com/problems/the-number-of-good-subsets/
# 题意：给定一个整数数组 nums ，求满足以下要求的 nums 的子序列的个数？
#       该子序列所有元素的积可以用互不相同的质数相乘得到。

# 数据限制：
#   1 <= nums.length <= 10 ^ 5
#   1 <= nums[i] <= 30

# 输入： nums = [1,2,3,4]
# 输出： 6
# 解释：
#   - [1,2]: 乘积是 2 ，含有互不相同的质数 2
#   - [1,2,3]: 乘积是 6 ，含有互不相同的质数 2 和 3
#   - [1,3]: 乘积是 3 ，含有互不相同的质数 3
#   - [2]: 乘积是 2 ，含有互不相同的质数 2
#   - [2,3]: 乘积是 6 ，含有互不相同的质数 2 和 3
#   - [3]: 乘积是 3 ，含有互不相同的质数 3

# 输入： nums = [4,2,3,15]
# 输出： 5
# 解释： 
#   - [2]: 乘积是 2 ，含有互不相同的质数 2
#   - [2,3]: 乘积是 6 ，含有互不相同的质数 2 和 3
#   - [2,15]: 乘积是 30 ，含有互不相同的质数 2, 3 和 5
#   - [3]: 乘积是 3 ，含有互不相同的质数 3
#   - [15]: 乘积是 15 ，含有互不相同的质数 3 和 5


# 思路： 状压 DP
#
#       可以发现 nums[i] 在 30 以内，而这里面的质数只有 2, 3, 5, 7, 11, 13, 17, 19, 23, 29 ，
#       所以我们可以先离散化这些质数 prime_to_bit ，
#       并预处理所有合法的数及其对应的质数压缩后的状态 valid_num_to_bits （所有合法的数，以及其所有质因子下标），
#
#       然后我们统计所有在 valid_num_to_bits 内的数，其他的数必定不合法（因子至少存在一个质数出现多次）
#       如果只有一个数，且该数是 1 ，则不存在结果，
#       
#       题目要求所有的质因子最多只有一个，所以每个合法的数最多只能选择一次，
#       那么我们需要先枚举要选择的数，然后再枚举压缩后的状态，
#       dp[i] 表示选择了 i 中下标对应的质数时，所有合法的方法的个数，
#       初始化： dp[i] = 0, dp[0] = 1 （未选择任何数只有一种方法）
#       状态转移：
#           当前要考虑的合法的数为 num ，其出现次数为 cnt ，其对应质因子状态为 bits ，
#           当前状态为 i ，
#           那么当且仅当 bits & i == 0 时可以进行状态转移，因为它们没有相同的质数，
#           （cnt 表示我们可以选择 cnt 个 num ，都能转移到该状态上）
#           dp[i | bits] = dp[i | bits] + dp[i] * cnt
#
#       最后还要特殊处理 1 的个数，因为 1 不影响结果，所有 1 的所有子集都可以选择，
#       即最终结果为 sum(dp[1:]) * pow(2, one_cnt)
#
#       时间复杂度： O(k * (2 ^ k)) ， k = 10 ，为质因子的个数
#       空间复杂度： O(2 ^ k)


MOD = 1000000007
# 30 以内的质数
primes = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29]
cnt = len(primes)
mx = 1 << cnt
# 离散化
prime_to_bit = {prime: bit for bit, prime in enumerate(primes)}

# 获取一个数的质因子状态表示（如果一个质因子出现多次，则返回 0 ）
def get_bits(num: int):
    bits = 0
    # 质因数分解
    for bit, prime in enumerate(primes):
        if num % prime == 0:
            # 加入该质因数的状态
            bits |= 1 << bit
            num //= prime
            # 如果还含有 prime 因子，那么直接返回 0
            if num % prime == 0:
                return 0
    return bits

# 30 以内合法的数（每个质因子只出现一次）
valid_num_to_bits = {}
for i in range(2, 31):
    bits = get_bits(i)
    if bits:
        valid_num_to_bits[i] = bits

class Solution:
    def numberOfGoodSubsets(self, nums: List[int]) -> int:
        # 统计合法的数的个数
        num_to_cnt = defaultdict(int)
        for num in nums:
            if num == 1 or num in valid_num_to_bits:
                num_to_cnt[num] += 1
        
        if len(num_to_cnt) == 0:
            return 0
        # 如果只有一个数，且该数是 1 ，则不存在结果
        if len(num_to_cnt) == 1 and 1 in num_to_cnt:
            return 0
        # 1 的数量需要特殊计算，因为 1 不是质数
        # 最后的结果要乘以 2 ^ one_cnt ，因为 1 不影响结果，
        # 所以可能的结果还需要配合 1 的所有子集
        one_cnt = num_to_cnt.pop(1, 0)

        # dp[i] 表示选择了 i 中下标对应的质数时，所有合法的方法的个数
        dp = [0] * mx
        # dp[0] 表示没有选择任何数，只有一种方法
        dp[0] = 1
        for num, num_cnt in num_to_cnt.items():
            # 获取 num 对应的质因子状态 bits
            bits = valid_num_to_bits[num]
            for i in range(mx):
                # 如果两个状态没有相同的质因子，则可以进行状态转移
                if dp[i] and i & bits == 0:
                    # 所有 num_cnt 个 num 都可以被选择一次
                    dp[i | bits] = (dp[i | bits] + dp[i] * num_cnt) % MOD

        return (sum(dp[1:]) * pow(2, one_cnt, MOD)) % MOD
