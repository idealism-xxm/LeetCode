// 链接：https://leetcode.com/problems/container-with-most-water/
// 题意：给定一个长度为 n 的非负整数数组 height ，
//      其中 height[i] 表示在端点为 (i, 0) 和 (i, height[i]) 的线段，
//      求任意两条线段组成的矩形的最大面积？


// 数据限制：
//  n == height.length
//  2 <= n <= 10 ^ 5
//  0 <= height[i] <= 10 ^ 4


// 输入： height = [1,8,6,2,5,4,8,3,7]
// 输出： 49

// 输入： height = [1,1]
// 输出： 1


// 思路： 双指针
//
//      维护题目所需的最大面积 ans ，初始化为 0 。
//
//      我们定义左右指针 l 和 r ，分别初始化为 0 和 n - 1 。
//
//      每次计算当前左右指针对应的线段组成的矩形的面积，
//      即 area = min(height[l], height[r]) * (r - l) 。
//
//      然后更新 ans 的最大值，即 ans = max(ans, area) 。
//
//      接下来按照以下规则更新左右指针，并重复上述过程，直至 l >= r ：
//          1. height[l] <= height[r]: l 向右移动一位，r 不动。
//			   假设 r 向左移动一位，则面积必定变小，
//              因为： (r - 1 - l) * min(height[l], height[r - 1])
//			       <= (r - 1 - l) * height[l]
//			       <  (r - l) * height[i] = area
//          2. height[l] > height[r]: r 向左移动一位，l 不动。
//			   假设 l 向右移动一位，则面积必定变小，
//              因为： (r - (l + 1)) * min(height[l + 1], height[r])
//			       <= (r - (l + 1)) * height[r]
//			       <  (r - l) * height[r] = area
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 height 中全部 O(n) 个数字
//      空间复杂度：O(1)
//          1. 只需要维护常数个额外变量


impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        // 维护题目所需的最大面积 ans ，初始化为 0
        let mut ans = 0;
        // 左指针 l ，初始化为 0
        let mut l = 0;
        // 右指针 r ，初始化为 height.len() - 1
        let mut r = height.len() - 1;
        // 当矩形底还存在时，继续循环
        while l < r {
            // 计算当前矩形的面积
            let area = (r - l) as i32 * (height[l].min(height[r]));
            // 更新 ans 的最大值
            ans = ans.max(area);
            if height[l] <= height[r] {
                // 如果 height[l] <= height[r] ，
                // 则 l 向右移动一位，r 不动
                l += 1;
            } else {
                // 如果 height[l] > height[r] ，
                // 则 r 向左移动一位，l 不动
                r -= 1;
            }
        }

        ans
    }
}
