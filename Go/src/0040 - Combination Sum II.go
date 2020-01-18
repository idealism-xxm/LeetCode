// 链接：https://leetcode.com/problems/combination-sum-ii/
// 题意：有一个正整数列表 candidates （不唯一）和一个目标整数 target
//      列表中的每个数只可以选取一次，从中选取若干个数形成一个列表
//      求所有和为 target 的列表？

// 输入：candidates = [10,1,2,7,6,1,5], target = 8
// 输出：[
//        [1, 7],
//        [1, 2, 5],
//        [2, 6],
//        [1, 1, 6]
//      ]

// 输入：candidates = [2,5,2,1,2], target = 5
// 输出：[
//        [1,2,2],
//        [5]
//      ]

// 思路：递归模拟即可（思路和前一题一致，多了校验使用次数的逻辑）
//      从第一个数开始进行递归选取数字，入参数为当前所选数的列表和总和
//      每次递归时，只能选下标 大于等于 前一次所选的数字的下标 的数字（必须还有剩余个数）

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    num := make(map[int]int)  // 记录每个数出现的次数
    var uniqueCandidates []int  // 唯一的待选数字列表
    for _, candidate := range candidates {
        if num[candidate] == 0 {
            uniqueCandidates = append(uniqueCandidates, candidate)
        }
        num[candidate]++
    }
    sort.Ints(candidates)  // 按升序排序
    list := make([]int, len(candidates))
    return dfs(uniqueCandidates, target, num, 0, list, 0, 0)
}

func dfs(uniqueCandidates []int, target int, num map[int]int, current int, list []int, count int, total int) [][]int {
    if total > target {  // 如果当前总和已大于目标数，则直接返回 nil
        return nil
    }
    if total == target {  // 如果当前总和等于目标数，则返回当前选择的数字
        return [][]int {append(list[:0:0], list[:count]...)}
    }

    var result [][]int
    for i := current; i < len(uniqueCandidates); i++ {
        if num[uniqueCandidates[i]] == 0 {  // 如果当前数已全部用完，则继续下一个数
            continue
        }

        list[count] = uniqueCandidates[i]  // 第 count 个数 选取 candidates[i]
        num[uniqueCandidates[i]]--
        // 收集循环中产生的结果列表
        result = append(result, dfs(uniqueCandidates, target, num, i, list, count + 1, total + uniqueCandidates[i])...)
        num[uniqueCandidates[i]]++
    }
    return result
}
