// 链接：https://leetcode.com/problems/the-skyline-problem/
// 题意：给定第一象限的一些矩形（底边在 x 轴上），按左边升序排序，
//      用 [l, r, h] 表示左边 x 值，右边 x 值和 y 值，
//      求这些矩形形成的轮廓线横线上左侧的点列表？

// 输入： [ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ]
// 输出： [ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ]

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

// 思路2： 扫描线
//
//      最开始自己还是想直接去解决：
//          先离散化，然后用线段树维护每一个点的最大值，最后再遍历一遍即可，
//          时间复杂度： O(nlogn) 空间复杂度： O(nlogn)
//
//      写完分治的解法又在看其他的解法，发现扫描线这个解法其实与我的思路差不多，
//      只不过没有用线段树先维护每一点的最大值，而是动态计算每一点的最大值，
//      我们维护每一点的高度变化值列表：正数表示一个矩形刚开始，负数表示一个矩形刚结束，
//      然后需要考虑的就是顺序遍历 x 值，而 rust 中有 BTreeMap 可以按照 key 的顺序遍历，
//      这样我们就可以动态维护每一个 x 值对应的所有高度，如果还用 BTreeMap 维护，
//      那么就可以在 O(logn) 中取得高度最大值
//      找到高度最大值后，和上一次的高度最大值比较，如果不同就是转折点，放入结果列表中即可
//
//      时间复杂度： O(nlogn)
//      空间复杂度： O(n)

use std::collections::BTreeMap;

impl Solution {
    pub fn get_skyline(buildings: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        // 高度列表 map ，key 是 x 值，value 是这一点开始/结束的矩形的高度/高度相反数
        let mut heights_map: BTreeMap<i32, Vec<i32>> = BTreeMap::new();
        // 遍历每个矩形
        for building in buildings.iter() {
            // 左边 x 对应的列表放入高度
            heights_map.entry(building[0]).or_insert_with(Vec::new).push(building[2]);
            // 右边 x 对应的列表放入高度相反数
            heights_map.entry(building[1]).or_insert_with(Vec::new).push(-building[2]);
        }

        // 最终结果
        let mut result: Vec<Vec<i32>> = Vec::new();
        // 维护当前点的所有存在的高度相反数及其次数
        //（目前 last_key_valu 等 api 是 nightly-only 的，所以这里存储高度相反数）
        let mut height_count: BTreeMap<i32, i32> = BTreeMap::new();
        // 放入高度 0 一次，方便后面所有高度都为空时取值处理
        height_count.insert(0, 1);
        // 初始化上次转折点高度为 0
        let mut last_height = 0;
        // 遍历每一个 x 及对应的高度列表
        for (x, heights) in heights_map.iter() {
            // 遍历每一个高度
            for height in heights.iter() {
                if *height > 0 {
                    // 矩形左边，将当前高度放入
                    *height_count.entry(-height).or_insert(0) += 1;
                } else {
                    // 矩形右边，将对应高度移除
                    let mut count = height_count.get_mut(height).unwrap();
                    *count -= 1;
                    // 如果移除以后当前高度全部没了，就从 map 中移除
                    if *count == 0 {
                        height_count.remove(height);
                    }
                }
            }
            // 当前点已经考虑所有的进出情况后，找到最大高度
            let (height, _) = height_count.iter().next().unwrap();
            let height = -height;
            // 如果当前最大高度不等于前一个最大高度，则出现了转折点，放入结果列表中
            if height != last_height {
                // 当前转折点放入结果列表
                result.push([*x, height].to_vec());
                // 更新上次转折点高度
                last_height = height;
            }
        }

        result
    }
}
