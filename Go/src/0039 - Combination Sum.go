// 链接：https://leetcode.com/problems/combination-sum/
// 题意：有一些互不相同的正整数列表 candidates 和一个目标整数 target
//      列表中的每个数可以无限选取，从中选取若干个数形成一个列表
//      求所有和为 target 的列表？

// 输入：candidates = [2,3,6,7], target = 7
// 输出：[
//        [7],
//        [2,2,3]
//      ]

// 输入：candidates = [2,3,5], target = 8
// 输出：[
//        [2,2,2,2],
//        [2,3,3],
//        [3,5]
//      ]

// 思路：模拟即可
//      从第一个数开始进行递归选取数字，入参数为当前所选数的列表和总和
//      每次递归时，只能选下标 大于等于 前一次所选的数字的下标 的数字

import "sort"

func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)  // 按升序排序
    list := make([]int, target / candidates[0] + 1)
    return dfs(candidates, target, 0, list, 0, 0)
}

func dfs(candidates []int, target int, current int, list []int, count int, total int) [][]int {
    if total > target {  // 如果当前总和已大于目标数，则直接返回 nil
        return nil
    }
    if total == target {  // 如果当前总和等于目标数，则返回当前选择的数字
        return [][]int {append(list[:0:0], list[:count]...)}
    }

    var result [][]int
    for i := current; i < len(candidates); i++ {
        list[count] = candidates[i]  // 第 count 个数 选取 candidates[i]
        // 收集循环中产生的结果列表
        result = append(result, dfs(candidates, target, i, list, count + 1, total + candidates[i])...)
    }
    return result
}
