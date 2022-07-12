// 链接：https://leetcode.com/problems/queue-reconstruction-by-height/
// 题意：给定一个数组 people ，其中 people[i] = [h_i, k_i] ，
//      h_i 表示第 i 个人的高度， k_i 表示在第 i 个人之前身高大于等于 h_i 的人数。
//
//      根据 people 构建满足上述身高限制的队列 queue ，
//      其中 queue[j] = [h_j, k_j] ， queue[0] 表示队首的人。


// 数据限制：
//  1 <= people.length <= 2000
//  0 <= h ^ i <= 10 ^ 6
//  0 <= k ^ i < people.length
//  题目确保 queue 一定能够构建成功


// 输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
// 输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
// 解释：第 0 个人的高度为 5 ，前面没有更高或相等的人。
//      第 1 个人的高度为 7 ，前面没有更高或相等的人。
//      第 2 个人的高度为 5 ，前面有 2 个更高或相等的人，这 2 个人的下标为 0 和 1
//      第 3 个人的高度为 6 ，前面有 1 个更高或相等的人，这个人的下标为 1
//      第 4 个人的高度为 4 ，前面有 4 个更高或相等的人，这 4 个人的下标为 0, 1, 2, 3
//      第 5 个人的高度为 7 ，前面有 1 个更高或相等的人，这个人的下标为 1
//
//      因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是一个满足题意的队列

// 输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
// 输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
// 解释：第 0 个人的高度为 4 ，前面没有更高或相等的人。
//      第 1 个人的高度为 5 ，前面没有更高或相等的人。
//      第 2 个人的高度为 2 ，前面有 2 个更高或相等的人，这 2 个人的下标为 0 和 1
//      第 3 个人的高度为 3 ，前面有 2 个更高或相等的人，这 2 个人的下标为 0 和 1
//      第 4 个人的高度为 1 ，前面有 4 个更高或相等的人，这 4 个人的下标为 0, 1, 2, 3
//      第 5 个人的高度为 6 ，前面没有更高或相等的人。
//
//      因此 [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]] 是一个满足题意的队列


// 思路：贪心 + 排序
//
//      我们可以发现对于高度相同的两个人 i 和 j 来说，必定有 k_i != k_j ，
//      因为两个人无法并排站。
//
//      假设 k_i == k_j ，不妨让 i 站在 j 前面，则在 i 之前高度大于等于 h_i 的人数为 k_i 。
//
//      那么在 j 之前高度大于等于 h_j = h_i 的人数为 k_j >= k_i + 1 > k_i ，
//      这与我们的假设矛盾，所以必定有 k_i != k_j 。
//
//      我们可以先贪心地排列高度最高的那群人，按照 k_i 值升序依次插入到 queue[k_i] 中，
//      因为 queue 中所有人的高度都大于等于 h_i 。
//
//      然后再贪心地排列高度次高的人，按照 k_j 值升序依次插入到 queue[k_j] 中，
//      因为 queue 中所有人的高度都大于等于 h_j 。
//
//      如此反复，直到所有高度的人都插入到 queue 中。
//
//      上述方法的遍历顺序是：先按照高度降序，再按照 k 值升序。
//
//      所以我们可以按照这个顺序对 people 进行排序，然后依次插入到 queue 中即可。
//
//
//		时间复杂度： O(n ^ 2)
//          1. 需要对 people 全部 O(n) 个元素排序，时间复杂度为 O(nlogn)
//          2. 需要对 queue 进行插入排序，时间复杂度为 O(n ^ 2)
//		空间复杂度： O(n)
//          1. 需要维护 queue 存储全部 O(n) 个元素


impl Solution {
    pub fn reconstruct_queue(mut people: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        people.sort_by(|a, b| {
            if a[0] != b[0] {
                // 先按照高度降序排序
                b[0].cmp(&a[0])
            } else {
                // 高度相同时，再按照 k 升序排序
                a[1].cmp(&b[1])
            }
        });

        // 初始化 queue 为空
        let mut queue = Vec::with_capacity(people.len());
        for person in people {
            // 此时 queue 中所有人的高度都大于等于 person[0] ，
            // 所以在 person[1] 处插入 person 即可满足题意
            queue.insert(person[1] as usize, person);
        }

        queue
    }
}
