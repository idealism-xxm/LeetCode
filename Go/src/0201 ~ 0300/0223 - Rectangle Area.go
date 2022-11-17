// 链接：https://leetcode.com/problems/rectangle-area/
// 题意：给定两个矩形的坐标，求它们形成的图形的总面积？
//      第一个矩形的左下角为 (ax1, ay1) ，右上角为 (ax2, ay2) 。
//      第一个矩形的左下角为 (bx1, by1) ，右上角为 (bx2, by2) 。


// 数据限制：
//  -10 ^ 4 <= ax1 <= ax2 <= 10 ^ 4
//  -10 ^ 4 <= ay1 <= ay2 <= 10 ^ 4
//  -10 ^ 4 <= bx1 <= bx2 <= 10 ^ 4
//  -10 ^ 4 <= by1 <= by2 <= 10 ^ 4


// 输入： ax1 = -3, ay1 = 0, ax2 = 3, ay2 = 4, bx1 = 0, by1 = -1, bx2 = 9, by2 = 2
// 输出： 45

// 输入： ax1 = -2, ay1 = -2, ax2 = 2, ay2 = 2, bx1 = -2, by1 = -2, bx2 = 2, by2 = 2
// 输出： 16


// 思路： 模拟
//
//      题目 LeetCode 2446 的加强版，需要求出相交的长度。
//
//      对于 x 轴上的两个线段 (ax1, ax2) 和 (bx1, bx2) ，
//      当且仅当 ax1 <= bx2 && bx1 <= ax2 时，两个线段相交。
//
//      相交线段的左端点 left = max(ax1, bx1) ，右端点 right = min(ax2, bx2) ，
//      那么相交线段长度为 right - left 。
//
//      同理，对于 y 轴上的两个线段 (ay1, ay2) 和 (by1, by2) ，
//      当且仅当 ay1 <= by2 && by1 <= ay2 时，两个线段相交。
//
//      相交线段的下端点 bottom = max(ay1, by1) ，上端点 top = min(ay2, by2) ，
//      那么相交线段长度为 top - bottom 。
//
//      那么当且仅当两个矩形在 x 方向和 y 方向都相交时，它们存在重叠区域，
//      重叠区域的面积为 (right - left) * (top - bottom) 。
//
//      我们可以用做差法得计算公式：总面积 = 两个矩形的面积和 - 两个矩形重叠区域的面积。
//
//
//      时间复杂度：O(1)
//          1. 只需要常数次布尔运算和算数运算即可
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
    // x 表示矩形 x 方向相交长度。初始化为 0 ，表示不相交，方便后续处理
    x := 0
    // 如果 x 方向相交，则更新相交长度
    if ax1 <= bx2 && bx1 <= ax2 {
        x = min(ax2, bx2) - max(ax1, bx1)
    }

    // y 表示矩形 y 方向相交长度。初始化为 0 ，表示不相交，方便后续处理
    y := 0
    // 如果 y 方向相交，则更新相交长度
    if ay1 <= by2 && by1 <= ay2 {
        y = min(ay2, by2) - max(ay1, by1)
    }

    // 总面积 = 两个矩形的面积和 - 两个矩形重叠区域的面积
    return area(ax1, ay1, ax2, ay2) + area(bx1, by1, bx2, by2) - x * y
}

// 计算矩形（左下角为 (x1, y1) ，右上角为 (x2, y2) ）的面积
func area(x1, y1, x2, y2 int) int {
    return (x2 - x1) * (y2 - y1)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
