// 链接：https://leetcode.com/problems/shift-2d-grid/
// 题意：给定一个二维数组 grid ，将其循环右移 k 次。
//      每次循环右移按如下方式操作：
//          1. grid[i][j] 移动至 grid[i][j + 1]
//          2. grid[i][n - 1] 移动至 grid[i + 1][0]
//          3. grid[n - 1][m - 1] 移动至 grid[0][0]


// 数据限制：
//  m == grid.length
//  n == grid[i].length
//  1 <= m <= 50
//  1 <= n <= 50
//  -1000 <= grid[i][j] <= 1000
//  0 <= k <= 100


// 输入： grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
// 输出： [[9,1,2],[3,4,5],[6,7,8]]

// 输入： grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], k = 4
// 输出： [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]

// 输入： grid = [[1,2,3],[4,5,6],[7,8,9]], k = 9
// 输出： [[1,2,3],[4,5,6],[7,8,9]]


// 思路： 三次翻转
//
//      将这个二维数组 grid 看成一个一维数组 nums ，
//      长度为 m * n = length ，
//      那就转换成了 LeetCode 189 这题。
//
//		可以发现循环右移 k 次后，
//      数组末尾的 k % length 个数字会移动到数组开始，
//      而数组开始的 length - k % length 个数字，
//      则会向右移动 k % length 次。
//
//      如果我们想将数组 grid 整体翻转，
//      则可以使得末尾的 k % length 个数移动至数组开始，
//      不过此时 nums[..k % length] 和 nums[k % length..] 的顺序都是反的，
//      所以还需要分别对 nums[..k % length] 和 nums[k % length..] 再次翻转，
//      这样就能获得循环右移 k 次的结果。
//
//
//		时间复杂度： O(m * n)
//          1. 需要遍历二维数组 grid 中的全部 O(m * n) 个数字
//		空间复杂度： O(1)
//          1. 只需要使用常数个额外变量


impl Solution {
    pub fn shift_grid(mut grid: Vec<Vec<i32>>, mut k: i32) -> Vec<Vec<i32>> {
        // 计算 grid 中的数字个数
        let length = (grid.len() * grid[0].len()) as i32;
        // 计算最后有多少数字会被移动到数组开始
        k %= length;
        // 整体翻转
        Self::reverse(&mut grid, 0, length - 1);
        // 翻转前 k 个数字
        Self::reverse(&mut grid, 0, k - 1);
        // 翻转后 n - k 个数字
        Self::reverse(&mut grid, k, length - 1);

        grid
    }

    pub fn reverse(grid: &mut Vec<Vec<i32>>, mut l: i32, mut r: i32) {
        let n = grid[0].len();
        // 使用双指针翻转
        while l < r {
            let (ll, rr) = (l as usize, r as usize);
            // 交换 l 和 r 位置的数字
            let tmp = grid[ll / n][ll % n];
            grid[ll / n][ll % n] = grid[rr / n][rr % n];
            grid[rr / n][rr % n] = tmp;
            // l 向右移动一位
            l += 1;
            // r 向左移动一位
            r -= 1;
        }
    }
}
