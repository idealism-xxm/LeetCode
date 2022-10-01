// 链接：https://leetcode.com/problems/the-skyline-problem/
// 题意：给定第一象限的一些矩形（底边在 x 轴上），按左边升序排序，
//      用 [l, r, h] 表示左边 x 值，右边 x 值和 y 值，
//      求这些矩形形成的轮廓线横线上左侧的点列表？


// 数据限制：
//  1 <= buildings.length <= 10 ^ 4
//  0 <= left_i < right_i <= 2 ^ 31 - 1
//  1 <= height_i <= 2 ^ 31 - 1
//  buildings 按照 left_i 升序排序


// 输入： buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
// 输出： [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]

// 输入： buildings = [[0,2,3],[2,5,3]]
// 输出： [[0,3],[5,0]]


// 思路1： 分治
//
//      最开始自己还是想直接去解决：
//          先离散化，然后用线段树维护每一个点的最大值，最后再遍历一遍即可，
//          时间复杂度： O(nlogn) 空间复杂度： O(nlogn)
//      但觉得这样没必要，而且比较难写，就直接看题解了
//
//      发现可以用分治解决就有点恍然大悟：我们每次可以将需要求的矩形分为两半分治处理
//          1. 如果当前只有一个矩形，则返回两个坐标的列表
//              第一个是矩形左上角的坐标，第二个是矩形右下角的坐标
//          2. 如果不止一个矩形，则分成两半分别求左右部分的列表，然后进行合并，
//              合并的时候类似于归并排序，维护以下值
//                (1) 两个指针 l 和 r ，分别表示左右部分当前考虑的元素
//                (2) 两个高度 left_y 和 right_y ，
//                  分别表示左右部分前一个考虑的元素的 y 值
//                (3) 一个高度 result_y ，表示前一个结果的 y 值
//              如果左右部分都还有待考虑的元素，则比较它们 x 值，
//                (1) 左边部分 x 较小，则准备使用其放入结果中，
//                  更新 left_y 为对应的 y 值，同时右移 l
//                (2) 右边部分 x 较小，则准备使用其放入结果中，
//                  更新 right_y 为对应的 y 值，同时右移 r
//              然后令 max_y = max(left_y, right_y) ，表示当前 x 值对应的 y 值
//              若 max_y != result_y ，则存在转折点，需要放入结果中
//              最后将所有剩余有转折点的部分全部放入结果中
//
//      时间复杂度： O(nlogn)
//      空间复杂度： O(n)

impl Solution {
    pub fn get_skyline(buildings: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        // 如果没有矩阵，直接返回空列表
        if buildings.len() == 0 {
            return Vec::new();
        }

        // 分治处理
        Solution::dfs(&buildings)
    }

    fn dfs(buildings: &[Vec<i32>]) -> Vec<Vec<i32>> {
        // 如果当前只有一个矩形，则返回两个坐标的列表
        // 第一个是矩形左上角的坐标，第二个是矩形右下角的坐标
        if buildings.len() == 1 {
            return [
                [buildings[0][0], buildings[0][2]].to_vec(),
                [buildings[0][1], 0].to_vec(),
            ].to_vec();
        }

        // 如果不止一个矩形，递归分治处理
        let mid = buildings.len() >> 1;
        let mut left_result = Solution::dfs(&buildings[..mid]);
        let mut right_result = Solution::dfs(&buildings[mid..]);
        // 合并左右部分结果
        let mut result: Vec<Vec<i32>> = Vec::new();
        // 左右部分的指针
        let mut l = 0;
        let mut r = 0;
        // 左右部分当前高度
        let mut left_y = 0;
        let mut right_y = 0;
        // 结果当前高度
        let mut result_y = 0;
        // 当两者都还有元素时，循环处理
        while l < left_result.len() && r < right_result.len() {
            // 当前 result_x 取两者 x 的较小值
            let mut result_x = 0;
            // 【注意】相等时必须取右侧的，否则可能造成结果错误 [[0,2,3],[2,5,3]]
            // 因为 x 相等时，共有以三种合法情况：
            //  1. 左侧是左上角，右侧是左上角： y 值相同，先后不影响结果
            //  2. 左侧是左上角，右侧是右下角：不存在这种情况，
            //      因为左侧的左边 <= 右侧的左边，且矩形宽度不为 0
            //  3. 左侧是右下角，右侧是左上角：若先处理左侧，左侧 y 值为 0 ，
            //      两个相接的矩形高度相同时，错误地把这个点也算入了结果中
            //  4. 左侧是右下角，右侧是右下角： y 值相同，先后不影响结果
            if left_result[l][0] < right_result[r][0] {
                // 如果左边的 x 小于 右边的 x ，则本次 result 使用左边的
                result_x = left_result[l][0];
                // 并且更新 left_y
                left_y = left_result[l][1];
                // 左边指针右移
                l += 1;
            } else {
                // 如果左边的 x 大于等于 右边的 x ，则本次 result 使用右边的
                result_x = right_result[r][0];
                // 并且更新 right_y
                right_y = right_result[r][1];
                // 右边指针右移
                r += 1;
            }

            // 当前 result_x 对应的高度为左边和右边在当前点高度的较大值
            let max_y = i32::max(left_y, right_y);
            // 如果高度和上次放入的高度不同，则存在转折点，
            // 需要更新 result_y ，并将当前转折放入 result
            if max_y != result_y {
                result_y = max_y;
                Solution::update_result(&mut result, result_x, result_y);
            }
        }
        // 如果是左边部分全部用完，则将右边部分赋值过去，方便后续处理
        if l >= left_result.len() {
            l = r;
            left_result = right_result;
            left_y = right_y;
        }
        // 遍历左边剩余的每个元素
        for i in l..left_result.len() {
            // 如果当前元素出现了转折点，则更新 result_y ，然后放入结果中
            if left_result[i][1] != result_y {
                result_y = left_result[i][1];
                Solution::update_result(&mut result, left_result[i][0], result_y);
            }
        }

        result
    }

