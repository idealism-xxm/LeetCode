// 链接：https://leetcode.com/problems/push-dominoes/
// 题意：有 n 个多米洛骨牌竖直朝上排成一条线，同时朝其中某些骨牌向左或向右推一下。
//      每一秒，那些被推的骨牌会向对应的方向倒下，并传递给该方向相邻的骨牌。
//      如果一个骨牌左右两边同时遇到向其倒下的骨牌，那么该骨牌由于力相互抵消而保持竖直。
//
//      现在给定一个长度为 n 的字符串 dominoes ，表示多米洛骨牌的初始状态，
//      求这些多米洛骨牌的最终状态？
//          1. dominoes[i] = 'L': 表示第 i 个骨牌被推向了左侧
//          2. dominoes[i] = 'R': 表示第 i 个骨牌被推向了右侧
//          3. dominoes[i] = '.': 表示第 i 个骨牌没有被推动


// 数据限制：
//  n == dominoes.length
//  1 <= n <= 10 ^ 5
//  dominoes[i] 是 'L', 'R' 或 '.'


// 输入： dominoes = "RR.L"
// 输出： "RR.L"
// 解释： 第 0 秒： "RR.L"
// 解释： 第 1 秒： "RR.L" ，此时所有的外力相互抵消，
//                状态保持不变，就是最终状态

// 输入： dominoes = ".L.R...LR..L.."
// 输出： "LL.RR.LLRRLL.."
// 解释： 第 0 秒： ".L.R...LR..L.."
//       第 1 秒： "LL.RR.LLRRLL.."
//       第 2 秒： "LL.RR.LLRRLL.." ，此时所有的外力相互抵消，
//                状态保持不变，就是最终状态


// 思路： DP
//
//      可以发现初始推动的骨牌状态不会改变，
//      只有原先保持竖直的骨牌会由于力的传导倒下而改变状态。
//
//      对于第 i 个骨牌来说，我们可以根据传导过来的力的时间点，得出该骨牌的最终状态
//          1. 左侧向右的力先传导过来，那么会向右倒下，并将向右的力继续向右传导下去
//          2. 右侧向左的力先传导过来，那么会向左倒下，并将向左的力继续向左传导下去
//          2. 左右侧的力同时传导过来（包括左右侧都没有力传导过来），那么力会相互抵消，保持竖直
//
//      那么我们可以使用 left 数组来维护右侧向左的力的传导情况，
//      left[i] 表示第 i 个骨牌会在第 left[i] 秒时受到向左的外力。
//      （其中 left[i] = n 表示不会受到向左的外力，方便后续统一处理）
//
//      然后我们就可以从右往左遍历骨牌的初始状态，进而递推出全部 left 的状态：
//          1. dominoes[i] == 'L': 如果初始向左推动，则第 0 秒就受到向左的外力，
//              即 left[i] = 0
//          2. dominoes[i] == 'R': 如果初始向右推动，则不会受到向左的外力，
//              即 left[i] = n
//          3. dominoes[i] == '.': 如果右侧有骨牌初始受到向左外力，
//              则其会在第 left[i + 1] + 1 秒传导过来，
//              即 left[i] = min(n, left[i + 1] + 1)
//              （与 n 取最小值保证表示没有向右的外力的情况只有一种，方便后续处理）
//
//      同理，我们可以从左往右递推出 right 数组，维护左侧向右的力的传导情况。
//      （这一步和最终步都是从左往右遍历，可以一起计算，这样就只用一个变量维护，无需使用数组维护）
//
//      最后，从左向右遍历第 i 个骨牌左右侧力的传导情况，按照前面提到的方式决定最终状态：
//          1. left[i] < right[i]: 向左的外力先到，则其向左倒下
//          2. left[i] > right[i]: 向右的外力先到，则其向右倒下
//          3. left[i] == right[i]: 左右两边的力同时到达，则相互抵消，保持竖直
//
//
//      该方法也可以使用双指针进行优化，这样就能在保持时间复杂度为 O(n) 的情况下，
//      将空间复杂度优化为 O(1) 。
//      本实现为了便于理解，不做优化处理。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 dominoes 中全部 O(n) 个字符
//          2. 需要遍历 left 中全部 O(n) 个状态
//      空间复杂度：O(n)
//          1. 需要维护 left 中全部 O(n) 个状态
//          2. 需要维护最终状态的结果字符串中全部 O(n) 的字符


func pushDominoes(dominoes string) string {
    n := len(dominoes)
    // left[i] 表示第 i 个骨牌会在第 left[i] 秒时受到向左的外力。
    // n 表示不会受到向左的外力，方便后续统一处理。
    // 初始化长度为 n + 1 ，方便处理最右侧为 '.' 的边界情况。
    left := make([]int, n + 1)
    left[n] = n
    for i := n - 1; i >= 0; i-- {
        switch dominoes[i] {
        // 如果初始向左推动，则第 0 秒就受到向左的外力
        case 'L': left[i] = 0
        // 如果初始向右推动，则不会受到向左的外力
        case 'R': left[i] = n
        // 如果右侧有骨牌初始为受到向左外力，则其会在第 left[i + 1] + 1 秒传导过来。
        // 与 n 取最小值保证表示没有向右的外力的情况只有一种，方便后续处理。
        case '.': left[i] = min(n, left[i + 1] + 1)
        // 题目数据保证不存在其他情况
        }
    }

    // ans[i] 表示第 i 个骨牌的最终状态，初始化都为竖直状态
    ans := make([]byte, n)
    // right 表示第 i 个骨牌会在第 right 秒时受到向右的外力。
    // 初始化为 n ，表示第 -1 个骨牌不会受到向右的外力，
    // 方便处理最左侧为 '.' 的边界情况。
    right := n
    for i := 0; i < n; i++ {
        switch dominoes[i] {
        // 如果初始向右推动，则第 0 秒就受到向右的外力
        case 'R': right = 0
        // 如果初始向左推动，则不会受到向右的外力
        // （注意，此时要将 right 重置为 n ，抵消前面向右传导的外力）
        case 'L': right = n
        // 如果左侧有骨牌初始为受到向右外力，则其会在第 right + 1 秒传导过来。
        // 与 n 取最小值保证表示没有向右的外力的情况只有一种，方便后续处理。
        case '.': right = min(n, right + 1)
        // 题目数据保证不存在其他情况
        }

        if left[i] < right {
            // 如果向左的外力先到，则其向左倒下
            ans[i] = 'L'
        } else if left[i] > right {
            // 如果向右的外力先到，则其向右倒下
            ans[i] = 'R'
        } else {
            // left[i] == right 时，左右两边的力同时到达，则相互抵消，保持竖直
            ans[i] = '.'
        }
        
    }

    // 转成字符串返回
    return string(ans)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
