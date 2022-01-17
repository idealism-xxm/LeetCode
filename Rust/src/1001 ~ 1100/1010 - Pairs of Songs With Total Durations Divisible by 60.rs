// 链接：https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/
// 题意：给定一个整型数组 time ，
//      判断有多少个数对 (i, j) 满足 i < j && (time[i] + time[j]) % 60 == 0 ？

// 数据限制：
//  1 <= time.length <= 6 * 10 ^ 4
//  1 <= time[i] <= 500

// 输入： time = [30,20,150,100,40]
// 输出： 3
// 解释： (time[0] = 30, time[2] = 150): 和为 180
//       (time[1] = 20, time[3] = 100): 和为 120
//       (time[1] = 20, time[4] = 40): 和为 60

// 输入： time = [60,60,60]
// 输出： 3
// 解释： 所有 3 个数对的和都是 120 ，都满足题意


// 思路：统计
//
//      维护一个数组 cnt ，cnt[remain] 表示现在遍历过的所有数中，
//      满足 time[i] % 60 == remain 的数的个数。
//
//      同时维护一个数字 ans ，表示当前遍历过的所有数中，满足题意的数对个数。
//
//      假设现在遍历到 time[j] ，计算余数 remain = time[j] % 60 ，
//      让 time[j] 作为数对中的第二个数，
//      那么 time[0..j] 中能作为数对中第一个数的个数为 cnt[(60 - remain) % 60] ，
//      （注意这里需要让 60 - remain 模上 60 ，因为当 remain 为 0 时，需要特殊处理）
//      所以满足要求的数对可以增加 cnt[(60 - remain) % 60] 对，
//      即： ans += cnt[(60 - remain) % 60]
//
//      然后将当前数统计到 cnt 中即可，即： cnt[remain] += 1
//
//      （开辟的数组长度固定为 60 ，所以空间复杂度还是 O(1) ）
//
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)


impl Solution {
    pub fn num_pairs_divisible_by60(time: Vec<i32>) -> i32 {
        // cnt[remain] 表示模 60 后为 remain 的数字个数
        let mut cnt = [0; 60];
        // ans 表示满足要求的数字对数
        let mut ans = 0;
        // 遍历 time 中的每个数字 cur
        for cur in time {
            // 让 cur 模上 60
            let remain = (cur % 60) as usize;
            // 选择 cur 为第二个数，
            // 那么可以选择前面出现过的 cnt[(60 - remain) % 60] 个数组成满足题意的数对
            // （注意这里需要让 60 - remain 模上 60 ，因为当 remain 为 0 时，需要特殊处理）
            ans += cnt[(60 - remain) % 60];
            // remain 出现的次数 +1
            cnt[remain] += 1;
        }

        // 返回最终统计结果
        ans
    }
}
