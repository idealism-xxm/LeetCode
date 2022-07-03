// 链接：https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/
// 题意：给定一个 h * w 的矩形蛋糕，同时给定两个一维数组 horizontalCuts 和 verticalCuts ，
//      其中： horizontalCuts[i] 表示在离蛋糕顶部这么距离处切一刀，
//            verticalCuts[j] 表示在离蛋糕左侧这么距离处切一刀。
//
//      求最终切出的所有蛋糕中，面积最大的蛋糕是多大？


// 数据限制：
//  2 <= h, w <= 10 ^ 9
//  1 <= horizontalCuts.length <= min(h - 1, 10 ^ 5)
//  1 <= verticalCuts.length <= min(w - 1, 10 ^ 5)
//  1 <= horizontalCuts[i] < h
//  1 <= verticalCuts[i] < w
//  horizontalCuts 中所有的数都各不相同
//  verticalCuts 中所有的数都各不相同


// 输入： h = 5, w = 4, horizontalCuts = [1,2,4], verticalCuts = [1,3]
// 输出： 4
// 解释： 面积最大的蛋糕（标记为 1 ）的高为 2 ，宽为 2 ，面积为 2 * 2 = 4
//       0|00|0
//       ------
//       0|00|0
//       ------
//       0|11|0
//       0|11|0
//       ------
//       0|00|0

// 输入： h = 5, w = 4, horizontalCuts = [3,1], verticalCuts = [1]
// 输出： 6
// 解释： 面积最大的蛋糕（标记为 1 和 2）的高为 2 ，宽为 3 ，面积为 2 * 3 = 6
//       0|000
//       -----
//       0|111
//       0|111
//       -----
//       0|222
//       0|222

// 输入： h = 5, w = 4, horizontalCuts = [3], verticalCuts = [3]
// 输出： 9
// 解释： 面积最大的蛋糕的高为 3 ，宽为 3 ，面积为 3 * 3 = 9
//       111|0
//       111|0
//       111|0
//       -----
//       000|0
//       000|0


// 思路： 贪心 + 排序
//
//      由于是直线横切和竖切，那么在同一个横切间隔中的所有蛋糕中，
//      高度是一样的，所以要想面积最大，必定是宽度最大；
//      对于在同一个竖切间隔中的所有蛋糕也是如此。
//
//      那么可以贪心地分别求横切的最大间隔 maxH 和竖切的最大间隔 maxW ，
//      那么面积最大的蛋糕大小为 maxH * maxW 。
//
//      为了计算最大间隔，我们可以先对要切的位置数组进行排序，
//      然后求所有间隔的最大值即可。
//
//
//      时间复杂度：O(nlogn + mlogm)
//          1. 需要对 horizontalCuts 进行排序，时间复杂度为 O(nlogn)
//          2. 需要遍历 horizontalCuts 中全部 O(n) 个数字
//          3. 需要对 verticalCuts 进行排序，时间复杂度为 O(mlogm)
//          4. 需要遍历 verticalCuts 中全部 O(m) 个数字
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量即可


import "sort"


func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    // 获取最大矩形的高和宽，并计算面积
    maxH, maxW := getMaxInterval(horizontalCuts, h), getMaxInterval(verticalCuts, w)
    return int((maxH * maxW) % 1000000007)
}

func getMaxInterval(vals []int, maxVal int) int64 {
    // 按升序排序，然后计算相邻值的最大间隔
    sort.Ints(vals)
    maxInterval := 0
    // 起始值为 0 ，方便计算第一个间隔
    preVal := 0
    for _, val := range vals {
        maxInterval = max(maxInterval, val - preVal)
        preVal = val
    }

    // 计算最后一个间隔值，并返回最大值
    return int64(max(maxInterval, maxVal - preVal))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
