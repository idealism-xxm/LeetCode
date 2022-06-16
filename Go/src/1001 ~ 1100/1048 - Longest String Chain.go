// 链接：https://leetcode.com/problems/longest-string-chain/
// 题意：给定一个单词列表 words ，如果一个单词 a 能通过恰好插入一个字母得到单词 b ，
//      那么 a 就是 b 的前驱单词。
//      求单词列表中的单词能形成的最长的单词链的长度？
//
//      单词链 [word_1, word_2, ..., word_k] 满足一下条件：
//          1. k >= 1
//          2. word_i 是 word_(i+1) 的前驱单词


// 数据限制：
//  1 <= words.length <= 1000
//  1 <= words[i].length <= 16
//  word 仅含有英文小写字母


// 输入： words = ["a","b","ba","bca","bda","bdca"]
// 输出： 4
// 解释： 一个最长的单词链是 ["a","ba","bda","bdca"]

// 输入： words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
// 输出： 5
// 解释： 一个最长的单词链是 ["xb", "xbc", "cxbc", "pcxbc", "pcxbcf"]

// 输入： words = ["abcd","dbqca"]
// 输出： 1
// 解释： 一个最长的单词链是 ["abcd"]


// 思路： DP
//
//      设 dp[i] 表示表示以 word 为结尾的最长单词链的长度。
//      由于 i 是单词，所以 dp 使用 map 维护，初始化为空 map 。
//
//      对于单词 b 和其前驱单词 a ，必有以下关系： len(a) + 1 = len(b) ，
//      所以我们只需要按照单词长度排序，然后进行状态转移即可。
//
//      枚举当前单词 word ，则此时所有长度小于 len(word) 的单词状态都已确定。
//
//      此时枚举单词 word 的前驱单词 pre = word[..i] + word[i+1..] ，
//      则有 dp[word] = dp[pre] + 1 。
//
//      注意，如果 pre 不在单词列表 words 中，则 dp[pre] 的值取 0 ，
//      这样就能兼顾 word 是单词链中第一个单词的情况。
//
//
//      设单词的长度为 L 。
//
//      时间复杂度：O(nlogn + n * L ^ 2)
//          1. 需要对单词列表 word 进行排序，时间复杂度为 O(nlogn)
//          2. 需要遍历全部 O(n) 个单词，每次要枚举其全部 O(L) 个前驱单词，
//              然后还要在 O(L) 内从 dp 中找到 pre 对应的状态，
//              总时间复杂度为 O(n * L ^ 2)
//      空间复杂度：O(n * L)
//          1. 需用维护全部 O(n) 个状态，每个状态都包含长度为 O(L) 的单词


func longestStrChain(words []string) int {
    // dp[word] 表示以 word 为结尾的最长单词链的长度
    dp := make(map[string]int, len(words))
    // 初始化最长单词链的长度为 0
    ans := 0
    // 按照字符串长度升序排序
    sort.SliceStable(words, func(i, j int) {return len(words[i]) < len(words[j])})
    // 从长度小的字符串开始枚举，这样处理能保证所有依赖的状态都已确定，
    // 能挣钱进行状态转移。
    // 因为长度为 len(word) 的单词状态仅依赖长度为 len(word) - 1 的单词状态
    for _, word := range words {
        // 枚举 word 所有可能的前驱单词，找到其中最长单词链的长度
        maxLen := 0
        for i := range word {
            pre := word[:i] + word[i + 1:]
            maxLen = max(maxLen, dp[pre])
        }
        // 转移至以 word 为结尾的最长单词链的长度
        dp[word] = maxLen + 1
        // 更新最长单词链的长度
        ans = max(ans, maxLen + 1)
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
