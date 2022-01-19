// 链接：https://leetcode.com/problems/can-place-flowers/
// 题意：给定一个长度为 m 的整型数组 flowerbed ，表示一个花坛，
//          flowerbed[i] = 0 表示这个位置没有种花，
//          flowerbed[i] = 1 表示这个位置种了一朵花。
//      现要求花坛中的任意两朵花不能相邻，求能否在花坛中再种 n 朵花？

// 数据限制：
//  1 <= flowerbed.length <= 2 * 10 ^ 4
//  flowerbed[i] 是 0 或 1
//  flowerbed 没有相邻的两朵花
//  0 <= n <= flowerbed.length

// 输入： flowerbed = [1,0,0,0,1], n = 1
// 输出： true
// 解释： 这朵花只能种在 i = 2 处

// 输入： flowerbed = [1,0,0,0,1], n = 2
// 输出： false
// 解释： 这个花坛中最多只能再种一朵花，在 i = 2 处


// 思路：贪心
//
//      我们可以贪心地计算做多能种的花的数量 cnt ，
//          如果 cnt >= n ，则说明能种 n 朵花，返回 true ；
//          如果 cnt < n ，则说明不能种 n 朵花，返回 false 。
//
//      贪心的方式就是每次发现能种花时，就一定要种花。
//      我们在遍历 flowerbed 时维护两个值： cnt 和 pre_status 
//          cnt 表示当前最多能种的花的数量
//          pre_status 表示前一个位置的种花状态，有三种状态：
//              EMPTY: 表示没有种花
//              PRE_PLANTED: 表示预种花，即前一个位置、再前一个位置都没有种花，
//                  基于贪心的想法，前一个位置可以预种花，能否真种花要取决于下一个位置是否为空
//              PLANTED: 表示种了一朵花
//
//      初始化 cnt = 0, pre_status = EMPTY ，
//      然后遍历 flowerbed 的值，设 cur = flowerbed[i] ，
//      则根据 cur 和 pre_status 的值进行如下处理：
//          1. cur == 1: 当前位置已种有一朵花，则当前位置不能再种花，即 cnt 保持不变
//              对于下一个位置来说， pre_status 需要设置为 PLANTED
//          2. cur == 0 && pre_status == EMPTY: 当前位置和前一个位置都没有种花，
//              则当前位置可以预种花（假装要种花，因为后面如果已经种了花的话，就不能假戏真做），
//              此时保持 cnt 不变，将 pre_status 设置为 PRE_PLANTED
//          3. cur == 0 && pre_status == PRE_PLANTED: 当前位置没有种花，
//              则前一个位置的预种花可以假戏真做，
//              此时另 cnt += 1 ，将 pre_status 需要设置为 PLANTED
//          4. cur == 0 && pre_status == PLANTED: 前一个位置已种一朵花，
//              则当前位置不能再种花，即 cnt 保持不变，将 pre_status 设置为 EMPTY
//
//      注意最后一个位置的状态如果是 PRE_PLANTED 的话，则也可以再种一朵花，
//          即另 cnt += 1
//      （如果在 flowerbed 后面加一个 0 ，则无需此操作，能使代码更加简洁）
//
//      最后返回 cnt >= n 即可
//      
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

use std::iter;

// 花坛某一位置的状态
enum Status {
    // 未种花
    Empty,
    // 预种花
    PrePlanted,
    // 已种花
    Planted,
}

impl Solution {
    pub fn can_place_flowers(flowerbed: Vec<i32>, n: i32) -> bool {
        // 迭代器遍历
        flowerbed.iter()
            // 在 flowerbed 末尾串上一个 0
            .chain(iter::once(&0))
            // 积累 cnt 和 pre_status ，
            //      cnt 表示当前最多能种的花的数量
            //      pre_status 表示前一个位置的种花状态
            .fold((0, Status::Empty), |(cnt, pre_status), &cur| {
                // 根据 cur, pre_status 处理返回对应的新 cnt 和 pre_status
                match (cur, pre_status) {
                    // 如果当前位置已种有一朵花，则当前位置不能再种花，即 cnt 保持不变，
                    // 对于下一个位置来说， pre_status 需要设置为 PLANTED
                    (1, _) => (cnt, Status::Planted),
                    // 当前位置和前一个位置都没有种花，
                    // 则当前位置可以预种花（假装要种花，因为后面如果已经种了花的话，就不能假戏真做），
                    // 此时保持 cnt 不变， pre_status 需要设置为 PRE_PLANTED
                    (0, Status::Empty) => (cnt, Status::PrePlanted),
                    // 当前位置没有种花，则前一个位置的预种花可以假戏真做，
                    // 此时 cnt 需要设置为 cnt + 1 ，将 pre_status 需要设置为 PLANTED
                    (0, Status::PrePlanted) => (cnt + 1, Status::Empty),
                    // 前一个位置已种一朵花，则当前位置不能再种花，
                    // 即 cnt 保持不变， pre_status 需要设置为 EMPTY
                    (0, Status::Planted) => (cnt, Status::Empty),
                    // 不存在其他情况
                    _ => unreachable!(),
                }
            })
            // 最后只要最多能种的花大于等于 n ，满足题意
            .0 >= n
    }
}
