// 链接：https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
// 题意：有 n 个游戏角色，每个角色有自己的攻击力和防御力。
//      现在用一个数组 properties 表示，
//      properties[i] = [attack_i, defense_i] 表示第 i 个角色的攻击力和防御力。
//
//      如果一个角色的攻击力和防御力都小于其他某个角色，那么这个角色就是脆弱的，
//      求有多少个脆弱的角色？


// 数据限制：
//  2 <= properties.length <= 10 ^ 5
//  properties[i].length == 2
//  1 <= attack_i, defense_i <= 10 ^ 5


// 输入： properties = [[5,5],[6,3],[3,6]]
// 输出： 0

// 输入： properties = [[2,2],[3,3]]
// 输出： 1
// 解释： 第一个角色的攻击力和防御力都小于第二个角色的

// 输入： properties = [[1,5],[10,4],[4,3]]
// 输出： 1
// 解释： 第三个角色的攻击力和防御力都小于第二个角色的


// 思路： 排序
//
//      我们对 properties 按照攻击力倒序排序，
//      然后遍历每个角色，并维护三个值：
//          1. preAttack: 前一个角色的攻击力
//          2. preMaxDefense: 攻击力为 preAttack 的所有角色中，防御力的最大值
//          3. maxDefense: 攻击力大于当前角色攻击力的所有角色中，防御力的最大值
//
//      这样我们每次先比较当前角色的攻击力和前一个角色的攻击力，
//      如果不相等，则可以更新 maxDefense = max(maxDefense, preMaxDefense) 。
//
//      然后再判断 maxDefense 是否大于当前角色的防御力，如果大于，则对结果 +1 。
//
//      最后再更新 preAttack 和 preMaxDefense 即可。
//
//
//      时间复杂度：O(nlogn)
//          1. 需要对全部 O(n) 个角色进行排序，时间复杂度为 O(nlogn)
//          2. 需要遍历排序后的全部 O(n) 个角色
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量即可


func numberOfWeakCharacters(properties [][]int) int {
    // 按照攻击力倒序排序
    sort.SliceStable(properties, func(i, j int) bool {
        return properties[i][0] > properties[j][0]
    })

    // ans 统计脆弱的角色数
    ans := 0
    // preAttack 表示前一个角色的攻击力
    preAttack := 0
    // preMaxDefense 表示攻击力为 preAttack 的所有角色中，防御力的最大值
    preMaxDefense := 0
    // maxDefense 表示攻击力大于当前角色攻击力的所有角色中，防御力的最大值
    maxDefense := 0
    // 遍历每个角色
    for _, property := range properties {
        attack, defense := property[0], property[1]
        if attack != preAttack {
            // 如果当前角色攻击力不等于 preAttack ，
            // 则 preMaxDefense 可以去更新 maxDefense
            maxDefense = max(maxDefense, preMaxDefense)
            // 同时重置 preMaxDefense
            preMaxDefense = 0
        }
        // 如果当前角色攻击力小于 maxDefense ，则当前角色是脆弱的角色
        if defense < maxDefense {
            ans += 1
        }
        // 更新前一个角色的攻击力和其对应的最大防御力
        preAttack = attack
        preMaxDefense = max(preMaxDefense, defense)
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
