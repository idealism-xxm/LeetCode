// 链接：https://leetcode.com/problems/number-of-pairs-of-interchangeable-rectangles/
// 题意：给定一组矩形 rectangles ， rectangles[i] = [width_i, height_i] 表示第 i 个矩形的宽和高。
//      如果两个矩形 i 和 j (i < j) 满足 width_i / height_i == width_j / height_j ，
//      那么它们是可互换的，求这些矩形中有多少对可互换的？

// 数据限制：
//  n == rectangles.length
//  1 <= n <= 10 ^ 5
//  rectangles[i].length == 2
//  1 <= width_i, height_i <= 10 ^ 5


// 输入： rectangles = [[4,8],[3,6],[10,20],[15,30]]
// 输出： 6
// 解释： 矩形 0 和 1: 4/8 == 3/6
//       矩形 0 和 2: 4/8 == 10/20
//       矩形 0 和 3: 4/8 == 15/30
//       矩形 1 和 2: 3/6 == 10/20
//       矩形 1 和 3: 3/6 == 15/30
//       矩形 2 和 3: 10/20 == 15/30

// 输入： rectangles = [[4,5],[7,8]]
// 输出： 0


// 思路： Map
//
//       先求出每个矩形的宽高比的最简分数，分子分母都除以最大公约数即可。
//
//       然后统计每个最简分数的矩形数。
//
//       最后就是计算组合，即从每个最简分数中的矩形都可以两两互换。
//
//
//       设宽和高的最大值为 M 。
//
//       时间复杂度： O(nM)
//          1. 需要遍历全部 O(n) 个矩形，
//             每次都要求一次最大公约数，辗转相除法的时间复杂度为 O(logM)
//          2. 需要遍历全部 O(n) 个不同的最简分数
//       空间复杂度： O(n)
//          1. 需要维护全部 O(n) 个不同的最简分数的个数


use std::collections::HashMap;
use std::mem::swap;
use std::ops::AddAssign;


impl Solution {
    pub fn interchangeable_rectangles(rectangles: Vec<Vec<i32>>) -> i64 {
        // fraction_to_cnt 维护每个最简分数的个数
        let mut fraction_to_cnt = HashMap::new();
        for rectangle in rectangles {
            let (w, h) = (rectangle[0], rectangle[1]);
            // 求最大公约数
            let g = Self::gcd(w, h);
            // 统计最简分数的个数
            fraction_to_cnt.entry((w / g, h / g)).or_insert(0).add_assign(1);
        }
        
        // 计算所有最简分数的组合
        fraction_to_cnt
            .values()
            // 每个最简分数中的矩形都可以两两互换
            .map(|&value| value * (value - 1) / 2)
            .sum()
    }

    // 辗转相除法计算最大公约数
    fn gcd(mut a: i32, mut b: i32) -> i32 {
        while b != 0 {
            // (a, b) = (b, a % b);
            // 由于 Rust 上述写法暂不可用，转换成以下写法
            a = a % b;
            swap(&mut a, &mut b)
        }

        a
    }
}
