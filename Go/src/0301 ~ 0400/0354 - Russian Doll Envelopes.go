// 链接：https://leetcode.com/problems/russian-doll-envelopes/
// 题意：给定一个二维数组 envelopes ，
//      其中 envelopes[i] = [w_i, h_i] 表示第 i 张信封的宽和高。
//
//      当且仅当一个信封的宽和高都大于另一个信封的宽和高时，后者能被放入前者中，
//      求最多能嵌套多少层信封。
//
//      注意：不能旋转信封


// 数据限制：
//  1 <= envelopes.k <= 10 ^ 5
//  envelopes[i].k == 2
//  1 <= w_i, h_i <= 10 ^ 5


// 输入： envelopes = [[5,4],[6,4],[6,7],[2,3]]
// 输出： 3
// 解释： 最多能嵌套 3 层信封 ([2,3] => [5,4] => [6,7])

// 输入： envelopes = [[1,1],[1,1],[1,1]]
// 输出： 1


// 思路： DP + 二分
//
//      本题是最长上升子序列 (LIS) 的加强版，需要保证子序列的两个值都是严格递增的。
//
//
//      最简单地求普通 LIS 就是普通 DP ：
//          设 dp[i] 为以第 i 个元素为结尾的最长上升子序列的长度，
//          那么在更新 dp[i] 时，需要更新 dp[i] = max(dp[j]) + 1 ，
//          其中 j < i，且 s[j] < s[i] 。
//
//      时间复杂度为 O(n ^ 2) ，空间复杂度为 O(n)
//
//      可以使用二分将时间复杂度优化为 O(nlogn) ，只需要注意到求解过程中的内在约束：
//          设 minHeight[k] 表示最长上升子序列的长度为 k 时的最小高度，
//          为了方便后续处理，初始化 minHeight = [0] ，
//          表示最长上升子序列的长度为 0 时的最小高度为 0 。
//
//          那么我们在求解过程中维护的 minHeight 必定是一个严格递增的数组。
//
//      注意到这个约束后，我们就不需要遍历前面求出的全部状态 dp[j] ，
//      只需要在 minHeight 中找到第一个大于等于当前高度的下标 k 即可。
//
//      此时 k 就是以当前高度为结尾的最长上升子序列的长度：
//          1. len(minHeight) == k: 
//              说明长度为 k 的最长上升子序列是第一次出现，
//              直接将当前高度加入 minHeight 中即可
//          2. len(minHeight) > k:
//              说明长度为 k 的最长上升子序列已经出现过了，
//              由于二分找到的是第一个大于等于当前高度的下标 k ，
//              所以必定有 minHeight[k] >= 当前高度 ，
//              可以直接更新 minHeight[k] 为当前高度
//
//      最后 len(minHeight) - 1 就是最长上升子序列的长度。
//
//
//      针对本题需要先对 envelopes 按照宽度升序排序，宽度相同时按照高度将序排序。
//          1. 宽度不同时，按照宽度升序排序，
//              保证按照顺序遍历时，宽度是递增的，
//              这样基本就转换成了普通的 LIS ，后续只需要二分高度
//          2. 宽度相同时，按照高度降序排序，结合宽度递增，就严格转换成了普通的 LIS ，
//              保证相同宽度的信封不会嵌套。例如：
//              (1) 遍历顺序为 [[3,3],[3,4]] 时，就会出现嵌套的情况，
//                  因为只考虑了高度递增，没有保证宽度递增。
//              (2) 遍历顺序为 [[3,4],[3,3]] 时，就不会出现嵌套的情况，
//                  因为是按照高度递减处理的，而处理时会按照高度递增处理。
//
//
//      时间复杂度：O(nlogn)
//          1. 需要对 envelopes 排序，时间复杂度为 O(nlogn)
//          2. 需要遍历 envelopes 全部 O(n) 个元素，
//              且每次遍历时都需要进行 O(logn) 的二分
//      空间复杂度：O(n)
//          1. 需要维护一个大小为 O(n) 的数组 minHeight
//          2. 需要对 envelopes 排序，不改变入参的情况下需要 O(n) 的空间


import "sort"


func maxEnvelopes(envelopes [][]int) int {
    // 按照宽度升序排序，宽度相同时按照高度将序排序
    sort.SliceStable(envelopes, func(i, j int) bool {
        if envelopes[i][0] == envelopes[j][0] {
            // 宽度相同时，按照高度将序排序，
            // 结合宽度递增，就严格转换成了普通的 LIS ，
            // 保证相同宽度的信封不会嵌套
            return envelopes[i][1] > envelopes[j][1]
        }
    
        // 宽度不同时，按照宽度升序排序，
        // 保证按照顺序遍历时，宽度是递增的，
        // 这样基本就转换成了普通的 LIS ，后续只需要二分高度
        return envelopes[i][0] < envelopes[j][0]
    })

    // minHeight[k] 表示嵌套的信封个数为 k 时，最外层信封的最小高度。
    // 初始嵌套信封个数为 0 时，最外层信封的最小高度为 0 ，方便后续处理。
    minHeight := []int{ 0 }
    // 遍历每个信封
    for _, envelope := range envelopes {
        // 寻找 minHeight 中第一个大于等于 envelope[1] 的下标 k ，
        // 则说明以当前信封为最外层时，最多能嵌套 k 层
        k := sort.Search(len(minHeight), func (i int) bool { return minHeight[i] >= envelope[1] });
        if len(minHeight) == k {
            // 当前信封是第一个嵌套 k 层的信封，所以直接放入 minHeight
            minHeight = append(minHeight, envelope[1])
        } else {
            // 此时存在嵌套 k 个信封的情况，
            // 因为前面二分寻找的是第一个大于等于 envelope[1] 的下标，
            // 所以 minHeight[k] >= envelope[1] ，可以直接更新为 envelope[1]
            minHeight[k] = envelope[1]
        }
    }

    // 最多嵌套 len(minHeight) - 1 层
    return len(minHeight) - 1
}
