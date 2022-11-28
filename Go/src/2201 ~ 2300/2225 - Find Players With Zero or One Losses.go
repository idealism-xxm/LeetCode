// 链接：https://leetcode.com/problems/find-players-with-zero-or-one-losses/
// 题意：给定一个数组 matches ，其中 matches[i] = [winner_i, loser_i] ，
//      表示玩家 winner_i 击败了 loser_i 。
//      返回一个长度为 2 的数组 answer:
//          1. answer[0] 是所有没输过任何比赛的玩家列表
//          2. answer[1] 时所有恰好输过 1 场比赛的玩家列表
//
//      注意：
//          1. answer 的两个列表应该按升序排序
//          2. 只考虑至少参加过一场比赛的玩家
//          3. matches 中不存在结果相同的两场比赛（即 matches[i] 各不相同）


// 数据限制：
//  1 <= matches.length <= 10 ^ 5
//  matches[i].length == 2
//  1 <= winneri, loseri <= 10 ^ 5
//  winner_i != loser_i
//  所有的 matches[i] 各不相同


// 输入： matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
// 输出： [[1,2,10],[4,5,7,8]]
// 解释： 玩家 1,2,10 各输了 0 场比赛；
//       玩家 4,5,7,8 各输了 1 场比赛；
//       玩家 3,6,9 各输了 2 场比赛。
//       所以 answer[0] = [1,2,10], answer[1] = [4,5,7,8]

// 输入： matches = [[2,3],[1,3],[5,4],[6,4]]
// 输出： [[1,2,5,6],[]]
// 解释： 玩家 1,2,5,6 各输了 0 场比赛；
//       玩家 3,4 各输了 2 场比赛。
//       所以 answer[0] = [1,2,5,6], answer[1] = [0]


// 思路： Map
//
//      我们用 playerToCnt 维护所有参与过比赛的玩家输过的比赛数，
//      playerToCnt[player] 表示玩家 player 输过的比赛数。
//
//      遍历 matches 的每一个比赛结果 [winner_i, loser_i] ，
//      统计每个玩家输过的比赛数。
//
//      如果 winner_i 不在 playerToCnt 中，则设置其输过的比赛数为 0 ，
//      方便后续收集为 answer[0]
//
//      对 loser_i 输过的比赛数 +1 ，即 playerToCnt[loser_i] += 1 。
//
//      然后遍历 playerToCnt ，收集输过的比赛数为 0 和 1 的玩家列表。
//
//      最后，对玩家列表按升序排序后返回。
//
//
//      设参加过比赛的玩家数为 m 。
//
//      时间复杂度：O(n + m)
//          1. 需要遍历 matches 中全部 O(n) 场比赛结果
//          2. 需要遍历 playerToCnt 中全部 O(m) 个玩家的输过的比赛数
//      空间复杂度：O(m)
//          1. 需要维护 playerToCnt 中全部 O(m) 个玩家的输过的比赛数
//          2. 需要维护 answer 满足题意的玩家列表，
//              最差情况下全部 O(m) 个玩家都满足题意


func findWinners(matches [][]int) [][]int {
    // playerToCnt[player] 表示玩家 player 输过的比赛数
    playerToCnt := make(map[int]int)
    // 遍历所有的比赛结果，统计每个玩家输过的比赛数
    for _, item := range matches {
        winner, loser := item[0], item[1]
        // 如果 winner 不在 playerToCnt 中，则设置其输过的比赛数为 0 ，
        // 方便后续收集为 answer[0]
        if _, exists := playerToCnt[winner]; !exists {
            playerToCnt[winner] = 0
        }
        
        // loser 输过的比赛数 +1
        playerToCnt[loser] += 1
    }

    // ans 收集满足题意的玩家列表
    ans := [][]int{[]int{}, []int{}}
    // 遍历所有参加过比赛的玩家 player 及其输过的比赛数 cnt
    for player, cnt := range playerToCnt {
        // 收集输过的比赛数为 0 和 1 的玩家列表
        if cnt <= 1 {
            ans[cnt] = append(ans[cnt], player)
        }
    }
    // 玩家列表按升序排序
    sort.Ints(ans[0])
    sort.Ints(ans[1])

    return ans
}
