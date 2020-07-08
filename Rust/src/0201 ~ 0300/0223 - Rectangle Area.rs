// 链接：https://leetcode.com/problems/rectangle-area/
// 题意：给定两个矩形，求它们形成的图形的总面积？

// 输入： A = -3, B = 0, C = 3, D = 4, E = 0, F = -1, G = 9, H = 2
// 输出： 45

// 思路： 枚举
//
//      首先保证第一个矩形的左侧在第二个矩形的左侧左边即可
//      先求两个矩形的面积和 area_sum ，
//		枚举判断所有情况即可：
//      1. 两个矩形不相交，直接返回面积和 area_sum
//      2. 两个矩形相交，相交矩形的面积为 area ，则总面积为 area_sum - area
//          (1) 相交矩形的左侧 = 第二个矩形的左侧
//          (2) 相交矩形的下侧 = max(第一个矩形的下侧，第二个矩形的下侧)
//          (3) 相交矩形的右侧 = min(第一个矩形的右侧，第二个矩形的右侧)
//          (4) 相交矩形的上侧 = min(第一个矩形的上侧，第二个矩形的上侧)
//
//      时间复杂度： O(height ^ 2)
//      空间复杂度： O(height) 【栈的开销】

use std::cmp;

impl Solution {
    pub fn compute_area(a: i32, b: i32, c: i32, d: i32, e: i32, f: i32, g: i32, h: i32) -> i32 {
        // 如果第一个矩形的左侧在第二个矩形的的左侧右边，则调整两个矩形的位置
        if a > e {
            return Solution::compute_area(e, f, g, h, a, b, c, d);
        }
        // 现在第一个矩形的左侧必定在第二个矩形的左侧左边
        // 算出两个矩形的面积和，等下减去相交矩形的面积即可
        let area_sum = Solution::area(a, b, c, d) + Solution::area(e, f, g, h);

        // 如果第二个矩形右侧在第一个矩形右侧右边 或者
        //    第二个矩形下侧在第一个矩形上侧上边 或者
        //    第二个矩形上侧在第一个矩形下侧下边
        // 则两个矩形不相交，直接返回面积和即可
        if e > c || f > d || h < b {
            return area_sum;
        }

        // 现在两个矩形必定相交
        // 相交矩形的左侧 = 第二个矩形的左侧
        let i = e;
        // 相交矩形的下侧 = max(第一个矩形的下侧，第二个矩形的下侧)
        let j = cmp::max(b, f);
        // 相交矩形的右侧 = min(第一个矩形的右侧，第二个矩形的右侧)
        let k = cmp::min(c, g);
        // 相交矩形的上侧 = min(第一个矩形的上侧，第二个矩形的上侧)
        let l = cmp::min(d, h);

        return area_sum - Solution::area(i, j, k, l)
    }

    pub fn area(a: i32, b: i32, c: i32, d: i32) -> i32 {
        (c - a) * (d - b)
    }
}
