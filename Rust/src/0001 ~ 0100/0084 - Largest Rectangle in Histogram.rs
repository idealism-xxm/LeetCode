// 链接：https://leetcode.com/problems/largest-rectangle-in-histogram/
// 题意：给定一个非负数整型数组，表示柱状图每处的高度，求能形成的最大矩形的面积？

// 数据限制：
//  1 <= heights.length <= 10 ^ 5
//  0 <= heights[i] <= 10 ^ 4

// 输入：heights = [2,1,5,6,2,3]
// 输出：10
// 解释：最大矩形面积是 10 = 2 * 5

// 输入：heights = [2,4]
// 输出：4
// 解释：有两种可能的最大矩形，他们的面积都是 4 = 2 * 2 = 1 * 4

// 思路1：单调栈
//
//      单调栈就是栈内元素保持严格单调性的栈，即栈内元素严格递增或严格递减。
//
//      单调递增栈可以使用如下方法维护（假设新元素为 num ，栈为 stack ，栈顶为 top ）：
//          1. num > stack[top]: 直接将 num 压入栈中；
//          2. num <= stack[top]: 不断弹出栈顶元素，直至 num > stack[top] ，然后再入栈
//              在这个过程中有以下两个特性：
//              (1) num 是 stack[top] 右侧第一个小于 stack[top] 的元素
//              (2) stack[top - 1] 是 stack[top] 左侧第一个小于 stack[top] 的元素
//
//      所以本题可以维护一个单调递增栈 stack ，存储单调递增的高度的下标，
//      根据这两个特性来计算最大矩形面积 max_area 。
//
//      遍历高度数组 heights ，假设当前遍历到第 i 个高度 heights[i] ，
//		    1. top < 0 || heights[stack[top]] < heights[i]: 
//              则当前能够使栈维持单调递增，直接入栈即可
//		    2. top >= 0 && heights[stack[top]] >= heights[i]:
//              对于当前栈顶下标的高度 heights[stack[top]] 有：
//                  (1) heights[i] 是 heights[stack[top]] 右侧第一个小于它的高度
//                  (2) heights[stack[top - 1]] 是 heights[stack[top]] 左侧第一个小于它的高度
//              则以 heights[stack[top]] 为矩形的高时，该矩形的右边界为 i - 1 ，
//              该矩形的左边界为 stack[top - 1] + 1 （若 top == 0 ，则左边界为 0 ）。
//
//              则该矩形的面积为 area = heights[stack[top]] * (i - stack[top - 1] - 1)  ，
//              更新矩形面积的最大值为 max_area = max(max_area, area)
//
//              不断执行这一过程，并弹出栈顶元素，直至不满足 2 的条件
//
//		为了方便操作，在 heights 后加上 -1 ，最后让所有元素出栈
//
//		时间复杂度： O(n) 
//      空间复杂度： O(n)

use std::iter;

impl Solution {
    pub fn largest_rectangle_area(heights: Vec<i32>) -> i32 {
        heights
            // 转成迭代器
            .iter()
            // 串上 -1 ，方便最后将所有数字出栈，保证考虑所有可能的情况
            .chain(iter::once(&-1))
            // 加上每个高度的下标，方便后续计算矩形宽度
            .enumerate()
            // 使用 fold 计算最新的最大矩形面积 max_area 和单调栈 stack
            .fold(
                // 初始不存在矩形，所以 max_area 为 0
                // 初始单调栈为空，所以 stack 为空
                (0, Vec::new()), 
                // 对于当前 max_area 和 stack ，根据当前的高度 height 和其下标 index ，
                // 计算出最新的 max_area 和 stack
                |(mut max_area, mut stack), (i, &height)| {
                    // 当栈不为空 且 栈顶下标的高度 > 当前高度 height 时，
                    // 则需要将栈顶元素出栈，并计算和更新最大矩形面积
                    while !stack.is_empty() && heights[*stack.last().unwrap()] >= height {
                        // 弹出栈顶元素，并获取矩形高度的下标
                        let j = stack.pop().unwrap();
                        // 如果栈为空，则矩形的左边界为 0 ，否则为现在栈顶元素 + 1
                        let left = if stack.is_empty() { 0 } else { *stack.last().unwrap() + 1 };
                        // 矩形右边界是 i - 1 ，则宽度是 i - 1 - left + 1 = i - left
                        let width = (i - left) as i32;
                        // 这一段矩形 [j, i - 1] 的高度是 heights[j] ，
                        // 计算面积 width * heights[j] 并进行更新
                        max_area = max_area.max(width * heights[j]);
                    }
                    // 将当前高度 height 的下标 i 入栈，维持严格单调递增
                    stack.push(i);
                    // 返回此时的最大矩形面积和单调栈
                    (max_area, stack)
                }
            )
            // 最后直接返回最大矩形面积即可
            .0
    }
}
