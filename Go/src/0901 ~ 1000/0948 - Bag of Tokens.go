// 链接：https://leetcode.com/problems/bag-of-tokens/
// 题意：你的初始能量为 power ，初始分数为 0 。
//      有一个长度为 n 的数组 tokens ，其中 tokens[i] 表示第 i 个硬币的值。
//      现在你需要通过以下操作，最大化自己的得分：
//          1. 如果 power >= tokens[i] ，则可以将第 i 个硬币正面朝上放置，
//              此时 power 减少 tokens[i] ，得分 score 增加 1 。
//          2. 如果 score > 0 ，则可以将第 i 个硬币反面朝上放置，
//              此时 power 增加 tokens[i] ，得分 score 减少 1 。
//
//      每一个硬币最多只能放置一次，且无需放置全部硬币，求最大能获得的分数？


// 数据限制：
//  0 <= tokens.length <= 1000
//  0 <= tokens[i], power < 10 ^ 4


// 输入： tokens = [100], power = 50
// 输出： 0
// 解释： 初始能量太小，不足以正面朝上放置第 0 个硬币。
//       初始分数太小，不足以反面朝上放置第 0 个硬币。

// 输入： tokens = [100,200], power = 150
// 输出： 1
// 解释： 正面朝上放置第 0 个硬币， power 减少至 50 ，score 增加至 1 。
//       不再放置其他硬币。

// 输入： tokens = [100,200,300,400], power = 200
// 输出： 2
// 解释： 1. 正面朝上放置第 0 个硬币， power 减少至 100 ，score 增加至 1 ；
//       2. 正面朝上放置第 3 个硬币， power 增加至 500 ，score 减少至 0 ；
//       3. 正面朝上放置第 1 个硬币， power 减少至 300 ，score 增加至 1 ；
//       4. 正面朝上放置第 2 个硬币， power 减少至 0 ，  score 增加至 2 。


// 思路： 贪心 + 排序 + 双指针
//
//      每个硬币正面朝上放置时，都增加 1 分，但需要扣除 tokens[i] 点能量，
//      那么我们必定要贪心地正面朝上放置值最小的那个硬币，尽可能减小能量损耗。
//
//      类似地，我们必定要贪心地反面朝上放置值最大的那个硬币，
//      尽可能获得更多能量。
//
//      所以我们可以先对 tokens 按升序排序，然后使用双指针的方式进行处理。
//
//      l 指向剩余硬币中值最小的那个硬币，初始化为 0 ；
//      r 指向剩余硬币中值最大的那个硬币，初始化为 len(tokens) - 1 。
//
//      同时我们需要维护操作过程中的分数 score ，以及所有情况下得分的最大值 ans 。
//
//      只要还有能放置的硬币 (l <= r) ，且还能放置硬币 (power >= tokens[l] || score > 0) ，
//      那么可以进行如下处理：
//          1. 先不断贪心地正面朝上放置第 l 个硬币，直至能量不足正面放置下一个硬币
//          2. 此时停止操作的话，我们能获得 score 分，可以更新 ans 的最大值
//          3. 如果还能反面朝上放置第 r 个硬币 (score > 0) ，
//              则放置以获取能量，然后不断重复上述操作
//
//      
//      时间复杂度：O(nlogn)
//          1. 需要对 tokens 中全部 O(n) 个数字进行排序，时间复杂度为 O(nlogn)
//          2. 需要遍历 tokens 中全部 O(n) 个数字一次
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量即可


func bagOfTokensScore(tokens []int, power int) int {
    if len(tokens) == 0 {
        return 0
    }

    // 按升序排序
    sort.Ints(tokens)
    // score 维护操作过程中的分数
    score := 0
    // ans 维护所有情况下得分的最大值
    ans := score
    // l 指向剩余硬币中值最小的那个硬币，初始化为 0 ；
    // r 指向剩余硬币中值最大的那个硬币，初始化为 len(tokens) - 1
    l, r := 0, len(tokens) - 1
    // 如果有能放置的硬币 (l <= r) ，
    // 且还能放置硬币 (power >= tokens[l] || score > 0) ，
    // 则继续循环处理
    for l <= r && (power >= tokens[l] || score > 0) {
        // 不断贪心地正面朝上放置第 l 个硬币，直至能量不足正面放置下一个硬币
        for l <= r && power >= tokens[l] {
            power -= tokens[l]
            score += 1
            l += 1
        }
        // 此时停止操作的话，我们能获得 score 分，可以更新 ans 的最大值
        ans = max(ans, score);
        // 如果还能反面朝上放置第 r 个硬币 (score > 0) ，则放置以获取能量
        if l <= r && score > 0 {
            power += tokens[r]
            score -= 1
            r -= 1
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