    fn update_result(result: &mut Vec<Vec<i32>>, x: i32, y: i32) {
        if !result.is_empty() && result.last_mut().unwrap()[0] == x {
            // 如果 result 不为空，且 last.x == x ，
            // 则需要更新对应的 y
            result.last_mut().unwrap()[1] = y;
        } else {
            // 如果 result 为空，或者 最后一个元素的 last.x != x，
            // 则需要插入 [x, y].to_vec()
            result.push([x, y].to_vec());
        }
    }
}


// 思路2： 排序 + 优先队列/堆
//
//      可以发现转折点只会在以下两种情况下产生：
//          1. 当前 x 处恰好好进入一个更高的矩形的左侧，那么转折点就是该矩形的左上点坐标
//          2. 当前 x 处恰好离开当前最高矩形的右侧，那么转折点的横坐标就是 x ，
//              纵坐标为此后最高矩形的高度
//
//      所以我们可以将一个矩形的坐标信息 (l, r, h) 拆成左侧 (l, -h) 和右侧 (r, h) ，
//      以 (x, height) 存储在 heights 数组中。
//
//      这里左侧的高度是用负数表示，一是要和矩形右侧区分开来，二是方便后续处理 x 相同时的情况。
//
//      然后我们就可以对 heights 按照 x 升序排序，再按 height 升序排序。
//      这样后续在遍历时， x 相同时，高的矩形先进入后离开，就无需按 x 分组遍历。
//
//      然后用一个最大堆 cur_heights 维护当前 x 处所有矩形的高度。
//      并用名为 height_count 的 map 维护当前 x 处所有矩形的不同高度的出现次数，
//      标记哪些高度是无效的，方便后续移除。
//
//      此时遍历 heights 中的每个元素 (x, height) ，根据 height 的正负进行处理：
//          1. height < 0: 矩形左侧，将当前高度放入 cur_heights ，并增加出现次数
//          2. height > 0: 矩形右侧，减小出现次数即可。
//              （优先队列/堆 无法删除指定元素，所以等实际取的时候再从 cur_heights 中删除）
//
//      当前 x 处已经考虑所有的进出情况后，移除无效的最大高度，找到有效的最大高度 max_height 。
//      若 max_height 不等于前一处转折点的高度 pre_height ，则出现了转折点，
//      将 (x, max_height) 放入结果列表中，再更新 pre_height 为 max_height 即可。
//
//
//      时间复杂度： O(nlogn)
//          1. 需要遍历 buildings 全部 O(n) 个元素
//          2. 需要对 heights 全部 O(n) 个元素进行排序，时间复杂度为 O(nlogn)
//          3. 需要将 heights 全部 O(n) 个元素放入堆一次，再从堆中取出一次，
//              每次时间复杂度为 O(logn)
//      空间复杂度： O(n)
//          1. 需要维护 heights, cur_heights, height_count 中全部 O(n) 个元素
//          2. 需要维护结果数组 ans 中全部 O(n) 个元素


use std::collections::{ BinaryHeap, HashMap };
use std::ops::{ AddAssign, SubAssign };


impl Solution {
    pub fn get_skyline(buildings: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        // heights[i] = (x_i, height_i) 表示 x_i 处的矩形高度的变化信息
        //      1. height_i < 0: 表示此处刚进入一个高度为 -height_i 的矩形左侧
        //      2. height_i > 0: 表示此处刚离开一个高度为  height_i 的矩形右侧
        let mut heights = Vec::with_capacity(buildings.len() << 1);
        // 遍历每个矩形，将高度变化信息放入 height 中
        for building in buildings.iter() {
            heights.push((building[0], -building[2]));
            heights.push((building[1], building[2]));
        }
        // 先按 x 升序排序，再按 height 升序排序。
        // 这样后续在遍历时， x 相同时，高的矩形先进入后离开，
        // 就无需按 x 分组遍历
        heights.sort();

        // ans 收集所有转折点的坐标
        let mut ans = vec![];
        // cur_heights 维护当前 x 处所有矩形的高度
        let mut cur_heights = BinaryHeap::new();
        // height_count 维护当前 x 处所有矩形的不同高度的出现次数
        let mut height_count = HashMap::new();
        // 初始存在地面高度 0 ，方便后面处理不存在任何矩形的情况
        cur_heights.push(0);
        height_count.insert(0, 1);
        // pre_height 维护上次转折点的高度，初始化为地面高度 0
        let mut pre_height = 0;
        // 遍历每一个 x 及对应的高度列表
        for &(x, height) in heights.iter() {
            if height < 0 {
                // 矩形左侧，将当前高度放入
                height_count.entry(-height).or_insert(0).add_assign(1);
                cur_heights.push(-height)
            } else {
                // 矩形右侧，将当前高度移除
                height_count.entry(height).or_insert(0).sub_assign(1);
                // 优先队列/堆 无法删除指定元素，所以等实际取的时候再从 cur_heights 中删除
            }

            // 当前 x 处已经考虑所有的进出情况后，移除无效的最大高度，找到有效的最大高度
            while height_count[cur_heights.peek().unwrap()] == 0 {
                cur_heights.pop();
            }
            let max_height = *cur_heights.peek().unwrap();
            // 如果当前最大高度不等于前一个最大高度，则出现了转折点，放入结果列表中
            if max_height != pre_height {
                // 当前转折点放入结果列表
                ans.push(vec![x, max_height]);
                // 更新上次转折点高度
                pre_height = max_height;
            }
        }

        ans
    }
}
