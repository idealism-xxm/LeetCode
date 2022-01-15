// 链接：https://leetcode.com/problems/jump-game-iv/
// 题意：给定一个数组 arr ，你初始在 i = 0 处，每次在 i 时有三种跳法：
//          1. 当 i + 1 < arr.len() 时，跳到 i + 1
//          2. 当 i - 1 >= 0 时，跳到 i - 1
//          3. 存在 j 满足 i != j && arr[i] == arr[j] 时，跳到 j
//      从跳到 arr.len() - 1 时的最小跳数？

// 数据限制：
//  1 <= arr.length <= 5 * 10 ^ 4
//  -(10 ^ 8) <= arr[i] <= 10 ^ 8

// 输入： arr = [100,-23,-23,404,100,23,23,23,3,404]
// 输出： 3
// 解释： 0 --> 4 --> 3 --> 9

// 输入： arr = [7]
// 输出： 0
// 解释： 0 就是 arr 最后的下标

// 输入： arr = [7,6,9,6,9,6,9,7]
// 输出： 1
// 解释： 1 --> 7


// 思路： BFS
//
//      先对 arr 按照 arr[i] 分组，
//      再按照题意进行 bfs 即可，因为路径长度都是 1 ，找的也是最短路。
//
//      维护两个值 jumps 和 q ，
//          jumps[i] 表示跳到 i 时的最小跳数， jumps[0] = 0, jumps[1:] = i32::MAX
//          q 表示现在已经跳到过的下标队列
//
//      当队列不为空时，我们取出 q 队首的下标 cur ，
//          计算跳到下一个位置的跳数 next_jump = jumps[cur] + 1 ，
//              1. 如果 next in [cur - 1, cur + 1] 合法且未到达过，
//                  则可以跳过去，并设置 jumps[next] = next_jump ，将 next 放入到 q 中
//              2. 遍历与 arr[i] 值相同的下标 next ，如果未到达过，
//                  则可以跳过去，并设置 jumps[next] = next_jump ，将 next 放入到 q 中
//
//      （如果发现 next == arr.len() - 1 则可以提前返回，这样最后就不会到达了）
//
//
//      时间复杂度： O(n)
//      空间复杂度： O(n)


use std::collections::{HashMap, LinkedList};


impl Solution {
    pub fn min_jumps(arr: Vec<i32>) -> i32 {
        let n = arr.len();
        // 先按照 val = arr[i] 的值分组
        let mut val_to_indices: HashMap<i32, LinkedList<usize>> = HashMap::new();
        // 转成迭代器，并附上每个值的下标
        for (i, &val) in arr.iter().enumerate() {
            // 从 HashMap 中获取 val 对应的 LinkedList ，不存在则创建一个，
            // 这里为了避免每次都创建 LinkedList ，使用了 lambda 进行惰性创建
            let indices = val_to_indices.entry(val).or_insert_with(|| LinkedList::new());
            indices.push_back(i);
        }

        // 初始时每个跳到 0 时需要 0 次，
        // 跳到 [1, len) 时需要 i32:MAX ，方便后续更新
        let mut jumps = vec![i32::MAX; n];
        jumps[0] = 0;
        // 初始化 bfs 用的队列，初始只有 0 跳到了
        let mut q: LinkedList<usize> = LinkedList::new();
        q.push_back(0);

        // 当队列不为空时，继续进行循环
        while !q.is_empty() {
            // 从队首取出当前跳到的下标 cur
            let cur = q.pop_front().unwrap();
            // 如果 cur 就是最后一个位置，则直接返回
            if cur == n - 1 {
                return jumps[cur];
            }
            // 计算跳到后续的下标共需要 next_jump = jumps[cur] + 1 跳
            let next_jump = jumps[cur] + 1;
            // cur - 1 和 cur + 1 是两个可能到达的下标，转成迭代器遍历
            [cur - 1, cur + 1].iter()
                // 过滤出合法的下标
                .filter(|&&next| 0 < next && next < n)
                // 串上必定能够到达的下标（ arr[cur] 相同值的下标的 LinkedList ），
                // 并将其从 val_to_indices 移除。
                //      因为相同的 arr[i] 遍历过一次后，下次遍历不会更优，
                //      反而可能使时间复杂度升为 O(n ^ 2) ，所以这里必须移除
                .chain(val_to_indices.remove(&arr[cur]).unwrap_or_default().iter())
                // 针对每个下标 next 处理
                .for_each(|&next| {
                    // 如果这个下标 next 没有跳到过，则可以跳 next_jump 次到 next
                    if jumps[next] == i32::MAX {
                        jumps[next] = next_jump;
                        // 放入队列
                        q.push_back(next)
                    }
                });
        }

        // 由于在循环里提前返回了，所以不会到达这里
        unreachable!()
    }
}
